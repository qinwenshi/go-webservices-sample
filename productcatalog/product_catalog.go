package productcatalog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	apiUrl   string
	apiKey   string
	debug    bool
	backends ProductCatalogBackend

	DefaultTTL = time.Second * 60

	httpClient = &http.Client{
		Timeout: DefaultTTL,
	}
)

type BackendConfiguration struct {
	Type       SupportedBackend
	URL        string
	HTTPClient *http.Client
}

type SupportedBackend string

const (
	PublicBackend SupportedBackend = "public"
)

type ProductCatalogBackend struct {
	Public Backend
}

type Backend interface {
	Call(method, path string, form *url.Values, header http.Header, content interface{}, v interface{}) error
}

func (s *BackendConfiguration) NewRequest(method, path, contentType string, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = s.URL + path

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		log.Printf("Cannot create product catalog request: %v\n", err)
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (s *BackendConfiguration) Do(req *http.Request, v interface{}) error {
	log.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
	start := time.Now()

	res, err := s.HTTPClient.Do(req)

	if debug {
		log.Printf("Completed in %v\n", time.Since(start))
	}

	if err != nil {
		log.Printf("Request to product catalog failed: %v\n", err)
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Cannot parse product catalog response: %v\n", err)
		return err
	}

	if debug {
		log.Printf("dimsum response: %s\n", string(resBody))
	}

	// parses error response if status code is not 2xx
	if res.StatusCode < 200 || res.StatusCode >= 400 {
		var err ErrorResp
		json.Unmarshal(resBody, &err)
		return err
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}

/**
 * @param method Either GET, POST, PUT, DELETE
 * @param path URL path
 * @param form Query string for GET method only
 * @param header Http header that need to be sent
 * @param content Interface{} of JSON object
 * @param v Any response object and fill after call success
 * @example
 *
 * // GET /some_resource?a=1
 * u := &url.Values{}l
 * u.Add("a", "1")
 * obj := RespObj{}
 * s.Call("GET", "/some_url", u, nil, obj)
 *
 * // POST /some_resource
 * c := dimsum.Content{}
 * c["data"] = "John"
 * obj := RespObj{}
 * s.Call("POST", "/some_resource", "1234567890", nil, nil, &c, obj)
 */
func (s BackendConfiguration) Call(method, path string, form *url.Values, header http.Header, content interface{}, v interface{}) error {
	var body io.Reader

	method = strings.ToUpper(method)
	if method == "GET" {
		if form != nil && len(*form) > 0 {
			path += "?" + form.Encode()
		}
	} else {
		// POST, PUT, DELETE
		if content != nil {
			encoded, _ := json.Marshal(content)

			fmt.Printf("body: %s\n", string(encoded))
			body = bytes.NewBuffer(encoded)
		}
	}

	req, err := s.NewRequest(method, path, "application/json", body)
	if err != nil {
		return err
	}

	for k, vs := range header {
		if len(vs) != 0 {
			req.Header.Add(k, vs[0])
		}
	}

	if err := s.Do(req, v); err != nil {
		return err
	}

	return nil
}

func GetBackend(backend SupportedBackend) Backend {
	var ret Backend
	switch backend {
	case PublicBackend:
		if backends.Public == nil {
			backends.Public = BackendConfiguration{backend, apiUrl, httpClient}
		}

		ret = backends.Public
	}

	return ret
}

func SetDebug(value bool) {
	debug = value
}

func Setup(url, key string) {
	apiUrl = url
	apiKey = key
}

func SetClientTimeout(t time.Duration) {
	httpClient.Timeout = t
}

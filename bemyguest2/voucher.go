package bemyguest2

type VouchersResp struct {
	Data []Voucher `json:"data"`

	// ignore this field since inconsistent type for meta.pagination.links
	//Meta Meta      `json:"meta"`
}

type Voucher struct {
	UUID         string  `json:"uuid"`
	GeneratedAt  *string `json:"generatedAt"`
	DownloadedAt *string `json:"downloadedAt"`
	Links        []Link  `json:"links"`
}

//func (r *VouchersResp) HasNext() bool {
//	return r.Meta.Pagination.HasNext()
//}
//
//func (r *VouchersResp) GetNextPage() int {
//	return r.Meta.Pagination.GetNextPage()
//}

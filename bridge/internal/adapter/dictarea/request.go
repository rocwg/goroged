package dictarea

type GetAreaByParentRequest struct {
	ParentCode string `form:"parent_code"`
}

type SearchAreaRequest struct {
	Keyword string `json:"keyword"`
	Limit   int32  `json:"limit"`
}

type BatchGetAreaByCodesRequest struct {
	AreaCodes []string `json:"area_codes"`
}

type GetDictByTypesRequest struct {
	TypeCodes []string `json:"type_codes"`
}

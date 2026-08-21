package dictarea

type GetAreaByParentRequest struct {
	ParentCode string
}

type SearchAreaRequest struct {
	Keyword string
	Limit   int32
}

type BatchGetAreaByCodesRequest struct {
	AreaCodes []string `json:"area_codes"`
}

type GetDictByTypesRequest struct {
	TypeCodes []string `json:"type_codes"`
}

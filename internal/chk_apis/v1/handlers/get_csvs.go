package handlers

type ListRequest struct {
	Filters []*Filter `json:"filters,omitempty"`
	Sort    Sort      `json:"sort,omitempty"`
	Limit   int       `json:"limit,omitempty"`
	Page    int       `json:"page,omitempty"`
}

func (p *ListRequest) GetLimit() (limit int) {

	if p.Limit <= 0 {
		limit = 5
	} else if p.Limit > 100 {
		limit = 100
	} else {
		limit = p.Limit
	}

	return
}

func (p *ListRequest) GetPage() (page int) {

	if p.Page <= 0 {
		page = 1
	} else {
		page = p.Page
	}

	return
}

type Filter struct {
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
	Method string `json:"method,omitempty"`
}

func (f *Filter) GetKey() (key string) {

	switch f.Key {
	case "id", "unix", "symbol", "open", "high", "low", "close":
		key = f.Key
	default:
		key = "id"
	}

	return
}

func (f *Filter) GetValue() (value string) {

	return f.Value
}

func (f *Filter) GetMethod() (method string) {

	switch f.Method {
	case ">", ">=", "<", "<=", "=":
		method = f.Method
	default:
		method = "="
	}

	return
}

type Sort struct {
	OrderBy string `json:"order_by"`
	IsACS   bool   `json:"is_acs"`
}

func (s *Sort) GetOrderBy() (oderBy string) {

	if s.OrderBy == "" {
		oderBy = "id"
	} else {
		oderBy = s.OrderBy
	}

	return
}

func (s *Sort) GetSort() (sort string) {

	if s.IsACS {
		sort = "ASC"
	} else {
		sort = "DESC"
	}

	return
}

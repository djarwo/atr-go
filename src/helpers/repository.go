package helpers

//FindAllParams is Parameters helpers for FindAllParams
type FindAllParams struct {
	Page           string
	Size           string
	Keyword        string
	KeywordName    string
	StatusID       string
	Query          string
	SortBy         string
	SortName       string
	GroupBy        string
	BusinessID     string
	OutletID       string
	QueryCondition string
	DataFinder     string
}

//ReturnRepo is Parameters helpers for ReturnRepo
type ReturnRepo struct {
	StatusCode int
	Message    string
	Err        error
	Data       interface{}
}

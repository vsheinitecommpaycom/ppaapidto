package ppaapidto

type QueryParametersDto struct {
	DateSince  string   `json:"dateSince"`
	DateTill   string   `json:"dateTill"`
	ProjectIds []uint32 `json:"projectIds,omitempty"`
}

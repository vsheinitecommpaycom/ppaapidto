package dto

type QueryParameters struct {
	DateSince  string   `json:"dateSince"`
	DateTill   string   `json:"dateTill"`
	ProjectIds []uint32 `json:"projectIds,omitempty"`
}

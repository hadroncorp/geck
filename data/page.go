package data

// PaginationType the type of pagination for a Page.
type PaginationType string

const (
	PaginationTypeOffset PaginationType = "OFFSET"
	PaginationTypeKeySet PaginationType = "KEY_SET"
	PaginationTypeCursor PaginationType = "CURSOR"
)

// Page a chunk of items from a dataset with metadata used to fetch more data.
type Page[T any] struct {
	PreviousPageToken PageToken `json:"previous_page_token"`
	NextPageToken     PageToken `json:"next_page_token"`
	TotalItems        int       `json:"total_items"`
	Items             []T       `json:"items"`
}

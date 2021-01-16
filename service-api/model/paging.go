package model

// Paging ...
type Paging struct {
	Sort         string
	SortOrder    int64
	Page         int64
	CountPerPage int64
}

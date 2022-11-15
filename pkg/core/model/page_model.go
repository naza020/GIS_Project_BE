package model

// PageLimit ...
type PageLimit struct {
	PageNumber int64 `json:"pageNumber" default:"1"`
	PageSize   int64 `json:"pageSize" default:"10"`
}

// PageResult ...
type PageResult struct {
	PageSize   int64 `json:"pageSize"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	PageNumber int64 `json:"pageNumber"`
}

// PageQuery ...
type PageQuery struct {
	Offset   int64 `json:"offset"`
	PageSize int64 `json:"pageSize"`
}

// Sort ...
type Sort struct {
	SortBy        string `json:"sortBy" `
	SortDirection string `json:"sortDirection" example:"desc" enums:"asc,desc"`
}

var (
	MaxLimitForQuery = &PageLimit{
		PageNumber: 1,
		PageSize:   1000,
	}

	LimitForCount = &PageLimit{
		PageNumber: 1,
		PageSize:   1,
	}
)

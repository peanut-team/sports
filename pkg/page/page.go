package page

type Paginate struct {
	Page     int64 `json:"page,omitempty" form:"page"`
	PageSize int64 `json:"page_size,omitempty" form:"page_size"`
	Total    int64 `json:"total,omitempty" form:"total" `
	// Sort     string `json:"sort,omitempty" form:"sort"`
	// Search   string `json:"search,omitempty" form:"search"`
}

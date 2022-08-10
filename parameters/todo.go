package parameters

type FilterTodo struct {
	Limit      uint   `url:"limit"`
	Offset     uint   `url:"offset"`
	IsDone     bool   `url:"is_done"`
	OrderBy    string `url:"order_by"`
	OrderField string `url:"order_field"`
}

type NewTodo struct {
	Title string `json:"title" form:"title" validator:"required"`
	Body  string `json:"body" form:"body" validator:"required"`
}

type UpdateTodo struct {
	Title *string `json:"title" form:"title"`
	Body  *string `json:"body" form:"body"`
}

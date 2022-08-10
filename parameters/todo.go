package parameters

type FilterTodo struct {
	Limit      *int    `url:"limit"`
	Offset     *int    `url:"offset"`
	IsDone     *bool   `url:"is_done"`
	OrderBy    *string `url:"order_by"`
	OrderField *string `url:"order_field"`
	Title      *string
}

func (filter *FilterTodo) GetLimit() int {
	if filter.Limit == nil {
		return 20
	}
	return *filter.Limit
}
func (filter *FilterTodo) GetOffest() int {
	if filter.Limit == nil {
		return 0
	}
	return *filter.Limit
}

func (filter *FilterTodo) OrderQueryBy() string {
	if filter.OrderBy == nil || filter.OrderField == nil {
		return `desc + created_at`
	} else {
		return *filter.OrderBy + " " + *filter.OrderField
	}
}

type NewTodo struct {
	Title string `json:"title" form:"title" validator:"required"`
	Body  string `json:"body" form:"body" validator:"required"`
}

type UpdateTodo struct {
	Title *string `json:"title" form:"title"`
	Body  *string `json:"body" form:"body"`
}

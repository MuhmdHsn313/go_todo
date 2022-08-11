package parameters

type NewUserParams struct {
	FullName string `json:"full_name" form:"full_name" validator:"required"`
	Email    string `json:"email" form:"email" validator:"required,email"`
	Password string `json:"password" form:"password" validator:"required"`
}

type NewSessionParams struct {
	UserID uint
}

type LoginParams struct {
	Email    string `json:"email" form:"email" validator:"required,email"`
	Password string `json:"password" form:"password" validator:"required"`
}

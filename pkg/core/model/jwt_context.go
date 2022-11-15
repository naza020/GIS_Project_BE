package model

type JwtContext struct {
	User  JwtContextUser `json:"user"`
	Roles []string       `json:"roles"`
}

type JwtContextUser struct {
	ID        int     `json:"id"`
	Email     *string `json:"email"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
}

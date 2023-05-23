package dao

type Me struct {
	Email      string  `json:"email,omitempty" sql:"email" validate:"required" example:"myemail@gmail.com"`
	Name       string  `json:"name,omitempty" sql:"name" validate:"required" example:"user94"`
	FirstName  string  `json:"firstName,omitempty" sql:"first_name" validate:"required" example:"Ivan"`
	LastName   string  `json:"lastName,omitempty" sql:"last_name" validate:"required" example:"Ivanov"`
	Role       string  `json:"role,omitempty" sql:"role" validate:"required" example:"USER"`
	CompanyID  int     `json:"-" sql:"company_id"`
	FatherName *string `json:"fatherName,omitempty" sql:"father_name" example:"Ivanovich"`
	Position   *string `json:"position,omitempty" sql:"position" example:"Director"`
	Country    *string `json:"country,omitempty" sql:"country" example:"Россия"`
	City       *string `json:"city,omitempty" sql:"city" example:"Москва"`
	Biography  *string `json:"biography,omitempty" sql:"biography" example:"I'd like to relax"`
}

type Session struct {
	ID      int    `json:"id" redis:"ID"`
	IP      string `json:"ip" redis:"IP"`
	Device  string `json:"device" redis:"Device"`
	Browser string `json:"browser" redis:"Browser"`
	Updated int64  `json:"updated" redis:"Updated"`
}

package dao

type Me struct {
	Email      string  `json:"email,omitempty" sql:"email" example:"myemail@gmail.com"`
	Name       string  `json:"name,omitempty" sql:"name" example:"user94"`
	INN        string  `json:"inn,omitempty" sql:"company_inn" example:"7707083893"`
	FirstName  string  `json:"firstName,omitempty" sql:"first_name" example:"Ivan"`
	LastName   string  `json:"lastName,omitempty" sql:"last_name" example:"Ivanov"`
	FatherName *string `json:"fatherName,omitempty" sql:"father_name" example:"Ivanovich"`
	Position   *string `json:"position,omitempty" sql:"position" example:"Director"`
	Country    *string `json:"country,omitempty" sql:"country" example:"Россия"`
	City       *string `json:"city,omitempty" sql:"city" example:"Москва"`
	Biography  *string `json:"biography,omitempty" sql:"biography" example:"I'd like to relax"`
}

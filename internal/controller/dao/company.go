package dao

type Company struct {
	INN     string  `json:"inn,omitempty" sql:"inn" validate:"required,inn" example:"7707083893"`
	Name    *string `json:"name,omitempty" sql:"name" validate:"omitempty,gte=2,lte=150" example:"ООО 'Парк'"`
	Website *string `json:"website,omitempty" sql:"website" validate:"omitempty,link" example:"https://www.rusprofile.ru"`
}

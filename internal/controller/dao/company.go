package dao

type Company struct {
	INN     string  `json:"inn,omitempty" validate:"required,inn" example:"7707083893"`
	Name    *string `json:"name,omitempty" validate:"omitempty,gte=2,lte=150" example:"ООО 'Парк'"`
	Website *string `json:"website,omitempty" validate:"omitempty,link" example:"https://www.rusprofile.ru"`
}

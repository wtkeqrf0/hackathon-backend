package dao

type Company struct {
	INN     string  `json:"inn,omitempty" validate:"required,inn" required:"true" example:"7707083893"`
	Name    *string `json:"name,omitempty" example:"ООО 'Парк'"`
	Website *string `json:"website,omitempty" example:"https://www.rusprofile.ru"`
}

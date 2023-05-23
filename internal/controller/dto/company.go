package dto

type UpdateCompany struct {
	INN     *string `json:"inn,omitempty" validate:"omitempty,inn" example:"7707083893"`
	Name    *string `json:"name,omitempty" validate:"omitempty,gte=2,lte=150" example:"ООО 'Парк'"`
	Website *string `json:"website,omitempty" validate:"omitempty,link" example:"https://www.rusprofile.ru"`
}

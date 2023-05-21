package dto

type UpdateCompany struct {
	Name    *string `json:"name,omitempty" example:"ООО 'Парк'"`
	Website *string `json:"website,omitempty" example:"https://www.rusprofile.ru"`
}

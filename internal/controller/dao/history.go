package dao

type Histories struct {
	ID          int    `json:"id" sql:"id" validate:"required" example:"12"`
	CompanyName string `json:"companyName" sql:"company_name" validate:"required" example:"ООО 'Парк'"`
}

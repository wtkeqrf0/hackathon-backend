package dao

type Histories struct {
	ID   int    `json:"id" sql:"id" validate:"required" example:"12"`
	Name string `json:"name" sql:"name" validate:"required" example:"ООО 'Парк'"`
}

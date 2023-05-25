package dto

type IndustryType struct {
	Type string `json:"type,omitempty" validate:"required,lte=150"`
}

package dto

type TaxationSystem struct {
	Operations int     `json:"operations,omitempty" validate:"required,lte=0" sql:"operations" example:"20"`
	USN6       float64 `json:"usn6,omitempty" validate:"required,lte=0" sql:"usn6" example:"84.2"`
	USN15      float64 `json:"usn15,omitempty" validate:"required,lte=0" sql:"usn15" example:"57.2"`
	OSN        float64 `json:"osn,omitempty" validate:"required,lte=0" sql:"osn" example:"49.5"`
}

type BusinessActivity struct {
	Type    string `json:"type,omitempty" validate:"required" sql:"type" example:"Много букав"`
	SubType string `json:"subType,omitempty" validate:"required" sql:"sub_type" example:"Много букав"`
}

type Equipment struct {
	Type  string  `json:"type,omitempty" sql:"type" validate:"required" example:"Станок"`
	Price float64 `json:"price,omitempty" sql:"price" validate:"omitempty,lte=0" example:"3058.12"`
}

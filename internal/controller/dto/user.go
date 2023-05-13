package dto

type EmailWithPassword struct {
	Email    string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Password string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" example:"onkr3451"`
}

type Token struct {
	Authorization string `json:"Authorization" header:"Authorization" validate:"required,jwt"`
}

package dto

import "github.com/wtkeqrf0/while.act/internal/controller/dao"

type SignUp struct {
	Email      string      `json:"email,omitempty" validate:"required,email" required:"true" example:"myemail@gmail.com"`
	Password   string      `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" required:"true" example:"bob126"`
	FirstName  string      `json:"firstName,omitempty" validate:"required,gte=2,lte=30" required:"true" example:"Ivan"`
	LastName   string      `json:"lastName,omitempty" validate:"required,gte=2,lte=30" required:"true" example:"Ivanov"`
	FatherName *string     `json:"fatherName,omitempty" example:"Ivanovich"`
	Position   *string     `json:"position,omitempty" example:"Director"`
	Country    *string     `json:"country,omitempty" example:"Россия"`
	City       *string     `json:"city,omitempty" example:"Москва"`
	Biography  *string     `json:"biography,omitempty" example:"I'd like to relax"`
	Company    dao.Company `json:"company,omitempty" required:"true"`
}

type SignIn struct {
	Email    string `json:"email,omitempty" validate:"required,email" required:"true" example:"myemail@gmail.com"`
	Password string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" required:"true" example:"bob126"`
}

type UpdateUser struct {
	FirstName  *string `json:"firstName,omitempty" required:"true" example:"Ivan"`
	LastName   *string `json:"lastName,omitempty" required:"true" example:"Ivanov"`
	FatherName *string `json:"fatherName,omitempty" example:"Ivanovich"`
	Position   *string `json:"position,omitempty" example:"Director"`
	Country    *string `json:"country,omitempty" example:"Россия"`
	City       *string `json:"city,omitempty" example:"Москва"`
	Biography  *string `json:"biography,omitempty" example:"I'd like to relax"`
}

type UpdatePassword struct {
	Password    string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" required:"true" example:"bob126"`
	NewPassword string `json:"newPassword,omitempty" validate:"required,printascii,gte=4,lte=20" required:"true" example:"mob126"`
}

type Token struct {
	Authorization string `json:"Authorization" header:"Authorization" required:"true" validate:"required,jwt"`
}

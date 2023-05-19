package dto

import "github.com/wtkeqrf0/while.act/internal/controller/dao"

type SignUp struct {
	Email      string      `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Password   string      `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" example:"bob126"`
	FirstName  string      `json:"firstName,omitempty" validate:"required,gte=2,lte=30" example:"Ivan"`
	LastName   string      `json:"lastName,omitempty" validate:"required,gte=2,lte=30" example:"Ivanov"`
	FatherName *string     `json:"fatherName,omitempty" example:"Ivanovich"`
	Position   *string     `json:"position,omitempty" example:"Director"`
	Country    *string     `json:"country,omitempty" example:"Россия"`
	City       *string     `json:"city,omitempty" example:"Москва"`
	Biography  *string     `json:"biography,omitempty" example:"I'd like to relax"`
	Company    dao.Company `json:"company,omitempty"`
}

type SignIn struct {
	Email    string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Password string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" example:"bob126"`
}

type UpdateUser struct {
	FatherName *string `json:"fatherName,omitempty" example:"Ivanovich"`
	Position   *string `json:"position,omitempty" example:"Director"`
	Country    *string `json:"country,omitempty" example:"Россия"`
	City       *string `json:"city,omitempty" example:"Москва"`
	Biography  *string `json:"biography,omitempty" example:"I'd like to relax"`
}

type Token struct {
	Authorization string `json:"Authorization" header:"Authorization" validate:"required,jwt"`
}

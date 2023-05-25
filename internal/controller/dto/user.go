package dto

import "github.com/while-act/hackathon-backend/internal/controller/dao"

type SignUp struct {
	Email      string      `json:"email,omitempty" validate:"required,email" required:"true" example:"myemail@gmail.com"`
	Password   string      `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" required:"true" example:"bob126"`
	FirstName  string      `json:"firstName,omitempty" validate:"required,gte=2,lte=30" required:"true" example:"Ivan"`
	LastName   string      `json:"lastName,omitempty" validate:"required,gte=2,lte=30" required:"true" example:"Ivanov"`
	FatherName *string     `json:"fatherName,omitempty" validate:"omitempty,gte=2,lte=30" example:"Ivanovich"`
	Position   *string     `json:"position,omitempty" validate:"omitempty,gte=2,lte=50" example:"Director"`
	Country    *string     `json:"country,omitempty" validate:"omitempty,title" example:"Россия"`
	City       *string     `json:"city,omitempty" validate:"omitempty,title" example:"Москва"`
	Biography  *string     `json:"biography,omitempty" validate:"omitempty,lte=1024" example:"I'd like to relax"`
	Company    dao.Company `json:"company,omitempty"`
}

type SignIn struct {
	Email    string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Password string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" example:"bob126"`
}

type Email struct {
	Email string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
}

type UpdateUser struct {
	FirstName  *string `json:"firstName,omitempty" validate:"omitempty,gte=2,lte=30" example:"Ivan"`
	LastName   *string `json:"lastName,omitempty" validate:"omitempty,gte=2,lte=30" example:"Ivanov"`
	FatherName *string `json:"fatherName,omitempty" validate:"omitempty,gte=2,lte=30" example:"Ivanovich"`
	Position   *string `json:"position,omitempty" validate:"omitempty,gte=2,lte=50" example:"Director"`
	Country    *string `json:"country,omitempty" validate:"omitempty,title" example:"Россия"`
	City       *string `json:"city,omitempty" validate:"omitempty,title" example:"Москва"`
	Biography  *string `json:"biography,omitempty" validate:"omitempty,lte=1024" example:"I'd like to relax"`
}

type UpdatePassword struct {
	Email       string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Code        string `json:"code,omitempty" validate:"required,len=5" example:"N1OSP"`
	NewPassword string `json:"newPassword,omitempty" validate:"required,printascii,gte=4,lte=20" example:"mob126"`
}

type UpdateEmail struct {
	NewEmail string `json:"newEmail,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Password string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" example:"mob126"`
}

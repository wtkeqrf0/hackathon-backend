package bind

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type Valid struct {
	v *validator.Validate
}

func NewValid(v *validator.Validate) *Valid {
	if err := v.RegisterValidation("name", validateName); err != nil {
		logrus.WithError(err).Warn("can't validate name")
	}

	if err := v.RegisterValidation("inn", validateInn); err != nil {
		logrus.WithError(err).Warn("can't validate inn")
	}

	if err := v.RegisterValidation("link", validateLink); err != nil {
		logrus.WithError(err).Warn("can't validate link")
	}

	if err := v.RegisterValidation("title", validateTitle); err != nil {
		logrus.WithError(err).Warn("can't validate title")
	}

	return &Valid{v: v}
}

func (v *Valid) Engine() any {
	return v.v
}

func (v *Valid) ValidateStruct(o any) error {
	return v.v.Struct(o)
}

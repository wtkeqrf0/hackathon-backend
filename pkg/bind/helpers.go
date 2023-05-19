package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
	"regexp"
)

var (
	NameRegexp  = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{3,18}([a-zA-Z0-9])$`)
	InnRegexp   = regexp.MustCompile(`^(\d{10}|\d{12})$`)
	LinkRegexp  = regexp.MustCompile(`^https?://(?:www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b[-a-zA-Z0-9()@:%_+.~#?&/=]*$`)
	TitleRegexp = regexp.MustCompile(`^\p{L}+(?:([ \-']|(\. ))\p{L}+)*$`)
	EmailRegexp = regexp.MustCompile(`^\S+@\S+\.\S+$`)
)

// FillStruct of given generic type by request JSON body
func FillStruct[T any](c *gin.Context) (t T, ok bool) {
	if err := c.ShouldBindJSON(&t); err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}
	return t, true
}

func validateName(fl validator.FieldLevel) bool {
	return NameRegexp.MatchString(fl.Field().String())
}

func validateInn(fl validator.FieldLevel) bool {
	return InnRegexp.MatchString(fl.Field().String())
}

func validateLink(fl validator.FieldLevel) bool {
	return LinkRegexp.MatchString(fl.Field().String())
}

func validateTitle(fl validator.FieldLevel) bool {
	return TitleRegexp.MatchString(fl.Field().String())
}

func validateEmail(fl validator.FieldLevel) bool {
	return EmailRegexp.MatchString(fl.Field().String())
}

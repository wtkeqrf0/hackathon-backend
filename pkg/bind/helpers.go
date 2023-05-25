package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/while-act/hackathon-backend/pkg/middleware/errs"
	"regexp"
)

var (
	NameRegexp  = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{3,18}([a-zA-Z0-9])$`)
	InnRegexp   = regexp.MustCompile(`^(\d{10}|\d{12})$`)
	LinkRegexp  = regexp.MustCompile(`^https?://(?:www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b[-a-zA-Z0-9()@:%_+.~#?&/=]*$`)
	TitleRegexp = regexp.MustCompile(`^\p{L}+(?:([ \-']|(\. ))\p{L}+)*$`)
	EmailRegexp = regexp.MustCompile(`^\S+@\S+\.\S+$`)
	UUID4       = regexp.MustCompile(`^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$`)
)

// FillStructJSON of given generic type by request JSON body
func FillStructJSON[T any](c *gin.Context) (t *T) {
	if err := c.ShouldBindJSON(&t); err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}
	return
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

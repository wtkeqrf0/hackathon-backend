package errs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/wtkeqrf0/while.act/ent"
	"io"
	"net/http"
)

// Sign-in errors
var (
	PasswordError = newError(http.StatusBadRequest, "Wrong password", "You can still sign in by your email!")
)

// Auth errors
var (
	NoSuchUser   = newError(http.StatusNotFound, "There is no such user", "But you can still find another existing user!")
	UnAuthorized = newError(http.StatusUnauthorized, "You are not logged in", "Click on the button below to sign in!")
)

// Input errors
var (
	ValidError = newError(http.StatusBadRequest, "Validation error", "Try to enter the correct data")
)

// Enity errors
var (
	NoSuchCompany = newError(http.StatusBadRequest, "There is no such company inn", "But you can still find another existing company!")
)

// ServerError errors
var (
	ServerError = newError(http.StatusInternalServerError, "Server exception was occurred", "Try to restart the page")
)

type ErrHandler struct {
	log *logrus.Logger
}

func NewErrHandler(log *logrus.Logger, out io.Writer) *ErrHandler {
	log.SetLevel(logrus.ErrorLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006/01/02 15:32:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "status",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyMsg:   "message",
		},
	})
	log.SetOutput(out)

	return &ErrHandler{log: log}
}

func (e ErrHandler) HandleErrors(c *gin.Context) {
	c.Next()

	errs := c.Errors

	if errs.Last() == nil {
		return
	}

	for i, err := range errs {
		if my, ok := err.Err.(MyError); ok {
			e.log.WithError(my.Err).Errorf("%02d# %s", i+1, my.Msg)
			res := gin.H{"error": my.Msg, "advice": my.Advice}

			if vErrs, ok := my.Err.(validator.ValidationErrors); ok {
				fields := make(gin.H)

				for _, vErr := range vErrs {
					field := vErr.Field()
					if field == "" {
						field = "Field"
					}
					switch vErr.Tag() {
					case "email":
						fields[field] = fmt.Sprintf("%s is not the correct email", field)
					case "required":
						fields[field] = fmt.Sprintf("%s should not be empty", field)
					case "numeric":
						fields[field] = fmt.Sprintf("%s must be a number", field)
					case "gte":
						fields[field] = fmt.Sprintf("%s must be greater or equal %s", field, vErr.Param())
					case "lte":
						fields[field] = fmt.Sprintf("%s must be lesser or equal %s", field, vErr.Param())
					case "len":
						fields[field] = fmt.Sprintf("%s must have a length of %s", field, vErr.Param())
					case "gt":
						fields[field] = fmt.Sprintf("%s must be greater than %s", field, vErr.Param())
					case "lt":
						fields[field] = fmt.Sprintf("%s must be lesser than %s", field, vErr.Param())
					case "printascii":
						fields[field] = fmt.Sprintf("%s can contain only /:@-._~!?$&'()*+,;= and any english letters", field)
					case "required_without_all":
						fields[field] = fmt.Sprintf("%s should not be empty", field)
					case "name":
						fields[field] = fmt.Sprintf("name is not valid")
					case "jwt":
						fields[field] = fmt.Sprintf("%s is not jwt format", field)
					case "title":
						fields[field] = fmt.Sprintf("%s is not a title", field)
					case "inn":
						fields[field] = fmt.Sprintf("%s is not an INN", field)
					case "link":
						fields[field] = fmt.Sprintf("%s is not a link", field)
					}
				}
				res["fields"] = fields

			} else if ent.IsValidationError(err.Err) {
				my.Status = http.StatusBadRequest
				res["fields"] = gin.H{"FieldName": err.Err.Error()}
			}

			if i == 0 {
				c.JSON(my.Status, res)
			}
		} else {
			e.log.WithError(err.Err).Error("UNEXPECTED ERROR")
			if i == 0 {
				c.JSON(ServerError.Status, ServerError.Msg)
			}
		}
	}
}

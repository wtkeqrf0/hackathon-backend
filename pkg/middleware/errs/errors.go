package errs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/while-act/hackathon-backend/ent"
	"net/http"
	"os"
)

// Sign-in errors
var (
	PasswordError = newError(http.StatusBadRequest, "Wrong password", "You can still sign in by your email!")
	CodeError     = newError(http.StatusBadRequest, "Code is not correct", "Try to request a new one")
)

// Auth errors
var (
	UnAuthorized = newError(http.StatusUnauthorized, "You are not logged in", "Click on the button below to sign in!")
)

// Input errors
var (
	ValidError = newError(http.StatusBadRequest, "Validation error", "Try to enter the correct data")
)

// Entity errors
var (
	NoSuchUser     = newError(http.StatusNotFound, "There is no such user", "But you can still find another existing user!")
	NoSuchCompany  = newError(http.StatusBadRequest, "There is no such industry branch", "But you can still find another existing industry!")
	NoSuchIndustry = newError(http.StatusBadRequest, "There is no such company inn", "But you can still find another existing company!")
)

// ServerError errors
var (
	EmailError  = newError(http.StatusInternalServerError, "Can't send message to your email", "Try to send it later")
	ServerError = newError(http.StatusInternalServerError, "Server exception was occurred", "Try to restart the page")
)

type ErrHandler struct {
	log *logrus.Logger
}

func NewErrHandler(log *logrus.Logger) *ErrHandler {
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
	log.SetOutput(os.Stderr)

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

			vErrs, oki := my.Err.(validator.ValidationErrors)
			switch {
			case oki:
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

			case ent.IsValidationError(my.Err):
				my.Status = http.StatusBadRequest
				my.Msg = "That's not a correct value"
				my.Advice = "Try to write another value"

			case ent.IsConstraintError(my.Err):
				my.Status = http.StatusBadRequest
				my.Msg = "This value doesn't exist"
				my.Advice = "Try to write another value"

			case ent.IsNotFound(err):
				my.Status = http.StatusNotFound
				my.Msg = "There is no such object"
				my.Advice = "But you can still find another existing object!"
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

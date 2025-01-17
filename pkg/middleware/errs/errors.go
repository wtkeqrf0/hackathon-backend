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

// ServerError errors
var (
	EmailError  = newError(http.StatusInternalServerError, "Can't send message to your email", "Try to send it later")
	ServerError = newError(http.StatusInternalServerError, "Server exception was occurred", "Try to restart the page")
	PDFError    = newError(http.StatusInternalServerError, "Can't create pdf file", "Try to create it later")
)

const (
	message = "message"
	advice  = "advice"
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

// HandleErrors of MyError type
func (e *ErrHandler) HandleErrors(c *gin.Context) {
	c.Next()

	errs := c.Errors

	if errs.Last() == nil {
		return
	}

	for i, err := range errs {
		res := gin.H{message: ServerError.Msg, advice: ServerError.Advice}
		status := http.StatusInternalServerError
		resErr := err.Err

		switch err.Err.(type) {
		case MyError:
			my := err.Err.(MyError)
			resErr = my.Err
			status = my.Status
			res[message] = my.Msg
			res[advice] = my.Advice

		case validator.ValidationErrors:
			vErrs := err.Err.(validator.ValidationErrors)
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
			status = http.StatusBadRequest
			res[message] = "Validation failed"
			res[advice] = "Try to send another data"
		}
		switch {
		case ent.IsValidationError(err.Err):
			res[message] = "Validation failed"
			res[advice] = "Try to send another data"
			status = http.StatusBadRequest

		case ent.IsConstraintError(err.Err):
			res[message] = "This value doesn't exist"
			res[advice] = "Try to write another value"
			status = http.StatusBadRequest

		case ent.IsNotFound(err.Err):
			res[message] = "There is no such object"
			res[advice] = "But you can still find another existing object!"
			status = http.StatusBadRequest
		case ent.IsNotSingular(err):
			res[message] = "There are many objects"
			res[advice] = "You must find a unique object"
			status = http.StatusBadRequest
		}

		e.log.WithError(resErr).Errorf("%02d# %s", i+1, res[message])
		if i == 0 {
			c.JSON(status, res)
		}
	}
}

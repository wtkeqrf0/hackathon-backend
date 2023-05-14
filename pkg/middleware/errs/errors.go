package errs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// Sign-in errors
var (
	NoSuchUser       = newError(http.StatusNotFound, "There is no such user", "But you can still find another existing user!")
	PasswordError    = newError(http.StatusBadRequest, "Wrong password", "You can still sign in by your email!")
	PasswordNotFound = newError(http.StatusNotFound, "You have not registered a password for you account", "Try change the password in your profile")
)

// Auth errors
var (
	PermissionError = newError(http.StatusForbidden, "You don't have this permission", "Try to ask the owner about it")
	UnAuthorized    = newError(http.StatusUnauthorized, "You are not logged in", "Click on the button below to sign in!")
)

// Input errors
var (
	ValidError   = newError(http.StatusBadRequest, "Validation error", "Try to enter the correct data")
	DataError    = newError(http.StatusBadRequest, "Insufficient data", "Try to enter the remaining data")
	AlreadyExist = newError(http.StatusBadRequest, "Already exist", "Try to enter another data")
)

// ServerError errors
var (
	ServerError = newError(http.StatusInternalServerError, "Server exception was occurred", "Try to restart the page")
)

var logger = logrus.Logger{
	Level:        logrus.ErrorLevel,
	Out:          os.Stderr,
	ReportCaller: true,
	Formatter: &logrus.JSONFormatter{
		TimestampFormat: "2006/01/02 15:32:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "status",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyMsg:   "message",
		},
	},
}

// ErrorHandler used for error handling. Handles only MyError type errors
func ErrorHandler(c *gin.Context) {
	c.Next()

	errs := c.Errors

	if errs.Last() == nil {
		return
	}

	for i, err := range errs {
		if my, ok := err.Err.(MyError); ok {
			logger.WithError(my.Err).Errorf("%02d# %s", i+1, my.Msg)
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
						break
					case "name":
						fields[field] = fmt.Sprintf("name is not valid")
					case "jwt":
						fields[field] = fmt.Sprintf("%s is not jwt format", field)

					}
				}
				res["fields"] = fields
			}

			if i == 0 {
				c.JSON(my.Status, res)
			}
		} else {
			logger.WithError(err.Err).Error("UNEXPECTED ERROR")
			if i == 0 {
				c.JSON(ServerError.Status, ServerError.Msg)
			}
		}
	}
}
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
	"net/http"
)

// GetMe godoc
// @Summary Get detail data about the user by jwt
// @Description Returns detail information about me (jwt required)
// @Tags User Get
// @Param Authorization header string true "User's session"
// @Success 200 {object} ent.User "Info about jwt user"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 404 {object} errs.MyError "User doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /auth [get]
func (h Handler) getMe(c *gin.Context) {
	id, ok := h.jwt.GetUserId(c)
	if !ok {
		return
	}

	user, err := h.users.FindUserByID(id)
	if err != nil {
		if ent.IsNotFound(err) {
			c.Error(errs.NoSuchUser.AddErr(err))
		} else {
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

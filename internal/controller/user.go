package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
	"github.com/wtkeqrf0/while.act/pkg/bind"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
	"net/http"
)

// GetMe godoc
// @Summary Get detail data about the user by jwt
// @Security ApiKeyAuth
// @Description Returns detail information about me
// @Tags User
// @Success 200 {object} dao.Me "Info about session"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 404 {object} errs.MyError "User doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /auth/session [get]
func (h *Handler) getMe(c *gin.Context) {
	id, ok := h.jwt.GetUserId(c)
	if !ok {
		return
	}

	user, err := h.user.FindUserByID(id)
	if err != nil {
		if myErr, ok := err.(errs.MyError); ok {
			c.Error(myErr)
		} else {
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateMe godoc
// @Summary Update user's data
// @Security ApiKeyAuth
// @Description Updates user's not required info
// @Tags User
// @Param updFields body dto.UpdateUser true "Fields to update"
// @Success 200 "Successfully Updated"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 404 {object} errs.MyError "User doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /user [patch]
func (h *Handler) updateMe(c *gin.Context) {
	id, ok := h.jwt.GetUserId(c)
	if !ok {
		return
	}

	updFields, ok := bind.FillStruct[dto.UpdateUser](c)
	if !ok {
		return
	}

	if err := h.user.UpdateUser(updFields, id); err != nil {
		switch {
		case ent.IsNotFound(err):
			c.Error(errs.NoSuchUser.AddErr(err))
		case ent.IsValidationError(err):
			c.Error(errs.ValidError.AddErr(err))
		default:
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.Status(http.StatusOK)
}

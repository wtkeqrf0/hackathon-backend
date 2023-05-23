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
// @Summary Get detail data about the user by session
// @Security ApiKeyAuth
// @Description Returns detail information about me
// @Tags Session
// @Success 200 {object} dao.Me "Info about session"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 404 {object} errs.MyError "User doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /auth/session [get]
func (h *Handler) getMe(c *gin.Context) {
	s, err := h.session.GetSession(c)
	if err != nil {
		c.Error(errs.UnAuthorized.AddErr(err))
		return
	}

	user, err := h.user.FindUserByID(s.ID)
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
// @Description Updates user's info
// @Tags User
// @Param updFields body dto.UpdateUser true "Fields to update"
// @Success 200 "Successfully Updated"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 404 {object} errs.MyError "User doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /user [patch]
func (h *Handler) updateMe(c *gin.Context) {
	s, err := h.session.GetSession(c)
	if err != nil {
		c.Error(errs.UnAuthorized.AddErr(err))
		return
	}

	updFields, ok := bind.FillStruct[dto.UpdateUser](c)
	if !ok {
		return
	}

	if err = h.user.UpdateUser(updFields, s.ID); err != nil {
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

// UpdatePassword godoc
// @Summary Update user's password
// @Description Updates user's password
// @Tags User
// @Param updPassword body dto.UpdatePassword true "Email with new password"
// @Success 200 "Successfully Updated"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 404 {object} errs.MyError "User doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /user/password [patch]
func (h *Handler) updatePassword(c *gin.Context) {
	updPassword, ok := bind.FillStruct[dto.UpdatePassword](c)
	if !ok {
		return
	}

	if oki, err := h.auth.EqualsPopCode(updPassword.Email, updPassword.Code); err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	} else if !oki {
		c.Error(errs.CodeError.AddErr(err))
		return
	}

	if err := h.user.UpdatePassword([]byte(updPassword.NewPassword), updPassword.Email); err != nil {
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

// UpdateEmail godoc
// @Summary Update user's email
// @Security ApiKeyAuth
// @Description Updates user's email
// @Tags User
// @Param updEmail body dto.UpdateEmail true "New email with password"
// @Success 200 "Successfully Updated"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 404 {object} errs.MyError "User doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /user/email [patch]
func (h *Handler) updateEmail(c *gin.Context) {
	s, err := h.session.GetSession(c)
	if err != nil {
		c.Error(errs.UnAuthorized.AddErr(err))
		return
	}

	updEmail, ok := bind.FillStruct[dto.UpdateEmail](c)
	if !ok {
		return
	}

	if err = h.user.UpdateEmail([]byte(updEmail.Password), updEmail.NewEmail, s.ID); err != nil {
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

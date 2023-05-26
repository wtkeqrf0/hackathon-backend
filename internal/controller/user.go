package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"github.com/while-act/hackathon-backend/pkg/bind"
	"github.com/while-act/hackathon-backend/pkg/middleware/errs"
	"net/http"
)

// GetMe godoc
// @Summary Get detail data about the user by session
// @Security ApiKeyAuth
// @Description Returns detail information about me
// @Tags Session
// @Success 200 {object} dao.Me "Info about session"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /auth/session [get]
func (h *Handler) getMe(c *gin.Context) {
	s, err := h.session.GetSession(c)
	if err != nil {
		return
	}

	user, err := h.user.FindUserByID(s.ID)
	if err != nil {
		c.Error(errs.ValidError.AddErr(err))
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
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user [patch]
func (h *Handler) updateMe(c *gin.Context) {
	s, err := h.session.GetSession(c)
	if err != nil {
		return
	}

	updFields, err := bind.FillStructJSON[dto.UpdateUser](c)
	if err != nil {
		c.Error(errs.ValidError.AddErr(err))
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
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 500 {object} errs.MyError
// @Router /user/password [patch]
func (h *Handler) updatePassword(c *gin.Context) {
	updPassword, err := bind.FillStructJSON[dto.UpdatePassword](c)
	if err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}

	if ok, err := h.auth.EqualsPopCode(updPassword.Email, updPassword.Code); err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	} else if !ok {
		c.Error(errs.CodeError.AddErr(err))
		return
	}

	if err = h.user.UpdatePassword([]byte(updPassword.NewPassword), updPassword.Email); err != nil {
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
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user/email [patch]
func (h *Handler) updateEmail(c *gin.Context) {
	s, err := h.session.GetSession(c)
	if err != nil {
		return
	}

	updEmail, err := bind.FillStructJSON[dto.UpdateEmail](c)
	if err != nil {
		c.Error(errs.ValidError.AddErr(err))
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

// GetHistory godoc
// @Summary Generate PDF file from user history
// @Description Returns PDF file got from user history
// @Security ApiKeyAuth
// @Tags User
// @Param company_name path string true "Unique company name from history"
// @Success 200 "PDF file"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user/{company_name} [get]
func (h *Handler) getHistory(c *gin.Context) {
	s, err := h.session.GetSession(c)
	if err != nil {
		return
	}

	_, err = h.user.GetOneHistory(c.Param("company_name"), s.ID)
	if err != nil {
		return
	}

	c.Status(http.StatusNotModified)
}

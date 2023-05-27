package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"github.com/while-act/hackathon-backend/internal/service"
	"github.com/while-act/hackathon-backend/pkg/bind"
	"github.com/while-act/hackathon-backend/pkg/middleware/errs"
	"net/http"
	"strconv"
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
	s := h.session.GetSession(c)
	if s == nil {
		return
	}

	user, err := h.user.FindUserByID(s.ID)
	if err != nil {
		c.Error(err)
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
	s := h.session.GetSession(c)
	if s == nil {
		return
	}

	updFields := bind.FillStructJSON[dto.UpdateUser](c)
	if updFields == nil {
		return
	}

	if err := h.user.UpdateUser(updFields, s.ID); err != nil {
		c.Error(err)
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
	updPassword := bind.FillStructJSON[dto.UpdatePassword](c)
	if updPassword == nil {
		return
	}

	if ok, err := h.auth.EqualsPopCode(updPassword.Email, updPassword.Code); err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	} else if !ok {
		c.Error(errs.CodeError.AddErr(err))
		return
	}

	if err := h.user.UpdatePassword([]byte(updPassword.NewPassword), updPassword.Email); err != nil {
		c.Error(err)
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
	s := h.session.GetSession(c)
	if s == nil {
		return
	}

	updEmail := bind.FillStructJSON[dto.UpdateEmail](c)
	if updEmail == nil {
		return
	}

	if err := h.user.UpdateEmail([]byte(updEmail.Password), updEmail.NewEmail, s.ID); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

// GetHistory godoc
// @Summary Generate PDF file from user history
// @Description Returns PDF file got from user history
// @Security ApiKeyAuth
// @Tags User
// @Param history_id path string true "Unique id from history"
// @Success 200 "PDF file"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user/{history_id} [get]
func (h *Handler) getHistory(c *gin.Context) {
	s := h.session.GetSession(c)
	if s == nil {
		return
	}

	historyIdStr := c.Param("history_id")

	err := binding.Validator.ValidateStruct(dto.HistoryId{Id: historyIdStr})
	if err != nil {
		c.Error(err)
		return
	}
	historyId, _ := strconv.Atoi(historyIdStr)

	history, err := h.user.GetOneHistory(historyId, s.ID)
	if err != nil {
		c.Error(err)
		return
	}

	p := service.Params{
		IndustryBranch:      history.IndustryBranch,
		OrganizationType:    history.OrganizationType,
		FullTimeEmployers:   history.FullTimeEmployees,
		LandArea:            history.LandArea,
		Total:               0,
		Staff:               0,
		RentalProperty:      0,
		Taxes:               0,
		Services:            0,
		StaffNum:            0,
		MinStaffMaintenance: 0,
		MaxStaffMaintenance: 0,
		MinPensionInsurance: 0,
		MaxPensionInsurance: 0,
		MinHealthInsurance:  0,
		MaxHealthInsurance:  0,
	}

	err = h.pdf.GeneratePDF(c.Writer, p)
	if err != nil {
		c.Error(errs.PDFError.AddErr(err))
		return
	}

	c.Header("Content-Disposition", "attachment; filename=result.pdf")
	c.Header("Content-Type", "application/pdf")
	c.Status(http.StatusOK)
}

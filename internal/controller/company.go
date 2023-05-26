package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"github.com/while-act/hackathon-backend/pkg/bind"
	"github.com/while-act/hackathon-backend/pkg/middleware/errs"
	"net/http"
)

// GetMyCompany godoc
// @Summary Get data about company by session
// @Security ApiKeyAuth
// @Description Returns information about company by session
// @Tags Company
// @Success 200 {object} dao.Company "Info about company"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /company [get]
func (h *Handler) getMyCompany(c *gin.Context) {
	s, err := h.session.GetSession(c)
	if err != nil {
		return
	}

	user, err := h.user.FindUserByID(s.ID)
	if err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}

	company, err := h.company.GetCompanyDTO(user.CompanyID)
	if err != nil {
		c.Error(errs.NoSuchCompany.AddErr(err))
		return
	}

	c.JSON(http.StatusOK, company)
}

// UpdateCompany godoc
// @Summary Update data about company
// @Description Updates information about company by INN
// @Tags Company
// @Security ApiKeyAuth
// @Param updCompany body dto.UpdateCompany true "Company"
// @Success 200 "OK"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /company [patch]
func (h *Handler) updateCompany(c *gin.Context) {
	s, err := h.session.GetSession(c)
	if err != nil {
		return
	}

	updCompany, err := bind.FillStructJSON[dto.UpdateCompany](c)
	if err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}

	user, err := h.user.FindUserByID(s.ID)
	if err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}

	if err = h.company.UpdateCompany(updCompany, user.CompanyID); err != nil {
		switch {
		case ent.IsValidationError(err):
			c.Error(errs.ValidError.AddErr(err))
		default:
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.Status(http.StatusOK)
}

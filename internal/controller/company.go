package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"github.com/while-act/hackathon-backend/pkg/bind"
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
	s := h.session.GetSession(c)
	if s == nil {
		return
	}

	user, err := h.user.FindUserByID(s.ID)
	if err != nil {
		c.Error(err)
		return
	}

	company, err := h.company.GetCompanyDTO(user.CompanyID)
	if err != nil {
		c.Error(err)
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
	s := h.session.GetSession(c)
	if s == nil {
		return
	}

	updCompany := bind.FillStructJSON[dto.UpdateCompany](c)
	if updCompany == nil {
		return
	}

	user, err := h.user.FindUserByID(s.ID)
	if err != nil {
		c.Error(err)
		return
	}

	if err = h.company.UpdateCompany(updCompany, user.CompanyID); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

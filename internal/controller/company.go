package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
	"net/http"
)

// GetMyCompany godoc
// @Summary Get data about company by jwt
// @Security ApiKeyAuth
// @Description Returns information about company by jwt
// @Tags Company
// @Success 200 {object} dao.Company "Info about company"
// @Failure 404 {object} errs.MyError "Company doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /company [get]
func (h *Handler) getMyCompany(c *gin.Context) {
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

	company, err := h.company.GetCompanyDTO(user.INN)
	if err != nil {
		if myErr, ok := err.(errs.MyError); ok {
			c.Error(myErr)
		} else {
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.JSON(http.StatusOK, company)
}

// GetCompany godoc
// @Summary Get data about company inn
// @Description Returns information about company by INN
// @Tags Company
// @Param inn path string true "company's inn"
// @Success 200 {object} dao.Company "Info about company"
// @Failure 404 {object} errs.MyError "Company doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /company/{inn} [get]
func (h *Handler) getCompany(c *gin.Context) {
	inn := c.Param("inn")

	company, err := h.company.GetCompanyDTO(inn)
	if err != nil {
		if myErr, ok := err.(errs.MyError); ok {
			c.Error(myErr)
		} else {
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.JSON(http.StatusOK, company)
}

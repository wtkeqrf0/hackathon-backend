package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"github.com/while-act/hackathon-backend/internal/service"
	"github.com/while-act/hackathon-backend/pkg/bind"
	"github.com/while-act/hackathon-backend/pkg/middleware/errs"
	"net/http"
)

// SaveCalcData godoc
// @Summary Save calc data to history
// @Security ApiKeyAuth
// @Description Saves given values to user's history
// @Tags Calc
// @Param from body dto.History true "Completed application form"
// @Success 200 "OK"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /calc/save [post]
func (h *Handler) saveCalcData(c *gin.Context) {
	s := h.session.GetSession(c)
	if s == nil {
		return
	}

	history := bind.FillStructJSON[dto.History](c)
	if history == nil {
		return
	}
	id, err := h.business.GetBusiness(history.BusinessActivity)
	if err != nil {
		c.Error(err)
		return
	}

	if err = h.history.CreateHistory(history, id, s.ID); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

// CalcData godoc
// @Summary Generate PDF from body
// @Description Returns PDF file, gotten from body
// @Tags Calc
// @Param from body dto.History true "Completed application form"
// @Success 200 "PDF file"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 500 {object} errs.MyError
// @Router /calc [post]
func (h *Handler) calcData(c *gin.Context) {
	history := bind.FillStructJSON[dto.History](c)
	if history == nil {
		return
	}

	dist, err := h.district.GetDistrict(history.DistrictTitle)
	if err != nil {
		c.Error(err)
		return
	}
	var tax float64
	if history.AccountingSupport {
		tax, err = h.tax.GetTax(history.TaxationSystemOperations, history.OperationsType)
		if err != nil {
			c.Error(err)
			return
		}
	}

	p := service.Params{
		Other:             history.Other,
		IndustryBranch:    history.IndustryBranch,
		OrganizationType:  history.OrganizationLegal,
		FullTimeEmployers: history.FullTimeEmployees,
		LandArea:          history.LandArea,
	}

	wageFund := float64(history.FullTimeEmployees) * history.AvgSalary * 12
	p.InsurancePayment = wageFund * 0.3
	p.IncomeTax = wageFund * 0.13
	p.WageFund = wageFund + p.InsurancePayment + p.IncomeTax
	p.SocialInsurance = wageFund * 30
	p.Total += p.WageFund + p.SocialInsurance + p.InsurancePayment + p.IncomeTax

	if history.IsBuy {
		p.LandValue = history.LandArea * dist.AvgCadastralVal
		p.LandValueMin = history.ConstructionFacilitiesArea * 80
		p.LandValueMax = history.ConstructionFacilitiesArea * 120
	} else {
		p.LandValue = history.LandArea * dist.AvgCadastralVal * 0.003
	}
	p.Total += p.LandValue
	for _, v := range history.Equipment {
		p.Equipment += v.Price
	}
	p.Total += p.Equipment
	if history.AccountingSupport {
		p.Taxes = tax + (0.5 * float64(history.FullTimeEmployees))
		p.Total += p.Taxes
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

// GetIndustryInfo godoc
// @Summary Get data about industry
// @Description Returns detail information about industry
// @Tags Calc
// @Param industry path string true "Industry Branch"
// @Success 200 {object} dao.Industry "Info about industry"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 500 {object} errs.MyError
// @Router /calc/{industry} [get]
func (h *Handler) getIndustryInfo(c *gin.Context) {
	industry, err := h.industry.GetIndustry(c.Param("industry"))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, industry)
}

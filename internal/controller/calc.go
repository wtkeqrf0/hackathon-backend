package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
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
// @Router /calc [post]
func (h *Handler) saveCalcData(c *gin.Context) {
	s, err := h.session.GetSession(c)
	if err != nil {
		return
	}

	history := bind.FillStructJSON[dto.History](c)
	if history == nil {
		return
	}

	if err = h.history.CreateHistory(history, s.ID); err != nil {
		c.Error(errs.ServerError.AddErr(err))
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

	c.Status(http.StatusNotModified)
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
		c.Error(errs.NoSuchIndustry.AddErr(err))
		return
	}

	c.JSON(http.StatusOK, industry)
}

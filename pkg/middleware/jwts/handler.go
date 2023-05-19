package jwts

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
	"github.com/wtkeqrf0/while.act/pkg/conf"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
)

var cfg = conf.GetConfig()

// RequireAuth authorizes the user
func (a Auth) RequireAuth(c *gin.Context) {
	t := new(dto.Token)
	if err := c.ShouldBindHeader(t); err != nil {
		c.Error(errs.UnAuthorized.AddErr(err))
		return
	}

	id, err := a.ValidateJWT(t.Authorization)
	if err != nil {
		c.Error(errs.UnAuthorized.AddErr(err))
		return
	}

	c.Set("id", id)
	c.Next()
}

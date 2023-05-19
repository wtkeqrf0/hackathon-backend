package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
	"github.com/wtkeqrf0/while.act/pkg/bind"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// SignUp godoc
// @Summary Sign in by password
// @Description Compare the user's password with an existing user's password. If it matches, create session of the user. If the user does not exist, create new user
// @Param SignUp body dto.SignUp true "User's email, password, firstName, lastName, inn"
// @Tags Authorization
// @Success 200 {string} string "user's session"
// @Failure 400 {object} errs.MyError "Data is not valid"
// @Failure 500 {object} errs.MyError
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	auth, ok := bind.FillStruct[dto.SignUp](c)
	if !ok {
		return
	}

	user, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		company := new(ent.Company)

		company, err = h.company.GetCompany(auth.Company.INN)
		if err != nil {

			company, err = h.company.CreateCompany(
				auth.Company.INN,
				auth.Company.Name,
				auth.Company.Website)

			if err != nil {
				c.Error(errs.ServerError.AddErr(err))
				return
			}
		}

		user, err = h.auth.CreateUserWithPassword(auth, company)

		if err != nil {
			c.Error(errs.ServerError.AddErr(err))
			return
		}
	}

	if err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(auth.Password)); err != nil {
		c.Error(errs.PasswordError.AddErr(err))
		return
	}

	t, err := h.jwt.GenerateJWT(user.ID)
	if err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": t,
	})
}

// SignIn godoc
// @Summary Sign in by password
// @Description Compare the user's password with an existing user's password. If it matches, create session of the user. If the user does not exist, create new user
// @Param SignIn body dto.SignIn true "User's email, password"
// @Tags Authorization
// @Success 200 {string} string "user's session"
// @Failure 400 {object} errs.MyError "Data is not valid"
// @Failure 500 {object} errs.MyError
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	auth, ok := bind.FillStruct[dto.SignIn](c)
	if !ok {
		return
	}

	user, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		c.Error(errs.NoSuchUser.AddErr(err))
		return
	}

	if err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(auth.Password)); err != nil {
		c.Error(errs.PasswordError.AddErr(err))
		return
	}

	t, err := h.jwt.GenerateJWT(user.ID)
	if err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": t,
	})
}

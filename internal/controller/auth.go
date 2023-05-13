package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
	"github.com/wtkeqrf0/while.act/pkg/bind"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// SignInWithPassword godoc
// @Summary Sign in by password
// @Description Compare the user's password with an existing user's password. If it matches, create session of the user. If the user does not exist, create new user
// @Param EmailWithPassword body dto.EmailWithPassword true "User's email, password"
// @Tags Authorization
// @Success 200 {header} string "user's session"
// @Failure 400 {object} errs.MyError "Data is not valid"
// @Failure 404 {object} errs.MyError "Password is not registered for this account"
// @Failure 500 {object} errs.MyError
// @Router /auth/sign-in [post]
func (h Handler) signIn(c *gin.Context) {
	auth, ok := bind.FillStruct[dto.EmailWithPassword](c)
	if !ok {
		return
	}

	customer, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		customer, err = h.auth.CreateUserWithPassword(auth)

		if err != nil {
			c.Error(errs.ServerError.AddErr(err))
			return
		}

	} else if customer.PasswordHash == nil {
		c.Error(errs.PasswordNotFound)
		return
	}

	if err = bcrypt.CompareHashAndPassword(*customer.PasswordHash, []byte(auth.Password)); err != nil {
		c.Error(errs.PasswordError.AddErr(err))
		return
	}

	t, err := h.jwt.GenerateJWT(customer.ID)
	if err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": t,
	})
}

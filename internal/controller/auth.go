package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"github.com/while-act/hackathon-backend/pkg/bind"
	"github.com/while-act/hackathon-backend/pkg/middleware/errs"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
)

// SignUp godoc
// @Summary Sign up by password
// @Description Compare the user's password with an existing user's password. If it matches, create session of the user. If the user does not exist, create new user
// @Param SignUp body dto.SignUp true "User's email, password, firstName, lastName, inn"
// @Tags Auth
// @Success 200 {string} string "user's session"
// @Failure 400 {object} errs.MyError "Data is not valid"
// @Failure 404 {object} errs.MyError "User doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	auth := bind.FillStructJSON[dto.SignUp](c)
	if auth == nil {
		return
	}

	user, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		company := new(ent.Company)

		company, err = h.company.CreateCompany(
			auth.Company.INN,
			auth.Company.Name,
			auth.Company.Website,
		)

		if err != nil {
			c.Error(errs.ValidError.AddErr(err))
			return
		}

		user, err = h.auth.CreateUserWithPassword(auth, company)
		if err != nil {
			c.Error(errs.ValidError.AddErr(err))
			return
		}
	}

	if err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(auth.Password)); err != nil {
		c.Error(errs.PasswordError.AddErr(err))
		return
	}

	h.session.SetNewCookie(user.ID, c)

	c.Status(http.StatusOK)
}

// SignIn godoc
// @Summary Sign in by password
// @Description Compare the user's password with an existing user's password. If it matches, create session of this user
// @Param SignIn body dto.SignIn true "User's email, password"
// @Tags Auth
// @Success 200 {string} string "user's session"
// @Failure 400 {object} errs.MyError "Data is not valid"
// @Failure 404 {object} errs.MyError "User doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	auth := bind.FillStructJSON[dto.SignIn](c)
	if auth == nil {
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

	h.session.SetNewCookie(user.ID, c)

	c.Status(http.StatusOK)
}

// SendCodeToEmail godoc
// @Summary Send code to specified email
// @Description Generates secret code and sends it to specified email
// @Param SignIn body dto.Email true "User's email"
// @Tags Auth
// @Success 200 {string} string "user's session"
// @Failure 400 {object} errs.MyError "Data is not valid"
// @Failure 500 {object} errs.MyError
// @Router /email/send-code [post]
func (h *Handler) sendCodeToEmail(c *gin.Context) {
	to := bind.FillStructJSON[dto.Email](c)
	if to == nil {
		return
	}

	code := generateSecretCode()
	if err := h.auth.SetCodes(to.Email, code); err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	if err := h.mail.SendEmail("Verify email for while.act account", code, cfg.Email.From, to.Email); err != nil {
		c.Error(errs.EmailError.AddErr(err))
		return
	}

	c.Status(http.StatusOK)
}

// SignOut godoc
// @Summary Delete cookie session
// @Description Get cookie and delete them from db
// @Tags Session
// @Success 200 {string} string "deleted"
// @Failure 500 {object} errs.MyError
// @Router /auth/session [delete]
func (h *Handler) signOut(c *gin.Context) {
	h.session.PopCookie(c)
	c.Status(http.StatusOK)
}

// generateSecretCode for email auth
func generateSecretCode() string {
	b := make([]rune, 5)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

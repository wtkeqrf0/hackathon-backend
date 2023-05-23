package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/internal/controller/dao"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
	"github.com/wtkeqrf0/while.act/pkg/conf"
	"time"
)

var (
	chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	cfg   = conf.GetConfig()
)

// UserService interacts with the users table
type UserService interface {
	FindUserByID(id int) (*dao.Me, error)

	UpdateUser(updateUser dto.UpdateUser, id int) error
	UpdatePassword(newPassword []byte, email string) error
	UpdateEmail(password []byte, newEmail string, id int) error

	ContainsKeys(keys ...string) (int64, error)
	SetVariable(key string, value any, exp time.Duration) error
}

type CompanyService interface {
	CreateCompany(inn string, name, website *string) (*ent.Company, error)
	GetCompany(id int) (*ent.Company, error)
	GetCompanyDTO(id int) (*dao.Company, error)
	UpdateCompany(updateCompany dto.UpdateCompany, id int) error
}

type AuthService interface {
	CreateUserWithPassword(auth dto.SignUp, company *ent.Company) (*ent.User, error)
	AuthUserByEmail(email string) (*ent.User, error)

	EqualsPopCode(email string, code string) (bool, error)
	SetCodes(key string, value ...any) error
}

type AuthMiddleware interface {
	RequireSession(c *gin.Context)
	MaybeSession(c *gin.Context)
	GenerateSession(id int, ip, userAgent string) (string, error)
	SetNewCookie(id int, userAgent string, c *gin.Context)
	GetSession(c *gin.Context) (*dao.Session, error)
	PopCookie(c *gin.Context)
}

type MailSender interface {
	SendEmail(subj, body, from string, to ...string) error
}

type Handler struct {
	user    UserService
	company CompanyService
	auth    AuthService
	session AuthMiddleware
	mail    MailSender
}

func NewHandler(user UserService, company CompanyService, auth AuthService, session AuthMiddleware, mail MailSender) *Handler {
	return &Handler{user: user, company: company, auth: auth, session: session, mail: mail}
}

func (h *Handler) InitRoutes(rg *gin.RouterGroup, mailSet bool) {

	docs := rg.Group("/docs")
	{
		docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	auth := rg.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

		session := auth.Group("/session")
		{
			session.GET("", h.session.RequireSession, h.getMe)
			session.DELETE("", h.signOut)
		}
	}

	user := rg.Group("/user")
	{
		user.PATCH("", h.session.RequireSession, h.updateMe)
		user.PATCH("/password", h.session.RequireSession, h.updatePassword)
		user.PATCH("/email", h.session.RequireSession, h.updateEmail)
	}

	company := rg.Group("/company")
	{
		company.GET("", h.session.RequireSession, h.getMyCompany)
		company.PATCH("", h.session.RequireSession, h.updateCompany)
	}

	if mailSet {
		email := rg.Group("/email")
		{
			email.POST("/send-code", h.sendCodeToEmail)
		}
	}
}

type ErrHandler interface {
	HandleErrors(c *gin.Context)
}

type QueryHandler interface {
	HandleQueries(c *gin.Context)
}

type Middlewares struct {
	erh ErrHandler
	qh  QueryHandler
}

func NewMiddleWares(erh ErrHandler, qh QueryHandler) *Middlewares {
	return &Middlewares{erh: erh, qh: qh}
}

func (m *Middlewares) InitGlobalMiddleWares(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}), m.qh.HandleQueries, gin.Recovery(), m.erh.HandleErrors)
}

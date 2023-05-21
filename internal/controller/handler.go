package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/internal/controller/dao"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
)

// UserService interacts with the users table
type UserService interface {
	FindUserByID(id int) (*dao.Me, error)

	UpdateUser(updateUser dto.UpdateUser, id int) error
	UpdatePassword(updPassword dto.UpdatePassword, id int) error
}

type CompanyService interface {
	CreateCompany(inn string, name, website *string) (*ent.Company, error)
	GetCompany(inn string) (*ent.Company, error)
	GetCompanyDTO(inn string) (*dao.Company, error)
	UpdateCompany(updateCompany dto.UpdateCompany, inn string) error
}

type AuthService interface {
	CreateUserWithPassword(auth dto.SignUp, company *ent.Company) (*ent.User, error)
	AuthUserByEmail(email string) (*ent.User, error)
}

type AuthMiddleware interface {
	ValidateJWT(token string) (int, error)
	GenerateJWT(id int) (string, error)
	GetUserId(c *gin.Context) (int, bool)
	RequireAuth(c *gin.Context)
}

type Handler struct {
	user    UserService
	company CompanyService
	auth    AuthService
	jwt     AuthMiddleware
}

func NewHandler(user UserService, company CompanyService, auth AuthService, jwt AuthMiddleware) *Handler {
	return &Handler{user: user, company: company, auth: auth, jwt: jwt}
}

func (h *Handler) InitRoutes(rg *gin.RouterGroup) {

	docs := rg.Group("/docs")
	{
		docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	auth := rg.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

		session := rg.Group("/session")
		{
			session.GET("", h.jwt.RequireAuth, h.getMe)
		}
	}

	user := rg.Group("/user")
	{
		user.PATCH("", h.jwt.RequireAuth, h.updateMe)
		user.PATCH("/password", h.jwt.RequireAuth, h.updatePassword)
	}

	company := rg.Group("/company")
	{
		company.GET("/:inn", h.getCompany)
		company.GET("", h.jwt.RequireAuth, h.getMyCompany)
		company.PATCH("", h.jwt.RequireAuth, h.updateCompany)
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
	r.Use(m.qh.HandleQueries, gin.Recovery(), m.erh.HandleErrors)
}

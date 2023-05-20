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
}

type CompanyService interface {
	CreateCompany(inn string, name, website *string) (*ent.Company, error)
	GetCompany(inn string) (*ent.Company, error)
	GetCompanyDTO(inn string) (*dao.Company, error)
	UpdateCompany(inn string, name, website *string) error
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

type ErrHandler interface {
	HandleErrors(c *gin.Context)
}

type QueryHandler interface {
	HandleQueries(c *gin.Context)
}

type Handler struct {
	user    UserService
	company CompanyService
	jwt     AuthMiddleware
	auth    AuthService
	erh     ErrHandler
	qh      QueryHandler
}

func NewHandler(user UserService, company CompanyService, jwt AuthMiddleware, auth AuthService, erh ErrHandler, qh QueryHandler) *Handler {
	return &Handler{user: user, company: company, jwt: jwt, auth: auth, erh: erh, qh: qh}
}

func (h *Handler) InitRoutes(r *gin.Engine) {
	r.Use(h.qh.HandleQueries, gin.Recovery(), h.erh.HandleErrors)
	api := r.Group("/api")

	docs := api.Group("/docs")
	{
		docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

		session := api.Group("/session")
		{
			session.GET("", h.jwt.RequireAuth, h.getMe)
		}
	}

	user := api.Group("/user")
	{
		user.PATCH("", h.jwt.RequireAuth, h.updateMe)
	}

	company := api.Group("/company")
	{
		company.GET("/:inn", h.getCompany)
		company.GET("", h.jwt.RequireAuth, h.getMyCompany)
	}
}

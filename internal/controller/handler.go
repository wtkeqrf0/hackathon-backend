package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
	"github.com/wtkeqrf0/while.act/pkg/conf"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
	"github.com/wtkeqrf0/while.act/pkg/middleware/logger"
)

var cfg = conf.GetConfig()

// UserService interacts with the users table
type UserService interface {
	FindUserByID(id int) (*ent.User, error)
}

type AuthService interface {
	CreateUserWithPassword(auth dto.EmailWithPassword) (*ent.User, error)
	AuthUserByEmail(email string) (*ent.User, error)
}

type AuthMiddleware interface {
	ValidateJWT(token string) (int, error)
	GenerateJWT(id int) (string, error)
	GetUserId(c *gin.Context) (int, bool)
	RequireAuth(c *gin.Context)
}

type Handler struct {
	users UserService
	jwt   AuthMiddleware
	auth  AuthService
}

func NewHandler(users UserService, jwt AuthMiddleware, auth AuthService) *Handler {
	return &Handler{users: users, jwt: jwt, auth: auth}
}

func (h Handler) InitRoutes(r *gin.Engine) {
	r.Use(logger.QueryLogging, gin.Recovery(), errs.ErrorHandler)
	api := r.Group("/api")

	docs := api.Group("/docs")
	{
		docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	auth := api.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.GET("", h.jwt.RequireAuth, h.getMe)
	}
}

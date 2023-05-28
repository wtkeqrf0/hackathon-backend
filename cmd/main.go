package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/while-act/hackathon-backend/docs"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/internal/controller"
	"github.com/while-act/hackathon-backend/internal/repo/postgres"
	redisRepo "github.com/while-act/hackathon-backend/internal/repo/redis"
	"github.com/while-act/hackathon-backend/internal/service"
	"github.com/while-act/hackathon-backend/pkg/bind"
	"github.com/while-act/hackathon-backend/pkg/client/email"
	"github.com/while-act/hackathon-backend/pkg/client/postgresql"
	redisInit "github.com/while-act/hackathon-backend/pkg/client/redis"
	"github.com/while-act/hackathon-backend/pkg/conf"
	"github.com/while-act/hackathon-backend/pkg/middleware/errs"
	"github.com/while-act/hackathon-backend/pkg/middleware/logger"
	"github.com/while-act/hackathon-backend/pkg/middleware/sessions"
	"net/http"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006/01/02 15:32:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "status",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyMsg:   "message",
		},
	})
	logrus.SetReportCaller(true)

	// global validator of incoming values
	binding.Validator = bind.NewValid(validator.New())
}

// @title While.act API
// @version 1.0
// @description It's an API interacting with While.act using Golang
// @accept application/json
// @produce application/json
// @schemes http
// @BasePath /api

// @contact.name Contact us
// @contact.url https://github.com/while-act/hackathon-backend/issues/new/choose
// @contact.email  matvey-sizov@mail.ru

// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name session_id
// @host 37.230.195.26:3000

// @sessions.docs.description Authorization, registration and authentication
func main() {
	cfg := conf.GetConfig()
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", cfg.Listen.DomainName, cfg.Listen.Port)

	pClient, rClient, mailClient := getClients(cfg)

	h := initHandler(pClient, rClient, mailClient, cfg.TemplatePath)
	m := initMiddlewares()

	r := gin.New()

	m.InitGlobalMiddleWares(r)
	h.InitRoutes(r.Group(cfg.Listen.MainPath), mailClient != nil)

	run(cfg.Listen.Port, r, pClient, rClient, mailClient)
}

// run the Server with graceful shutdown
func run(port int, r *gin.Engine, pClient *ent.Client, rClient *redis.Client, mailClient *smtp.Client) {
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Fatalf("error occurred while running http server")
		}
	}()
	logrus.Infof("Server Started On Port %d", port)

	<-quit

	logrus.Info("Server Shutting Down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("Server Shutdown Failed")
	}

	if err := rClient.Close(); err != nil {
		logrus.WithError(err).Fatal("Redis Connection Shutdown Failed")
	}

	if err := pClient.Close(); err != nil {
		logrus.WithError(err).Fatal("PostgreSQL Connection Shutdown Failed")
	}

	if err := mailClient.Quit(); err != nil {
		logrus.WithError(err).Fatal("Email Connection Shutdown Failed")
	}

	logrus.Info("Server Exited Properly")
}

func getClients(cfg *conf.Config) (*ent.Client, *redis.Client, *smtp.Client) {
	pClient := postgresql.Open(cfg.DB.Postgres.Username, cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.DBName)

	rClient := redisInit.Open(cfg.DB.Redis.Host, cfg.DB.Redis.Port, cfg.DB.Redis.DbId)

	mailClient := email.Open(cfg.Email.User, cfg.Email.Password, cfg.Email.Host, cfg.Email.Port)

	return pClient, rClient, mailClient
}

func initHandler(pClient *ent.Client, rClient *redis.Client, mailClient *smtp.Client, tPath string) *controller.Handler {
	pUser := postgres.NewUserStorage(pClient.User)
	pComp := postgres.NewCompanyStorage(pClient.Company)
	pHist := postgres.NewHistoryStorage(pClient.History)
	pInd := postgres.NewIndustryStorage(pClient.Industry)
	pBus := postgres.NewBusinessStorage(pClient.BusinessActivity)
	rConn := redisRepo.NewRClient(rClient)

	auth := service.NewAuthService(pUser, rConn)
	history := service.NewHistoryService(pHist)
	industry := service.NewIndustryService(pInd)
	business := service.NewBusinessService(pBus)
	user := service.NewUserService(pUser, rConn)
	pdf := service.NewPDF(tPath)
	mail := service.NewEmailSender(mailClient)
	company := service.NewCompanyService(pComp)

	return controller.NewHandler(
		user,
		company,
		auth,
		history,
		industry,
		business,
		pdf,
		sessions.NewAuth(auth),
		mail,
	)
}

func initMiddlewares() *controller.Middlewares {
	return controller.NewMiddleWares(
		errs.NewErrHandler(logrus.New()),
		logger.NewQueryHandler(logrus.New()),
	)
}

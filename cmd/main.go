package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/wtkeqrf0/while.act/docs"
	_ "github.com/wtkeqrf0/while.act/docs"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/internal/controller"
	"github.com/wtkeqrf0/while.act/internal/repo/postgres"
	"github.com/wtkeqrf0/while.act/internal/service"
	"github.com/wtkeqrf0/while.act/pkg/bind"
	"github.com/wtkeqrf0/while.act/pkg/client/postgresql"
	"github.com/wtkeqrf0/while.act/pkg/conf"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
	"github.com/wtkeqrf0/while.act/pkg/middleware/jwts"
	"github.com/wtkeqrf0/while.act/pkg/middleware/logger"
	"io"
	"net/http"
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
// @name Authorization
// @host 68.183.76.225:3000

// @sessions.docs.description Authorization, registration and authentication
func main() {
	cfg := conf.GetConfig()
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", cfg.Listen.Host, cfg.Listen.Port)

	out := getLogsOut(cfg.LogsPath)

	pClient := postgresql.Open(cfg.DB.Postgres.Username, cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.DBName)

	h := initHandler(pClient)
	m := initMiddlewares(out)

	r := gin.New()
	m.InitGlobalMiddleWares(r)
	h.InitRoutes(r.Group(cfg.Listen.MainPath))

	run(cfg.Listen.Port, r, pClient, logrus.New())
}

// run the Server with graceful shutdown
func run(port int, r *gin.Engine, pClient *ent.Client, l *logrus.Logger) {
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.InfoLevel)
	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006/01/02 15:32:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "status",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyMsg:   "message",
		},
	})
	l.SetReportCaller(true)

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
			l.WithError(err).Fatalf("error occurred while running http server")
		}
	}()
	l.Infof("Server Started On Port %d", port)

	<-quit

	l.Info("Server Shutting Down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		l.WithError(err).Fatal("Server Shutdown Failed")
	}

	if err := pClient.Close(); err != nil {
		l.WithError(err).Fatal("PostgreSQL Connection Shutdown Failed")
	}

	l.Info("Server Exited Properly")
}

func getLogsOut(s string) io.Writer {
	if s != "cons" {
		file, err := os.OpenFile(s, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			logrus.WithError(err).Fatalf("can't write logs to %s", s)
		}
		writer := bufio.NewWriter(file)
		logrus.SetOutput(writer)
		return writer
	}
	return os.Stdout
}

func initHandler(pClient *ent.Client) *controller.Handler {
	pConn := postgres.NewUserStorage(pClient.User)

	user := service.NewUserService(pConn)
	company := service.NewCompanyService(postgres.NewCompanyStorage(pClient.Company))
	auth := service.NewAuthService(pConn)

	return controller.NewHandler(
		user,
		company,
		auth,
		jwts.NewAuth(auth),
	)
}

func initMiddlewares(out io.Writer) *controller.Middlewares {
	return controller.NewMiddleWares(
		errs.NewErrHandler(logrus.New(), out),
		logger.NewQueryHandler(logrus.New(), out),
	)
}

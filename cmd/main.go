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

// @title You Together API
// @version 1.0
// @description It's an API interacting with You Together using Golang
// @accept application/json
// @produce application/json
// @schemes http

// @host :3000
// @BasePath /api

// @sessions.docs.description Authorization, registration and authentication
func main() {
	cfg := conf.GetConfig()

	out := getLogsOut(cfg.LogsPath)

	pClient := postgresql.Open(cfg.DB.Postgres.Username, cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.DBName)

	pConn := postgres.NewUserStorage(pClient.User)
	auth := service.NewAuthService(pConn)

	h := controller.NewHandler(
		service.NewUserService(pConn),
		jwts.NewAuth(auth),
		auth,
		errs.NewErrHandler(logrus.New(), out),
		logger.NewQueryHandler(logrus.New(), out),
	)

	r := gin.New()
	h.InitRoutes(r)

	Run(fmt.Sprintf(":%d", cfg.Listen.Port), r, pClient, logrus.New())
}

// Run the Server with graceful shutdown
func Run(port string, r *gin.Engine, pClient *ent.Client, l *logrus.Logger) {
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
		Addr:           port,
		Handler:        r,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	docs.SwaggerInfo.Host = port

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.WithError(err).Fatalf("error occurred while running http server")
		}
	}()
	l.Infof("Server Started On Port %s", port)

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

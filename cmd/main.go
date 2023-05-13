package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	_ "github.com/wtkeqrf0/while.act/docs"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/internal/controller"
	"github.com/wtkeqrf0/while.act/internal/repo/postgres"
	"github.com/wtkeqrf0/while.act/internal/service"
	"github.com/wtkeqrf0/while.act/pkg/bind"
	"github.com/wtkeqrf0/while.act/pkg/client/postgresql"
	"github.com/wtkeqrf0/while.act/pkg/conf"
	"github.com/wtkeqrf0/while.act/pkg/middleware/jwts"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//TODO import docs file

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

// @host localhost:3000
// @BasePath /api

// @sessions.docs.description Authorization, registration and authentication
func main() {
	cfg := conf.GetConfig()

	pClient := postgresql.Open(cfg.DB.Postgres.Username, cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.DBName)

	pConn := postgres.NewUserStorage(pClient.User)
	auth := service.NewAuthService(pConn)

	h := controller.NewHandler(
		service.NewUserService(pConn),
		jwts.NewAuth(auth),
		auth,
	)

	r := gin.New()
	h.InitRoutes(r)

	Run(cfg.Listen.Port, r, pClient)
}

// Run the Server with graceful shutdown
func Run(port int, r *gin.Engine, pClient *ent.Client) {
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

	if err := pClient.Close(); err != nil {
		logrus.WithError(err).Fatal("PostgreSQL Connection Shutdown Failed")
	}

	logrus.Info("Server Exited Properly")
}

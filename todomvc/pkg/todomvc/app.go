package todomvc

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"todomvc/pkg/todomvc/conf"
	"todomvc/pkg/todomvc/model"
	"todomvc/pkg/todomvc/router"
)

var App *Application

const (
	ModeDebug   string = "debug"
	ModeRelease string = "release"
	ModeTest    string = "test"
)

type Application struct {
	Conf   *conf.Config
	Router *gin.Engine
	DB     *gorm.DB
}

func New(cfg *conf.Config) *Application {
	app := new(Application)

	gin.SetMode(gin.ReleaseMode)
	runMode := viper.GetString("app.run_mode")
	if runMode == ModeDebug {
		gin.SetMode(ModeDebug)
	}

	app.DB = model.Init(runMode)
	model.MigrateDB()

	app.Router = router.Router()

	return app
}

func (a *Application) Run() {
	server := &http.Server{
		Addr:    viper.GetString("app.addr"),
		Handler: a.Router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

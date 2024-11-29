package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ApiServer(log *zap.Logger, port, name string, app *gin.Engine) {
	srv := &http.Server{Addr: ":" + port, Handler: app}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("can't run service", zap.Error(err))
		}
	}()
	log.Info(name + " initiated at port " + port)

	// gracefully shutdown ------------------------------------------------------------------------
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown " + name + " service")

	cts, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(cts); err != nil {
		log.Error("can't shutdown "+name+" service", zap.Error(err))
	}

	log.Info(name + " service exiting")

	log.Info("Running cleanup tasks...")
}

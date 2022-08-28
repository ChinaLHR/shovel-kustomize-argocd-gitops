package main

import (
	"context"
	"github.com/chinalhr/go-stream"
	"github.com/chinalhr/go-stream/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	//Used for kubernetes readinessProbe and livenessProbe
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, Response{
			Code: 200,
			Msg:  "OK",
			Data: nil,
		})
	})

	router.GET("/env", func(c *gin.Context) {
		envKeys := c.Query("keys")
		data := stream.OfSlice(strings.Split(envKeys, ",")).
			Map(func(e types.T) (r types.R) {
				return map[string]string{e.(string): os.Getenv(e.(string))}
			}).
			ToSlice()
		c.JSON(200, Response{
			Code: 200,
			Msg:  "OK",
			Data: data,
		})
	})

	server := &http.Server{
		Addr:    ":9099",
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}

package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simpleRouter/core/conf"
	"time"
)

func InitGinServer() *http.Server {
	serverConf, err := conf.InitServerConf()
	if err == nil && serverConf != nil {
		router := gin.Default()
		//router.GET("/", func(c *gin.Context) {
		//    time.Sleep(5 * time.Second)
		//    c.String(http.StatusOK, "Welcome Gin Server")
		//})

		server := &http.Server{
			Addr:    "localhost:" + serverConf.ServerPort,
			Handler: router,
		}
		if serverConf.GraceShutDown {
			quit := make(chan os.Signal)
			signal.Notify(quit, os.Interrupt)
			go func() {
				<-quit
				log.Println("receive interrupt signal")
				if err := server.Close(); err != nil {
					log.Fatal("Server Close:", err)
				}
			}()
		}
		rl := serverConf.Rl
		if rl.Enable {
			rateLimit := ratelimit.New(rl.Rps)
			router.Use(leakBucket(rateLimit))
		}
		return server
	}
	return nil
}

func leakBucket(limit ratelimit.Limiter) gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		log.Printf("%v", now.Sub(prev))
		prev = now
	}
}

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
	serverConf := conf.InitServerConf()
	if serverConf != nil {
		router := gin.Default()
		//router.GET("/", func(c *gin.Context) {
		//    time.Sleep(5 * time.Second)
		//    c.String(http.StatusOK, "Welcome Gin Server")
		//})

		server := &http.Server{
			Addr:    "localhost:" + serverConf.ServerPort,
			Handler: router,
		}
		router.Any("/:path", handlerRequest)
		//router.POST("/:path", handlerRequest)
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

func handlerRequest(context *gin.Context) {
	req := context.Request
	path := req.URL.Path
	//filter.FilterRequest(context, path)
	//writeResponse(context, response)
	conf.FilterRequest(context, path)
	return
}

package server

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simpleRouter/core/conf"
	"simpleRouter/core/filter"
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
		router.GET("/:path", handlerRequest)
		router.POST("/:path", handlerRequest)
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
	deltegate := filter.GetDelegate(path)
	response := deltegate.Fc.Apply(context)
	writeResponse(context, response)
	return
}

func writeResponse(context *gin.Context, response *http.Response) {
	for k, vv := range response.Header {
		for _, v := range vv {
			context.Header(k, v)
		}
	}
	defer response.Body.Close()
	bufio.NewReader(response.Body).WriteTo(context.Writer)
}

//func(c *gin.Context) handler(){
//	// step 1: resolve proxy address, change scheme and host in requets
//	req := c.Request
//	proxy, err := url.Parse(getLoadBalanceAddr())
//	if err != nil {
//		log.Printf("error in parse addr: %v", err)
//		c.String(500, "error")
//		return
//	}
//	req.URL.Scheme = proxy.Scheme
//	req.URL.Host = proxy.Host
//
//	// step 2: use http.Transport to do request to real server.
//	transport := http.DefaultTransport
//	resp, err := transport.RoundTrip(req)
//	if err != nil {
//		log.Printf("error in roundtrip: %v", err)
//		c.String(500, "error")
//		return
//	}
//
//	// step 3: return real server response to upstream.
//	for k, vv := range resp.Header {
//		for _, v := range vv {
//			c.Header(k, v)
//		}
//	}
//	defer resp.Body.Close()
//	bufio.NewReader(resp.Body).WriteTo(c.Writer)
//	return
//}

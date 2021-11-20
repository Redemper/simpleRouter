package component

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/url"
	"simpleRouter/core/discover"
)

type Router struct {
	Name      string `gorm:"name",yaml:"name"`
	OriginUri string `gorm:"origin_uri",yaml:"origin-uri"`
	TargetUri string `gorm:"target_uri",yaml:"taget-uri"`
	Order     int    `gorm:"order",yaml:"order"`
	Enabled   bool   `gorm:"enabled",yaml:"enabled"`
	Filters   string `gorm:"filters",yaml:"filters"`
	gorm.Model
}

func (r *Router) RTrip() func(context *gin.Context) {
	return func(context *gin.Context) {
		uri := discover.GetTagetUriByOriginUri(r.OriginUri)
		req := context.Request
		proxy, err := url.Parse(uri)
		if err != nil {
			log.Printf("error in parse addr: %v", err)
			context.String(500, "error")
			return
		}
		req.URL.Scheme = proxy.Scheme
		req.URL.Host = proxy.Host
		transport := http.DefaultTransport
		resp, err := transport.RoundTrip(req)
		if err != nil {
			log.Printf("error in roundtrip: %v", err)
			context.String(500, "error")
			return
		}
		for k, vv := range resp.Header {
			for _, v := range vv {
				context.Header(k, v)
			}
		}
		defer resp.Body.Close()
		bufio.NewReader(resp.Body).WriteTo(context.Writer)
	}
}

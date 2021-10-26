package filter

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"simpleRouter/core/route/config"
)

type routerFilter struct {
	r *config.Router
}

func (rf *routerFilter) Apply(context *gin.Context) {
	req := context.Request
	uri := rf.r.TargetUri
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

func initRouterFilter(r *config.Router) *routerFilter {
	rf := new(routerFilter)
	rf.r = r
	return rf
}

//func (rf *routerFilter) Apply(context gin.Context) *http.Response{
//    	req := c.Request
//    	proxy, err := url.Parse(getLoadBalanceAddr())
//    	if err != nil {
//    		log.Printf("error in parse addr: %v", err)
//    		c.String(500, "error")
//    		return
//    	}
//    	req.URL.Scheme = proxy.Scheme
//    	req.URL.Host = proxy.Host
//
//    	// step 2: use http.Transport to do request to real server.
//    	transport := http.DefaultTransport
//    	resp, err := transport.RoundTrip(req)
//    	if err != nil {
//    		log.Printf("error in roundtrip: %v", err)
//    		c.String(500, "error")
//    		return
//    	}
//
//    	// step 3: return real server response to upstream.
//    	for k, vv := range resp.Header {
//    		for _, v := range vv {
//    			c.Header(k, v)
//    		}
//    	}
//    	defer resp.Body.Close()
//    	bufio.NewReader(resp.Body).WriteTo(c.Writer)
//    	return
//}

func (rf *routerFilter) NextFilter() FilterChan {
	return nil
}

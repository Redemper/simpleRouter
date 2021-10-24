package filter

import (
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"time"
)

type TimeWatchFilter struct {
	nextFilter FilterChan
}

func (tw *TimeWatchFilter) Ordered() int {
	return math.MaxInt8
}

var _ FilterChan = (*TimeWatchFilter)(nil)

func (tw *TimeWatchFilter) Apply(context *gin.Context) *http.Response {
	start := time.Now()
	log.Println("start filte , start time is ", start.Format("2006-01-02 15:04:05"))
	filter := tw.nextFilter
	if nil != filter {
		response := filter.Apply(context)
		if response.Status == "200" {
			log.Println("apply done")
		}
		log.Println("time used. ", time.Now().Sub(start))
		return response
	}
	return nil
}

func (tw *TimeWatchFilter) NextFilter() FilterChan {
	return tw.nextFilter
}

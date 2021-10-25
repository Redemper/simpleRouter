package filter

import (
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"time"
)

type timeWatchFilter struct {
	nextFilter FilterChan
}

func (tw *timeWatchFilter) Ordered() int {
	return math.MaxInt8
}

var _ FilterChan = (*timeWatchFilter)(nil)

func (tw *timeWatchFilter) Apply(context *gin.Context) {
	start := time.Now()
	log.Println("start filte , start time is ", start.Format("2006-01-02 15:04:05"))
	filter := tw.nextFilter
	if nil != filter {
		filter.Apply(context)
		log.Println("time used. ", time.Now().Sub(start))
	}
}

func (tw *timeWatchFilter) NextFilter() FilterChan {
	return tw.nextFilter
}

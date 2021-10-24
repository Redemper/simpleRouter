package filter

import (
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"time"
)

type TimeWatchFilter struct {
	nextFilter FilterChan
}

func (tw *TimeWatchFilter) Ordered() int {
	return math.MaxInt8
}

var _ FilterChan = (*TimeWatchFilter)(nil)

func (tw *TimeWatchFilter) Apply(context *gin.Context) {
	start := time.Now()
	log.Println("start filte , start time is ", start.Format("2006-01-02 15:04:05"))
	filter := tw.nextFilter
	if nil != filter {
		filter.Apply(context)
		log.Println("time used. ", time.Now().Sub(start))
	}
}

func (tw *TimeWatchFilter) NextFilter() FilterChan {
	return tw.nextFilter
}

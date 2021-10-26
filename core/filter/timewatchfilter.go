package filter

import (
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"time"
)

type timeWatchFilter struct {
}

func (tw *timeWatchFilter) Name() string {
	return "timeWatch"
}

func (tw *timeWatchFilter) Ordered() int {
	return math.MaxInt8
}

var _ Filter = (*timeWatchFilter)(nil)

func (tw *timeWatchFilter) Apply(context *gin.Context) {
	start := time.Now()
	log.Println("start filte , start time is ", start.Format("2006-01-02 15:04:05"))
	log.Println("time used. ", time.Now().Sub(start))
	//filter := tw.Nf
	//if nil != filter {
	//	filter.Apply(context)
	//}
}

//func (tw *timeWatchFilter) nextFilter() Filter {
//	return tw.Nf
//}

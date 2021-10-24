package filter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Filter interface {
	Filter(context gin.Context) *http.Response
}

type Ordered interface {
	GetOrder() int
}

type Chan interface {
	Apply(context gin.Context, filter Filter) *http.Response
}

package filter

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRouterFilter_Apply(t *testing.T) {
	http.HandleFunc("/test", SayHello)
	http.ListenAndServe("localhost:2258", nil)
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	fmt.Println("request path ", url)
	//name := url[len("/hello/"):]
	fmt.Fprintf(w, "hello,"+url)
}

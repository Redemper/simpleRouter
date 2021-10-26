package config

import (
	"encoding/json"
	"fmt"
	"simpleRouter/core/route"
	"testing"
)

func TestGetAllRouter(t *testing.T) {
	routers := route.GetAllRouter()
	for _, r := range routers {
		marshal, err := json.Marshal(r)
		if err == nil {
			fmt.Println(string(marshal))
		} else {
			fmt.Println("sry json fail")
		}
	}
}

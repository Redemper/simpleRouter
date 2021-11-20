package conf

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAllRouter(t *testing.T) {
	routers := GetAllRouter()
	for _, r := range routers {
		marshal, err := json.Marshal(r)
		if err == nil {
			fmt.Println(string(marshal))
		} else {
			fmt.Println("sry json fail")
		}
	}
}

func TestInitRouterFromYaml(t *testing.T) {
	routers := GetYamlRouters()
	fmt.Println(routers)
}

package conf

import (
	"fmt"
	"testing"
)

func TestInitServerConf(t *testing.T) {
	conf, err := InitServerConf()
	if err != nil {
		fmt.Println("errors ,", err)
	} else {
		fmt.Println(conf)
	}
}

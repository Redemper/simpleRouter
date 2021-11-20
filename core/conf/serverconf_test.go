package conf

import (
	"fmt"
	"testing"
)

func TestInitServerConf(t *testing.T) {
	conf := InitServerConf()
	fmt.Println(conf)

}

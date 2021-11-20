package conf

import (
	"fmt"
	"testing"
)

func TestReadYaml2(t *testing.T) {
	db := GetDbConf()
	fmt.Println(db)
	nacos := GetNacosConf()
	fmt.Println(nacos)
	server := GetServerConf()
	fmt.Println(server)
}

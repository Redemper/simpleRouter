package conf

import (
	"fmt"
	"testing"
)

func TestReadNacosConfFromYaml(t *testing.T) {
	yaml, err := ReadNacosConfFromYaml()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(yaml)
}

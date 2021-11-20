package conf

import (
	"fmt"
	"testing"
)

func TestReadNacosConfFromYaml(t *testing.T) {
	yaml := ReadNacosConfFromYaml()
	fmt.Println(yaml)
}

package conf

import (
	"fmt"
	"testing"
)

func TestReadNacosConfFromYaml(t *testing.T) {
	yaml := readNacosConfFromYaml()
	fmt.Println(yaml)
}

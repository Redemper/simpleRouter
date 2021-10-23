package conf

import (
	"log"
	"testing"
)

func TestReadYaml(t *testing.T) {
	_, err := GetDBFromYaml()
	if err != nil {
		log.Fatal("error occur", err)
	} else {
		log.Println("Test ReadYaml success")
	}
}

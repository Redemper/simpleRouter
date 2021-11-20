package conf

import (
	"log"
	"testing"
)

func TestGetDB(t *testing.T) {
	_, err := GetDB()
	if err != nil {
		log.Fatal("error occur", err)
	} else {
		log.Println("Test GetDB success")
	}
}

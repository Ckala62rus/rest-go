package main

import (
	"log"

	"github.com/Ckala62rus/rest-go"
)

func main() {
	srv := new(rest.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}

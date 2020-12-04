package main

import (
	"log"

	"github.com/vivaldy22/simpleRestApiLA/config"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Println(err)
		return
	}
	router := config.CreateRouter()
	config.InitRouters(db, router)
	config.RunServer(router)
}

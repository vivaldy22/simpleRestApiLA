package main

import "github.com/vivaldy22/simpleRestApiLA/config"

func main() {
	db := config.InitDB()
	router := config.CreateRouter()
	config.InitRouters(db, router)
	config.RunServer(router)
}

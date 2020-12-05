package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/simpleRestApiLA/tools/myViper"
)

func CreateRouter() *mux.Router {
	return mux.NewRouter()
}

func RunServer(r *mux.Router) {
	host := myViper.GetEnv("API_HOST", "")
	port := myViper.GetEnv("API_PORT", "8080")

	log.Printf("Starting Web Server at %v port: %v", host, port)
	fmt.Println("Available endpoints:")
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		if len(methods) != 0 {
			fmt.Println(path, methods)
		}
		return nil
	})
	fmt.Println()
	err := http.ListenAndServe(host+":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}

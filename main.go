package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	routes "github.com/bauidch/hyres-api/routers"
)

// Define HTTP request routes
func main() {
	router := routes.NewRouter()
	log.Printf("Server started")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Content-Range"})
	originsOk := handlers.AllowedOrigins([]string{"http://hyres-app.bauid.ch:3000", "*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	exposeOK := handlers.ExposedHeaders([]string{"Content-Range"})

	if err := http.ListenAndServe(":8080",  handlers.CORS(originsOk, headersOk, methodsOk, exposeOK)(router)); err != nil {
		log.Fatal(err)
	}

}

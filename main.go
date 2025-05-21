package main

import (
	"log"
	"net/http"

	"katalis.ai-home-test/middleware"
	"katalis.ai-home-test/routes"
)

func main() {
	go middleware.AsyncLogger()
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	mux.Handle("/swagger.yaml", http.FileServer(http.Dir("docs")))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("docs/swagger-ui"))))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

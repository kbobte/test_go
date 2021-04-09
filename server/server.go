package server

import (
	"fmt"
	"net/http"
	"todos/config"
	"todos/router"

	gohandlers "github.com/gorilla/handlers"
)

// Start server
func Start() {
	fmt.Println("Server running:" + config.ServerAddress)

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))
	// Load mux routes
	// Start server
	r := router.GetRouter()
	server := &http.Server{Addr: config.ServerAddress, Handler: ch(r)}
	server.ListenAndServe()
}

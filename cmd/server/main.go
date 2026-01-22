package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"parkping/internal/db"
	"parkping/internal/routes"

	_ "parkping/docs"
)

// @title ParkPing API
// @version 1.0
// @description Contact vehicle owners via Fastag or vehicle number
// @termsOfService https://parkping.io/terms

// @contact.name ParkPing Support
// @contact.email support@parkping.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Connect to DB (Postgres / SQLite in-memory for dev)
	database := db.Connect()

	// Create router
	r := mux.NewRouter()

	// Register all routes
	routes.RegisterRoutes(r, database)

	// Get port from env or default to 8080
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

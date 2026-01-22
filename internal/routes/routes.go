package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"parkping/internal/middleware"

	"parkping/internal/auth"
	"parkping/internal/contact"
	"parkping/internal/vehicle"

	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterRoutes wires all routes to the given router and DB
func RegisterRoutes(r *mux.Router, db *gorm.DB) {

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Auth routes
	authHandler := auth.NewHandler(db)
	r.HandleFunc("/auth/send-otp", authHandler.SendOTP).Methods("POST")
	r.HandleFunc("/auth/verify-otp", authHandler.VerifyOTP).Methods("POST")

	// Vehicle routes
	vehicleHandler := vehicle.NewHandler(db)

	r.HandleFunc("/vehicle", middleware.Auth(vehicleHandler.Create)).Methods("POST")
	r.HandleFunc("/vehicle/lookup", middleware.Auth(vehicleHandler.Lookup)).Methods("GET")

	// Contact routes
	contactHandler := contact.NewHandler(db)
	r.HandleFunc("/contact/call", middleware.Auth(contactHandler.Call)).Methods("POST")
	r.HandleFunc("/contact/message", middleware.Auth(contactHandler.Message)).Methods("POST")
}

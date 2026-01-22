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
	v1 := r.PathPrefix("/api/v1").Subrouter()
	authHandler := auth.NewHandler(db)
	v1.HandleFunc("/auth/send-otp", authHandler.SendOTP).Methods("POST")
	v1.HandleFunc("/auth/verify-otp", authHandler.VerifyOTP).Methods("POST")

	// Vehicle routes
	vehicleHandler := vehicle.NewHandler(db)

	v1.HandleFunc("/vehicle", middleware.Auth(vehicleHandler.Create)).Methods("POST")
	v1.HandleFunc("/vehicle/lookup", middleware.Auth(vehicleHandler.Lookup)).Methods("GET")
	v1.HandleFunc("/vehicle/qr", vehicleHandler.QR).Methods("GET")
	v1.HandleFunc("/scan/image", vehicleHandler.ScanImage).Methods("POST")

	// Contact routes
	contactHandler := contact.NewHandler(db)
	v1.HandleFunc("/contact/call", middleware.Auth(contactHandler.Call)).Methods("POST")
	v1.HandleFunc("/contact/message", middleware.Auth(contactHandler.Message)).Methods("POST")
}

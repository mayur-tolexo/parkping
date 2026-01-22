package vehicle

import (
	"encoding/json"
	"net/http"

	"parkping/internal/middleware"
	"parkping/internal/model"

	"gorm.io/gorm"
)

// Handler holds the DB connection
type Handler struct {
	DB *gorm.DB
}

// NewHandler creates a new Vehicle handler with DB
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

// CreateVehicleRequest represents API payload to register a vehicle
type CreateVehicleRequest struct {
	VehicleNumber string `json:"vehicle_number"`
	VehicleType   string `json:"vehicle_type"`
	FastagNumber  string `json:"fastag_number"`
}

// VehicleLookupResponse represents lookup response
type VehicleLookupResponse struct {
	VehicleNumber   string `json:"vehicle_number"`
	CallsEnabled    bool   `json:"calls_enabled"`
	MessagesEnabled bool   `json:"messages_enabled"`
}

// Create godoc
// @Summary Register vehicle
// @Tags Vehicle
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body CreateVehicleRequest true "Vehicle details"
// @Success 200 {object} map[string]string
// @Failure 401 {string} string
// @Router /vehicle [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(uint)

	var req CreateVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	vehicle := model.Vehicle{
		UserID:        userID,
		VehicleNumber: req.VehicleNumber,
		VehicleType:   req.VehicleType,
		FastagNumber:  req.FastagNumber,
	}

	if err := h.DB.Create(&vehicle).Error; err != nil {
		http.Error(w, "vehicle creation failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userID,
		"fastag":  vehicle.FastagNumber,
	})
}

// Lookup godoc
// @Summary Lookup vehicle by Fastag
// @Description Returns vehicle details and contact permissions using Fastag number
// @Tags Vehicle
// @Security BearerAuth
// @Produce json
// @Param fastag query string true "Fastag number"
// @Success 200 {object} VehicleLookupResponse
// @Failure 400 {string} string "missing fastag"
// @Failure 404 {string} string "vehicle not found"
// @Failure 401 {string} string "unauthorized"
// @Router /vehicle/lookup [get]
func (h *Handler) Lookup(w http.ResponseWriter, r *http.Request) {
	fastag := r.URL.Query().Get("fastag")
	if fastag == "" {
		http.Error(w, "missing fastag", http.StatusBadRequest)
		return
	}

	var vehicle model.Vehicle
	if err := h.DB.
		Select("vehicle_number", "calls_enabled", "messages_enabled").
		Where("fastag_number = ?", fastag).
		First(&vehicle).Error; err != nil {

		http.Error(w, "vehicle not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(VehicleLookupResponse{
		VehicleNumber:   vehicle.VehicleNumber,
		CallsEnabled:    vehicle.CallsEnabled,
		MessagesEnabled: vehicle.MessagesEnabled,
	})
}

package contact

import (
	"encoding/json"
	"net/http"

	"parkping/internal/model"

	"gorm.io/gorm"
)

// Handler holds DB connection
type Handler struct {
	DB *gorm.DB
}

// NewHandler creates a new Contact handler with DB
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

// ContactRequest represents request to call/message owner
type ContactRequest struct {
	Fastag  string `json:"fastag"`
	Type    string `json:"type"` // call | message
	Message string `json:"message"`
}

// Call handles masked call to vehicle owner
// @Summary Call vehicle owner (masked call)
// @Description Initiates a masked call to the vehicle owner using Fastag
// @Tags Contact
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body ContactRequest true "Fastag contact request"
// @Success 200 {string} string "call initiated"
// @Failure 400 {string} string "invalid request"
// @Failure 403 {string} string "calls disabled"
// @Failure 404 {string} string "vehicle not found"
// @Failure 401 {string} string "unauthorized"
// @Router /contact/call [post]
func (h *Handler) Call(w http.ResponseWriter, r *http.Request) {
	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	var vehicle model.Vehicle
	if err := h.DB.Joins("User").
		Where("fastag_number = ?", req.Fastag).
		First(&vehicle).Error; err != nil {

		http.Error(w, "vehicle not found", http.StatusNotFound)
		return
	}

	if !vehicle.CallsEnabled {
		http.Error(w, "calls disabled", http.StatusForbidden)
		return
	}

	// Integrate Exotel masked call here
	w.WriteHeader(http.StatusOK)
}

// Message handles sending message to vehicle owner
// @Summary Send message to vehicle owner
// @Description Sends a predefined or custom message to the vehicle owner using Fastag
// @Tags Contact
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body ContactRequest true "Fastag message request"
// @Success 200 {string} string "message sent"
// @Failure 400 {string} string "invalid request"
// @Failure 403 {string} string "messages disabled"
// @Failure 404 {string} string "vehicle not found"
// @Failure 401 {string} string "unauthorized"
// @Router /contact/message [post]
func (h *Handler) Message(w http.ResponseWriter, r *http.Request) {
	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	var vehicle model.Vehicle
	if err := h.DB.Joins("User").
		Where("fastag_number = ?", req.Fastag).
		First(&vehicle).Error; err != nil {

		http.Error(w, "vehicle not found", http.StatusNotFound)
		return
	}

	if !vehicle.MessagesEnabled {
		http.Error(w, "messages disabled", http.StatusForbidden)
		return
	}

	// Send predefined SMS here
	w.WriteHeader(http.StatusOK)
}

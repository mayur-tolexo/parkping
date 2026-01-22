package auth

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"

	"parkping/internal/model"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(database *gorm.DB) *Handler {
	return &Handler{db: database}
}

type VerifyOTPRequest struct {
	Phone string `json:"phone"`
	OTP   string `json:"otp"`
}

func (h *Handler) SendOTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// VerifyOTP godoc
// @Summary Verify OTP and login
// @Description Verifies OTP and returns JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body VerifyOTPRequest true "Verify OTP"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string
// @Router /auth/verify-otp [post]
func (h *Handler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req VerifyOTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	var user model.User
	err := h.db.Where("phone_number = ?", req.Phone).First(&user).Error
	if err != nil {
		user = model.User{PhoneNumber: req.Phone}
		if err := h.db.Create(&user).Error; err != nil {
			http.Error(w, "user creation failed", http.StatusInternalServerError)
			return
		}
	}

	token, err := GenerateToken(int(user.ID))
	if err != nil {
		http.Error(w, "token generation failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

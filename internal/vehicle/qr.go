package vehicle

import (
	"encoding/json"
	"net/http"
	"parkping/internal/model"

	"github.com/skip2/go-qrcode"
	dqrcode "github.com/tuotoo/qrcode"
)

func (h *Handler) QR(w http.ResponseWriter, r *http.Request) {
	fastag := r.URL.Query().Get("fastag")
	if fastag == "" {
		http.Error(w, "missing fastag", http.StatusBadRequest)
		return
	}

	png, err := qrcode.Encode(fastag, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "qr generation failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}

func (h *Handler) ScanImage(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "failed to read image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	qrMatrix, err := dqrcode.Decode(file)
	if err != nil {
		http.Error(w, "invalid QR code", http.StatusBadRequest)
		return
	}

	fastag := qrMatrix.Content

	var vehicle model.Vehicle
	if err := h.DB.Where("fastag_number = ?", fastag).First(&vehicle).Error; err != nil {
		http.Error(w, "vehicle not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"vehicle_number":   vehicle.VehicleNumber,
		"calls_enabled":    vehicle.CallsEnabled,
		"messages_enabled": vehicle.MessagesEnabled,
	})
}

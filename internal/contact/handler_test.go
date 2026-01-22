package contact

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"parkping/internal/auth"
	"parkping/internal/db"
	"parkping/internal/middleware"
	"parkping/internal/vehicle"

	"github.com/gorilla/mux"
)

func TestContactAPI(t *testing.T) {
	database := db.Connect()

	// First, create a vehicle for user_id 1
	vh := vehicle.NewHandler(database)
	vehicleData := map[string]string{
		"vehicle_number": "KA01XY9999",
		"vehicle_type":   "Car",
	}
	b, _ := json.Marshal(vehicleData)

	// Router for vehicle create
	r := mux.NewRouter()
	r.HandleFunc("/vehicle", middleware.Auth(vh.Create)).Methods("POST")

	// Generate test JWT
	token, _ := auth.GenerateTestToken(1)

	req := httptest.NewRequest("POST", "/vehicle", bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp map[string]string
	json.NewDecoder(w.Body).Decode(&resp)
	fastag := resp["fastag"]

	// -----------------------------
	// Now test contact call
	// -----------------------------
	ch := NewHandler(database)
	r2 := mux.NewRouter()
	r2.HandleFunc("/contact/call", middleware.Auth(ch.Call)).Methods("POST")
	r2.HandleFunc("/contact/message", middleware.Auth(ch.Message)).Methods("POST")

	// Call API
	callReq := map[string]string{
		"fastag": fastag,
		"type":   "call",
	}
	cb, _ := json.Marshal(callReq)
	req2 := httptest.NewRequest("POST", "/contact/call", bytes.NewReader(cb))
	req2.Header.Set("Authorization", "Bearer "+token)
	w2 := httptest.NewRecorder()
	r2.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w2.Code, w2.Body.String())
	}

	// Message API
	msgReq := map[string]string{
		"fastag":  fastag,
		"type":    "message",
		"message": "Please move your vehicle",
	}
	mb, _ := json.Marshal(msgReq)
	req3 := httptest.NewRequest("POST", "/contact/message", bytes.NewReader(mb))
	req3.Header.Set("Authorization", "Bearer "+token)
	w3 := httptest.NewRecorder()
	r2.ServeHTTP(w3, req3)

	if w3.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w3.Code, w3.Body.String())
	}
}

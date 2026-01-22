package vehicle

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"parkping/internal/auth"
	"parkping/internal/db"
	"parkping/internal/middleware"

	"github.com/gorilla/mux"
)

func TestVehicleAPI(t *testing.T) {
	// In-memory SQLite DB
	database := db.Connect()
	handler := NewHandler(database)

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/vehicle", middleware.Auth(handler.Create)).Methods("POST")
	r.HandleFunc("/vehicle/lookup", middleware.Auth(handler.Lookup)).Methods("GET")

	// Generate test JWT
	token, _ := auth.GenerateTestToken(1)

	// -----------------------------
	// 1️⃣ Test vehicle creation
	// -----------------------------
	body := map[string]string{
		"vehicle_number": "KA01AB1234",
		"vehicle_type":   "Car",
		"fastag_number":  "test",
	}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/vehicle", bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp map[string]string
	json.NewDecoder(w.Body).Decode(&resp)
	if resp["fastag"] == "" {
		t.Fatal("fastag not returned")
	}

	fastag := resp["fastag"]

	// -----------------------------
	// 2️⃣ Test vehicle lookup
	// -----------------------------
	req2 := httptest.NewRequest("GET", "/vehicle/lookup?fastag="+fastag, nil)
	req2.Header.Set("Authorization", "Bearer "+token)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w2.Code, w2.Body.String())
	}

	var lookupResp map[string]interface{}
	json.NewDecoder(w2.Body).Decode(&lookupResp)
	if lookupResp["vehicle_number"] != "KA01AB1234" {
		t.Fatalf("expected vehicle_number KA01AB1234, got %v", lookupResp["vehicle_number"])
	}
}

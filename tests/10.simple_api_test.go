package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"basics/api"
)

func TestPing(t *testing.T) {
	req := httptest.NewRequest("GET", "/api", nil)
	w := httptest.NewRecorder()

	api.Ping(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}

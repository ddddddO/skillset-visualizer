package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// ref: https://devblog.thebase.in/entry/2018/12/04/110000
func TestHealthCheckHandler(t *testing.T) {
	tests := []struct {
		method             string
		expectedStatusCode int
	}{
		{http.MethodGet, http.StatusOK},
		{http.MethodPost, http.StatusMethodNotAllowed},
		{http.MethodDelete, http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(tt.method, "/health", nil)

		healthCheckHandler(w, r)
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != tt.expectedStatusCode {
			t.Errorf("expected status code is %d, got %d", tt.expectedStatusCode, res.StatusCode)
		}
	}
}

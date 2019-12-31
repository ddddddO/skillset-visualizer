package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ref: https://devblog.thebase.in/entry/2018/12/04/110000
func TestHealthCheckHandler(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		expectedStatusCode int
	}{
		{"get", http.MethodGet, http.StatusOK},
		{"post", http.MethodPost, http.StatusMethodNotAllowed},
		{"delete", http.MethodDelete, http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(tt.method, "/health", nil)

		healthCheckHandler(w, r)
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != tt.expectedStatusCode {
			t.Errorf("%s pattern. expected status code is %d, got %d", tt.name, tt.expectedStatusCode, res.StatusCode)
		}
	}
}

func TestFetchGraphDataHandler(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		expectedStatusCode int
		expectedData       string
	}{
		{"get", http.MethodGet, http.StatusOK, `{"frontend":1,"backend":2,"database":3,"infrastructure":"4","network":"5","facilitation":4}`},
		{"post", http.MethodPost, http.StatusMethodNotAllowed, ""},
		{"delete", http.MethodDelete, http.StatusMethodNotAllowed, ""},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(tt.method, "/fetch", nil)

		fetchGraphDataHandler(w, r)
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != tt.expectedStatusCode {
			t.Errorf("%s pattern. expected status code is %d, got %d", tt.name, tt.expectedStatusCode, res.StatusCode)
		}

		// 正常にデータが返却されるパターン
		if tt.method == http.MethodGet {
			buf, _ := ioutil.ReadAll(res.Body)
			if tt.expectedData != string(buf) {
				t.Errorf("%s pattern. expected data is %s,\ngot %s",
					tt.name, tt.expectedData, string(buf))
			}
		}
	}
}

func TestPutGraphDataHandler(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		expectedStatusCode int
	}{
		{"options", http.MethodOptions, http.StatusOK},
		//{"put", http.MethodPut, http.StatusOK}, FIXME: 変更データを与えてみる
		{"get", http.MethodGet, http.StatusMethodNotAllowed},
		{"post", http.MethodPost, http.StatusMethodNotAllowed},
		{"delete", http.MethodDelete, http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(tt.method, "/put/1", nil)

		putGraphDataHandler(w, r)
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != tt.expectedStatusCode {
			t.Errorf("%s pattern. expected status code is %d, got %d", tt.name, tt.expectedStatusCode, res.StatusCode)
		}
	}
}

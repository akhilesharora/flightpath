package handler

import (
	"encoding/json"
	flight "github.com/akhilesharora/flightpath/internal/flightpath"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// TestHandleCalculate tests the handleCalculate handler function.
func TestHandleCalculate(t *testing.T) {
	// Create a new server and initialize its dependencies.
	s := &Server{
		flightSrv: &flight.Tracker{},
	}

	// Define the test cases.
	testCases := []struct {
		name     string
		body     string
		expected []string
		status   int
	}{
		{
			name:     "valid flights",
			body:     `{"flights":[["SFO","HKO"],["YYZ","SFO"],["YUL","YYZ"],["HKO","ORD"]]}`,
			expected: []string{"YUL", "ORD"},
			status:   http.StatusOK,
		},
		{
			name:     "invalid flights",
			body:     `{"flights":[["SFO","HKO"],["YYZ","SFO"],["YUL","YYZ"],["YUL","HKO"]]}`,
			expected: []string{},
			status:   http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new request and set the body.
			req, err := http.NewRequest("POST", "/calculate", strings.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}

			// Create a new response recorder.
			rr := httptest.NewRecorder()

			// Call the handler function.
			s.handleCalculate(rr, req)

			// Check the status code.
			if status := rr.Code; status != tc.status {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.status)
			}

			// Check the response body.
			var resp struct {
				Path []string `json:"path"`
			}
			if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(resp.Path, tc.expected) {
				t.Errorf("handler returned unexpected body: got %v want %v", resp.Path, tc.expected)
			}
		})
	}
}

package handler

import (
	"encoding/json"
	"net/http"
)

type Path struct {
	Path []string `json:"path"`
}

func (s *Server) handleCalculate(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req struct {
		Flights [][]string `json:"flights"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Sort the flights to determine the flight path
	path := s.flightSrv.FindItineraryStartAndEnd(req.Flights)

	w.Header().Set("Content-Type", "application/json")
	// Return the flight path in the response body as a JSON array
	if err := json.NewEncoder(w).Encode(&Path{Path: path}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

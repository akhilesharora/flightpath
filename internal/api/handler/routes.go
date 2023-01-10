package handler

import (
	"net/http"
	"time"

	flight "github.com/akhilesharora/flightpath/internal/tracker"
	"github.com/gorilla/mux"
)

//Server contains all components needed for this server
type Server struct {
	router    *mux.Router
	httpSrv   *http.Server
	flightSrv *flight.Tracker
}

func New() *Server {
	s := Server{
		router:    mux.NewRouter(),
		flightSrv: &flight.Tracker{},
	}

	// Accepts JSON input in the format:
	// { "flights": [["HEL","AMS"],["DEL", "HEL"],["LKO","DEL"]] }
	s.router.HandleFunc("/calculate", s.handleCalculate)
	return &s
}

//ListenAndServe starts listening for requests
func (s *Server) ListenAndServe(url string) error {
	s.httpSrv = &http.Server{
		Handler:      s.router,
		Addr:         url,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return s.httpSrv.ListenAndServe()
}

//Close the server
func (s *Server) Close() error {
	return s.httpSrv.Close()
}

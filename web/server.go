package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Trother555/portctl/app"
)

type Server struct {
	app app.App
}

func New(app app.App) *Server {
	return &Server{app: app}
}

func (s *Server) ReadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	portNumRaw := r.URL.Query().Get("portNum")
	port, err := strconv.Atoi(portNumRaw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	val, err := s.app.Read(int64(port))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["val"] = strconv.FormatInt(val, 10)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}

func (s *Server) WriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	portNumRaw := r.URL.Query().Get("portNum")
	port, err := strconv.Atoi(portNumRaw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	transactionIdRaw := r.URL.Query().Get("transactionId")
	transactionId, err := strconv.Atoi(transactionIdRaw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	valRaw := r.URL.Query().Get("val")
	val, err := strconv.Atoi(valRaw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.app.Write(int64(port), int64(transactionId), int64(val))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) ListenAndServe() error {
	mux := http.NewServeMux()

	mux.Handle("/read", http.HandlerFunc(s.ReadHandler))

	mux.Handle("/write", http.HandlerFunc(s.WriteHandler))

	return http.ListenAndServe(":8080", mux)
}

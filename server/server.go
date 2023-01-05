package server

import (
	"encoding/json"
	"log"
	"net/http"

	"Angular/api-rest/store"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type server struct {
	Router *chi.Mux
	Store  store.Store
}

func NewServerChi() *server {
	s := &server{
		Router: chi.NewRouter(),
	}
	s.Router.Use(middleware.Logger)
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
	s.routes()
	return s
}

func (s *server) respond(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v", err)
	}
}

func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

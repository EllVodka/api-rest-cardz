package server

import (
	"Angular/api-rest/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// cardz get all cardz
func (s *server) cardz(w http.ResponseWriter, r *http.Request) {
	c, err := s.Store.GetCardz()
	if err != nil {
		s.respond(w, r, "query failed :"+err.Error(), http.StatusInternalServerError)
		return
	}
	s.respond(w, r, c, http.StatusOK)
}

// cardzById cardz by id
func (s *server) cardzById(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Printf("Cannot parse id to int. err=%v", err)
		s.respond(w, r, nil, http.StatusBadRequest)
		return
	}
	c, err := s.Store.GetCardzById(id)
	if err != nil {
		s.respond(w, r, "query failed :"+err.Error(), http.StatusInternalServerError)
		return
	}
	s.respond(w, r, c, http.StatusOK)
}

func (s *server) deleteCardz(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Printf("Cannot parse id to int. err=%v", err)
		s.respond(w, r, nil, http.StatusBadRequest)
		return
	}

	s.respond(w, r, id, http.StatusOK)
}

func (s *server) createCardz(w http.ResponseWriter, r *http.Request) {
	c := models.Cardz{}

	if err := s.decode(w, r, &c); err != nil {
		s.respond(w, r, nil, http.StatusBadRequest)
		return
	}

	if err := s.Store.CreateCardz(c); err != nil {
		log.Printf("Query failed, err: %v", err)
		s.respond(w, r, nil, http.StatusInternalServerError)
		return
	}

	s.respond(w, r, nil, http.StatusOK)

}

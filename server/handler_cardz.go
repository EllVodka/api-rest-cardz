package server

import (
	"Angular/api-rest/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

//  @Summary      Show Cardz
//  @Description  get all Cardz
//  @Tags         Cardz
//  @Produce      json
//  @Success      200  {array}   models.Cardz
//  @Failure      404  {object}  models.Error
//  @Failure      500  {object}  models.Error
//  @Router       /cardz [get]
// cardz get all cardz
func (s *server) cardz(w http.ResponseWriter, r *http.Request) {
	c, err := s.Store.GetCardz()
	if err != nil {
		s.respond(w, r, "query failed :"+err.Error(), http.StatusInternalServerError)
		return
	}
	s.respond(w, r, c, http.StatusOK)
}

//  @Summary      Show Departments
//  @Description  Get specific Cardz
//  @Tags         Cardz
//  @Produce      json
//  @Param	      id path int true "id of a cardz"
//  @Success      200  {object}  models.Cardz
//  @Failure      400  {object}  models.Error
//  @Failure      404  {object}  models.Error
//  @Failure      500  {object}  models.Error
//  @Router       /cardz/{id} [get]
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

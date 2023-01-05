package server

import (
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *server) routes() {
	s.Router.Get("/", s.home)

	// Swagger
	s.Router.Get("/swagger/*", httpSwagger.WrapHandler)
	s.Router.Get("/cardz", s.cardz)
	s.Router.Get("/cardz/{id:[0-9]+}", s.cardzById)
	s.Router.Post("/create-cardz", s.createCardz)
	s.Router.Delete("/cardz/{id:[0-9]+}", s.deleteCardz)
}

func (s *server) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my incredible api !!")
}

package main

import (
	"Angular/api-rest/server"
	"Angular/api-rest/store"
	"fmt"
	"log"
	"net/http"
	"os"
)

// @title Api-Rest Cardz
// @version 0.0.1
// @description Api-Rest pour animer l'application cardz en angular

// @BasePath /
func main() {
	fmt.Println("Welcome to api rest")
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}

func run() error {

	srv := server.NewServerChi()

	srv.Store = &store.DbStore{}
	if err := srv.Store.Open(); err != nil {
		return err
	}
	defer srv.Store.Close()
	log.Printf("Serving HTTP on port 8080")
	if err := http.ListenAndServe(":8080", srv.Router); err != nil {
		return err
	}
	return nil
}

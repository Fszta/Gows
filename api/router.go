package api

import (
	"github.com/gorilla/mux"
)

func Routing() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/add", AddDag)
	router.HandleFunc("/list", ListDag)
	router.HandleFunc("/remove", RemoveDag)

	return router
}

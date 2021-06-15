package api

import (
	"github.com/gorilla/mux"
)

var (
	AddDagRoute     = "/dag/add"
	ListDagsRoute   = "/dag/list"
	RemoveDagRoute  = "/dag/remove"
	StopDagRoute    = "/dag/stop"
	TriggerDagRoute = "/dag/trigger"
	RestartDagRoute = "/dag/restart"
)

func Routing() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(AddDagRoute, AddDag)
	router.HandleFunc(ListDagsRoute, ListDag)
	router.HandleFunc(RemoveDagRoute, RemoveDag)
	router.HandleFunc(StopDagRoute, StopDag)
	router.HandleFunc(TriggerDagRoute, TriggerDag)
	router.HandleFunc(RestartDagRoute, RestartDag)

	return router
}

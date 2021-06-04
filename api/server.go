package api

import (
	"log"
	"net/http"
)

var port = "2128"

func Start() {
	router := Routing()
	handler := http.Handler(router)
	http.HandleFunc("/add", AddDag)
	http.HandleFunc("/list", ListDag)
	log.Printf("Server started at port %v", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, handler))
}

package api

import (
	"log"
	"net/http"
)

var port = "2128"

func Start() {
	router := Routing()
	handler := http.Handler(router)
	log.Printf("Server started at port %v", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, handler))
}

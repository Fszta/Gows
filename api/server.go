package api

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	router := Routing()
	handler := http.Handler(router)
	http.HandleFunc("/add", AddDag)
	http.HandleFunc("/list", ListDag)
	fmt.Println("Server started at port 2128")
	log.Fatal(http.ListenAndServe(":2128", handler))
}

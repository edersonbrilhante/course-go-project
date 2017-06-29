package main

import (
	"github.com/edersonbrilhante/go-course-project/api"
	"log"
	"net/http"
		"github.com/dimfeld/httptreemux"

)

func main() {
	addr := "127.0.0.1:8081"
	router := httptreemux.NewContextMux()
	router.Handler(http.MethodGet, "/music/:id", &api.GetMusicHandler{})
	router.Handler(http.MethodPut, "/music/:id", &api.UpdateMusicHandler{})
	router.Handler(http.MethodDelete, "/music/:id", &api.DeleteMusicHandler{})
	router.Handler(http.MethodPost, "/music/", &api.PostMusicHandler{})

	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))

	// execute
	// curl http://localhost:8081/cars/gol
	// curl -XPUT http://localhost:8081/cars/fusca -d'{"name": 1}'
}

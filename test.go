package main

import (
	"log"
	"net/http"

	handler "ac9/glad/services/sf_handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/account", handler.AccountHandler)
	router.HandleFunc("/course", handler.CourseHandler)
	router.HandleFunc("/product", handler.ProductHandler)
	router.HandleFunc("/timing", handler.TimingHandler)
	router.HandleFunc("/center", handler.CenterHandler)

	log.Println("now listening at port 4001")
	log.Println(http.ListenAndServe(":4001", router))
}

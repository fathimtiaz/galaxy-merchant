package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/fathimtiaz/galaxy-merchant/config"
	"github.com/fathimtiaz/galaxy-merchant/pkg/http/handlers"
)

func main() {
	config.LoadConfig()
	
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Home)
	router.HandleFunc("/post", handler.Result).Methods(http.MethodPost)
	http.Handle("/", router)

	fmt.Println("Running server on: ", config.CONF.Host)
	log.Fatal(http.ListenAndServe(config.CONF.Host, nil))
}

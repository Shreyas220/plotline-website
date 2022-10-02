package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"plotline-website/handler"
	"plotline-website/utils"

	handle "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var PORT = os.Getenv("PORT")

var portString = fmt.Sprintf(":%s", PORT)

// var portString = ":8080"

func main() {
	c := utils.InitializeGoogleApi()
	h := handler.NewHandler(c)

	r := mux.NewRouter()
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./frontend"))))

	r.HandleFunc("/getdirection", h.GetWithSDK).Methods("POST")
	r.HandleFunc("/getdirectionurl", h.GetWithUrl).Methods("POST")
	fmt.Println("Start listening on ", portString)

	log.Fatal(http.ListenAndServe(portString, handle.CORS(
		handle.AllowedHeaders([]string{"Content-Type", "Auth-Token", "token"}),
		handle.AllowedOrigins([]string{"*"}),
		handle.AllowCredentials(),
		handle.AllowedMethods([]string{"GET", "HEAD", "POST", "DELETE", "OPTIONS"}),
	)(r)))

	fmt.Println("stop listening")
}

package main

import (
"fmt"
"log"
"github.com/gorilla/mux"
"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleware(ProtectedEndpoint)).Methods("GET")

log.Println("Listen on port 8000")
log.Fatal(http.ListenAndServe(":8000", router))

}

func signup(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("successfully called signup" ))
}

func login(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("successfully called login" ))
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("successfully called ProtectedEndpoint" ))
}

func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc{
fmt.Println("TokenVerifyMiddleware Invoked")
	return nil
}

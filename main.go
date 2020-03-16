package main

import (
"fmt"
"log"
"github.com/gorilla/mux"
"github.com/lib/pq"
"net/http"
"database/sql"
)

type User struct {
	ID		int `json:"id"`
	Email		string `json:"email"`
	Password		string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

var db *sql.DB 

func main() {
	pgUrl, err := pq.ParseURL("postgres://ooioxunz:RT16WK...@drona.db.elephantsql.com:5432/ooioxunz")

	if err != nil{
		log.Fatal(err)
    }

	db, err = sql.Open("postgres", pgUrl)

    if err != nil{
		log.Fatal(err)
    }

	err = db.Ping()

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


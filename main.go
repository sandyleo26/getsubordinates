package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/kr/pretty"

	"github.com/gorilla/mux"
)

func postRoles(w http.ResponseWriter, r *http.Request) {
	var requestRoles []Role
	if err := json.NewDecoder(r.Body).Decode(&requestRoles); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	pretty.Println(requestRoles)
	setRoles(requestRoles)
	w.WriteHeader(http.StatusOK)
}

func postUsers(w http.ResponseWriter, r *http.Request) {
	var requestUsers []User
	if err := json.NewDecoder(r.Body).Decode(&requestUsers); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	pretty.Println(requestUsers)
	setUsers(requestUsers)
	w.WriteHeader(http.StatusOK)
}

func httpGetSubOrdinates(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	subordinates, err := getSubOrdinates(int(idInt))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(subordinates)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/roles", postRoles).Methods("POST")
	r.HandleFunc("/users", postUsers).Methods("POST")
	r.HandleFunc("/subordinates/{id}", httpGetSubOrdinates).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Listenning :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

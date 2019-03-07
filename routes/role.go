package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gmdmgithub/go_rest_api/models"
	"github.com/gorilla/mux"
)

var roles []models.Role

// InitRoles -
func InitRoles(router *mux.Router) {
	//read from db or something
	roles = models.ReadRoles()

	//init route
	router.HandleFunc("/goapi/roles", getRoles).Methods("GET")
	router.HandleFunc("/goapi/roles/{id:[0-9]+}", getRole).Methods("GET")
	router.HandleFunc("/goapi/roles", createRole).Methods("POST")
	router.HandleFunc("/goapi/roles/{id:[0-9]+}", updateRole).Methods("PUT")
	router.HandleFunc("/goapi/roles/{id:[0-9]+}", deleteRole).Methods("DELETE")
}

/**
 * Returns all the books
 * @param {w} http respose writer
 * @param {r} pointer to http request
 */
func getRoles(resw http.ResponseWriter, req *http.Request) {

	log.Printf("Get roles URI: %s %s", req.RequestURI, "getRoles")
	//fmt.Fprintf(w, "List of book will be serve soon ...")

	resw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resw).Encode(roles)
}

/**
 * Return a single book
 * @param {w} http respose writer
 * @param {r} pointer to http request
 */
func getRole(w http.ResponseWriter, r *http.Request) {
	log.Print("Get roles")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Println(params)

	for _, item := range roles {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid role number", http.StatusInternalServerError)
			return
		}
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

/**
 * Create a new book and add it to the list
 * @param {w} http respose writer
 * @param {r} pointer to http request
 */
func createRole(w http.ResponseWriter, r *http.Request) {
	log.Print("createRole")
	//fmt.Fprintf(w, "createRole ....")

	var role models.Role
	_ = json.NewDecoder(r.Body).Decode(&role)
	role.ID = rand.Intn(100000000) //simulate db id autoincrement

	roles = append(roles, role)

	json.NewEncoder(w).Encode(role)

}

/**
 * Update a particular book
 * @param {w} http respose writer
 * @param {r} pointer to http request
 */
func updateRole(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "updateRole ....")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Println(params)

	for index, item := range roles {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid role number", http.StatusInternalServerError)
			return
		}
		if item.ID == id {
			var role models.Role
			_ = json.NewDecoder(r.Body).Decode(&role)
			roles[index] = role
			break
		}
	}

	json.NewEncoder(w).Encode(roles)

}

/**
 * Delete a particular book
 * @param {w} http respose writer
 * @param {r} pointer to http request
 */
func deleteRole(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "deleteBook ....")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Println(params)

	for index, item := range roles {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid role number", http.StatusInternalServerError)
			return
		}
		if item.ID == id {
			roles = append(roles[:index], roles[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

// SayHiRole -
func SayHiRole() {
	fmt.Println("Hi tehere - role here!!!")
}

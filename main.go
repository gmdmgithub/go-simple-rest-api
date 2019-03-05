package main

// "github.com/gmdmgithub/go_rest_api/routes"

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gmdmgithub/go_rest_api/models"
	"github.com/gmdmgithub/go_rest_api/routes"
	"github.com/gorilla/mux"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hi there! Use:
		- /goapi/books - list of books
		- /goapi/books/id - get the book
		- /goapi/books - post a book
		- /goapi/book/id - put(update) the book
		- /goapi/book/id - delete the book`)
}

func forTesting(test, test2 string) string {
	res := test
	if len(test2) > 0 { //improvement after test scenario
		res += " " + test2
	}
	return res
}

func main() {
	//router initialization
	router := mux.NewRouter()

	routes.InitBooks()

	//Route handlers/endpots
	router.HandleFunc("/goapi/books", routes.GetBooks).Methods("GET")
	router.HandleFunc("/goapi/books/{id}", routes.GetBook).Methods("GET")
	router.HandleFunc("/goapi/books", routes.CreateBook).Methods("POST")
	router.HandleFunc("/goapi/books/{id}", routes.UpdateBook).Methods("PUT")
	router.HandleFunc("/goapi/books/{id}", routes.DeleteBook).Methods("DELETE")
	router.HandleFunc("/", sayHello).Methods("GET")

	routes.SayHi()
	models.RoleModel()

	port := 8001
	// Listen to port (8081) and handle requests - response is nil (null)
	message := "Server is running on port " + strconv.Itoa(port) + " ..."
	log.Println(message)
	log.Fatal(http.ListenAndServe(":8001", router))

}

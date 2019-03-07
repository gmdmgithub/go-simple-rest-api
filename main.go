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
	"github.com/joho/godotenv"
)

var myEnv map[string]string

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	Hi there, nice to see U! 
	For book use:
		- /goapi/books - list of books
		- /goapi/books/id - get the book
		- /goapi/books - post a book
		- /goapi/book/id - put(update) the book
		- /goapi/book/id - delete the book
	For roles use:
		- /goapi/roles - list of role
		- /goapi/roles/id - get the role
		- /goapi/roles - post a role
		- /goapi/role/id - put(update) the role
		- /goapi/role/id - delete the role
		`)
}

func forTesting(test, test2 string) string {
	res := test
	if len(test2) > 0 { //improvement after test scenario
		res += " " + test2
	}
	return res
}

func loadEnvData() {
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//test
	log.Println(myEnv["SECRET_KEY"])
}

func main() {
	loadEnvData()
	//router initialization
	router := mux.NewRouter()

	routes.InitBooks(router)
	routes.InitRoles(router)

	router.HandleFunc("/", sayHello).Methods("GET")

	routes.SayHiRole()
	routes.SayHi()
	models.SayModelRole()

	port := 8001
	message := "Server is running on port " + strconv.Itoa(port) + " ..."
	log.Println(message)

	// Listen to port (8081) and handle requests - response is nil (null)
	log.Fatal(http.ListenAndServe(":8001", router))

}

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

// type server struct {
// 	db     *someDatabase
// 	router *someRouter
// 	email  EmailSender
// }

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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("Logger url: %s query map is: %s ", r.RequestURI, r.URL.Query())
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
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
	//use logging middleware to log on all requests
	router.Use(loggingMiddleware)

	routes.SayHiRole()
	routes.SayHi()
	models.SayModelRole()

	port := 8001
	message := "Server is running on port " + strconv.Itoa(port) + " ..."
	log.Println(message)

	// Listen to port (8081) and handle requests - response is nil (null)
	log.Fatal(http.ListenAndServe(":8001", router))

}

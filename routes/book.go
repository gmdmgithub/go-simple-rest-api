package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book - model for book - Struct
type Book struct {
	ID     int     `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type Author struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

// InitBooks init a book rout
func InitBooks(router *mux.Router) {

	auth1 := Author{ID: "1", Firstname: "John", Lastname: "Doe"}
	auth2 := Author{ID: "1", Firstname: "Marin", Lastname: "Smith"}

	books = append(books, Book{ID: 1, Isbn: "442233", Title: "Nice title 1", Author: &auth1})
	books = append(books, Book{ID: 2, Isbn: "223344", Title: "Nice title 2", Author: &auth2})
	books = append(books, Book{ID: 3, Isbn: "556677", Title: "Nice title 3", Author: &auth1})
	books = append(books, Book{ID: 4, Isbn: "667788", Title: "Nice title 4", Author: &auth2})
	books = append(books, Book{ID: 5, Isbn: "778899", Title: "Nice title 5",
		Author: &Author{ID: "3", Firstname: "Adam", Lastname: "Frodo"}})

	router.HandleFunc("/goapi/books", getBooks).Methods("GET")
	router.HandleFunc("/goapi/books/{id:[0-9]+}", getBook).Methods("GET")
	router.HandleFunc("/goapi/books", createBook).Methods("POST")
	router.HandleFunc("/goapi/books/{id:[0-9]+}", updateBook).Methods("PUT")
	router.HandleFunc("/goapi/books/{id:[0-9]+}", deleteBook).Methods("DELETE")
}

/**
 * Returns all the books
 * @param {w} http respose writer
 * @param {r} pointer to http request
 */
func getBooks(resw http.ResponseWriter, req *http.Request) {
	log.Printf("Get books URI: %s", req.RequestURI)
	//fmt.Fprintf(w, "List of book will be serve soon ...")
	resw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resw).Encode(books)
}

/**
 * Return a single book
 * @param {w} http respose writer
 * @param {r} pointer to http request
 */
func getBook(w http.ResponseWriter, r *http.Request) {
	log.Print("Get book")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	log.Printf("Paramters in %v", params)

	for _, item := range books {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid book number", http.StatusInternalServerError)
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
func createBook(w http.ResponseWriter, r *http.Request) {
	log.Print("createBook")
	//fmt.Fprintf(w, "createBook ....")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = rand.Intn(100000000) //simulate db id autoincrement

	books = append(books, book)

	json.NewEncoder(w).Encode(books)

}

/**
 * Update a particular book
 * @param {w} http respose writer
 * @param {r} pointer to http request
 */
func updateBook(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "updateBook ....")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Println(params)

	for index, item := range books {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid book number", http.StatusInternalServerError)
			return
		}
		if item.ID == id {
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			books[index] = book
			break
		}
	}

	json.NewEncoder(w).Encode(books)

}

/**
 * Delete a particular book
 * @param {w} http respose writer
 * @param {r} pointer to http request
 */
func deleteBook(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "deleteBook ....")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Println(params)

	for index, item := range books {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid book number", http.StatusInternalServerError)
			return
		}
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func SayHi() {
	fmt.Println("Hi tehere!")
}

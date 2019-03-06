package models

import (
	"fmt"
)

//Role is popula bakery
type Role struct {
	ID          int    `json:"id"`
	ImageURL    string `json:"url"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

//initiate roles
var roles []Role

// RoleModel - gets the role model
func RoleModel() {
	roles = append(roles, Role{ID: 1, ImageURL: "https://google.com", Name: "John Doe", Ingredients: "Water, eggs, kefir"})
	roles = append(roles, Role{ID: 2, ImageURL: "https://google.com/name", Name: "Jon Blame", Ingredients: "Coffee, eggs, kefir"})
	fmt.Printf("Hi there! %v \n", roles)

}

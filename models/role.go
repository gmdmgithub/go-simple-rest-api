package models

import (
	"fmt"
)

//Role - is popular bakery
type Role struct {
	ID          int    `json:"id"`
	ImageURL    string `json:"url"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

//initiate roles
var roles []Role

// ReadRoles - gets the role model
func ReadRoles() []Role {
	roles = append(roles, Role{ID: 1, ImageURL: "https://google.com", Name: "John Doe", Ingredients: "Water, eggs, kefir"})
	roles = append(roles, Role{ID: 2, ImageURL: "https://google.com/name", Name: "Jon Blame", Ingredients: "Coffee, eggs, kefir"})
	roles = append(roles, Role{ID: 3, ImageURL: "https://google.com/3", Name: "Ana Blame", Ingredients: "Coffee, eggs, water"})
	roles = append(roles, Role{ID: 4, ImageURL: "https://google.com/4", Name: "Max Blame", Ingredients: "Coffee, eggs, kefir"})
	fmt.Printf("Hi there! %v \n", roles)
	return roles
}

// ! TODO - remove it

// SayModelRole - just for greeting and test
func SayModelRole() {
	fmt.Println("Hi there!")
}

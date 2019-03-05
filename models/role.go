package models

import (
	"fmt"
)

//Roles is popula bakery
type Role struct {
	ID          int    `json:"id"`
	ImageUrl    string `json:"url"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

//initiate roles
var roles []Role

func RoleModel() {
	roles = append(roles, Role{ID: 1, ImageUrl: "https://google.com", Name: "John Doe", Ingredients: "Water, eggs, kefir"})
	fmt.Printf("Hi there! % \n", roles)

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mux"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("couldnt connect to DB ALL user")
	}
	defer db.Close()

	var users []User
	db.Find(&users)

	db.AutoMigrate(&User{})
	json.NewEncoder(w).Encode(users)

	fmt.Fprintf(w, "All user Endpoint hit")
}

func NewUser(w http.ResponseWriter, r *http.Request) {

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("couldnt connect to DB ALL user")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	fmt.Fprintf(w, "New user endpoint hit and added succesfully: %s", vars)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("couldnt connect to DB ALL user")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Delete user Hit and User Deleted succesfully")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("couldnt connect to DB ALL user")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "update user hit")
}

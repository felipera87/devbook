package controllers

import (
	"net/http"
)

// CreateUser creates an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("creating user"))
}

// SearchUsers searches for all users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search all users"))
}

// SearchUser search for a single an user
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search user"))
}

// UpdateUser updates an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating user"))
}

// DeleteUser deletes an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting user"))
}

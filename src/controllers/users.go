package controllers

import "net/http"

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Teste get one user"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Teste get all users"))
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Teste save a user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Teste delete a user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Teste update a user"))
}

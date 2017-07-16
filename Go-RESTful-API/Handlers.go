package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)
/*
Concrete implementation to handle requet from the data sources
*/
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	var student Student

	result, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal([]byte(result), &student)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t := Repo_AddStudent(student)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func addClass(w http.ResponseWriter, r *http.Request) {
	var class Class

	result, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal([]byte(result), &class)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t := Repo_AddClass(class)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func getTotalScore(w http.ResponseWriter, r *http.Request) {
	var id string
	vars := mux.Vars(r)
	id = vars["id"]

	result := Repo_GetTotalScore(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func getTeacherFromScore(w http.ResponseWriter, r *http.Request) {
	result := Repo_GetTeacherFromScore()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
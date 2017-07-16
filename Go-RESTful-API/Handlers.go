package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	var student Student

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal([]byte(b), &student)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t := RepoCreateStudent(student)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func GetSumScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId string

	todoId = vars["id"]

	t := RepoGetSumScore(todoId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func GetTeacherFromScore(w http.ResponseWriter, r *http.Request) {
	t := RepoGetTopScoreTeacher()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func AddClass(w http.ResponseWriter, r *http.Request) {
	var todo class

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal([]byte(b), &todo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t := RepoCreateclass(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
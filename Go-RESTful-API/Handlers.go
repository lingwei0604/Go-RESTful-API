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
	var todo Student
	//
	//var todo map[string]interface{}
	//body,err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	panic(err)
	//}
	//if err := r.Body.Close(); err != nil {
	//	panic(err)
	//}
	var str = `{"Id":"26", "ClassNum":"01", "Score":"99"}`
	//消息体 分配到我们的 结构体上
	json.Unmarshal([]byte(str), &todo);
	//if err := json.Unmarshal([]byte(body), &todo); err != nil {
	//
	//	w.Header().Set("Content-Type", "multipart/form-data; charset=UTF-8")
	//	//w.WriteHeader(422) // unprocessable entity   multipart/form-data
	//	if err := json.NewEncoder(w).Encode(err); err != nil {
	//		panic(err)
	//		fmt.Print("this is")
	//	}
	//}


	//b, err := ioutil.ReadAll(r.Body)
	//defer r.Body.Close()
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//
	//err = json.Unmarshal([]byte(b), &todo)
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//

	//vars := mux.Vars(r)
	//todo.Id = vars["Id"]
	//todo.ClassNum = vars["ClassNum"]
	//todo.Score = vars["Score"]

	//返回创建实体的 json 表示
	t := RepoCreateStudent(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}


func GetSumScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId string
	//var err error
	//if todoId, err = strconv.Atoi(vars["id"]); err != nil {
	//	panic(err)
	//}

	todoId = vars["id"]
	//return t 为 sumScore
	t := RepoGetSumScore(todoId)

	//状态码不同
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}


func GetTeacherFromScore(w http.ResponseWriter, r *http.Request) {
	//var todo Student
	//body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	//if err != nil {
	//	panic(err)
	//}
	//if err := r.Body.Close(); err != nil {
	//	panic(err)
	//}
	//if err := json.Unmarshal(body, &todo); err != nil {
	//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//	w.WriteHeader(422) // unprocessable entity
	//	if err := json.NewEncoder(w).Encode(err); err != nil {
	//		panic(err)
	//	}
	//}

	t := RepoGetTopScoreTeacher()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func AddClass(w http.ResponseWriter, r *http.Request) {
	var todo class
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	panic(err)
	//}
	//if err := r.Body.Close(); err != nil {
	//	panic(err)
	//}
	//if err := json.Unmarshal([]byte(body), &todo); err != nil {
	//	w.Header().Set("Content-Type", "multipart/form-data; charset=UTF-8")
	//	w.WriteHeader(422) // unprocessable entity
	//	if err := json.NewEncoder(w).Encode(err); err != nil {
	//		panic(err)
	//	}
	//}

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
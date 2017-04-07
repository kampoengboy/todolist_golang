package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var errorTemplate = "<html><body><h1>Error rendering template %s</h1><p>%s</p></body></html>"

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w,
			fmt.Sprintf(errorTemplate, name, err),
			http.StatusInternalServerError,
		)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, "index/home", todos)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/", http.StatusFound)
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["todoId"])
	if err != nil {
		panic(err)
	}

	err2 := RepoDestroyTodo(todoId)
	if err2 != nil {
		panic(err)
	}
	redirect(w, r)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	var name = ""
	if r.Method == http.MethodPost {
		name = r.FormValue("name")
	}
	todo.Name = name
	RepoCreateTodo(todo)
	redirect(w, r)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(json.Unmarshal(body, &todo))
	// if err := r.Body.Close(); err != nil {
	// 	panic(err)
	// }
	// if err := json.Unmarshal(body, &todo); err != nil {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(422) // unprocessable entity
	// 	if err := json.NewEncoder(w).Encode(err); err != nil {
	// 		panic(err)
	// 	}
	// }

	// t := RepoCreateTodo(todo)

	// //redirect(w, r)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusCreated)
	// if err := json.NewEncoder(w).Encode(t); err != nil {
	// 	panic(err)
	// }
}

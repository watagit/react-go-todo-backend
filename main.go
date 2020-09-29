package main

import (
  "encoding/json"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

type Todo struct {
  ID int `json:"id"`
  Content string `json:"content"`
  Completed bool `json:"completed"`
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Add("Access-Control-Allow-Origin", "*")
  json.NewEncoder(w).Encode(todos)
}

func main() {
  r := mux.NewRouter()

  todos = append(todos, Todo{ID: 1, Content: "砂糖を買う", Completed: false})
  todos = append(todos, Todo{ID: 2, Content: "課題を終わらせる", Completed: false})
  todos = append(todos, Todo{ID: 3, Content: "歯医者に行く", Completed: false})

  r.HandleFunc("/api/v1/todos", getTodos).Methods("GET")

  log.Fatal(http.ListenAndServe(":8080", r))
}

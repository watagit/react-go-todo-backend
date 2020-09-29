package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strconv"

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

func addTodo(w http.ResponseWriter, r *http.Request) {
  var todo Todo

  json.NewDecoder(r.Body).Decode(&todo)
  fmt.Println("todo: ", todo)

  todos = append(todos, todo)

  json.NewEncoder(w).Encode(todos)
}

func removeTodo(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  fmt.Println("params: ", params)

  id, _ := strconv.Atoi(params["id"])
  fmt.Println("id: ", id)

  fmt.Println("todos: ", todos)

  for i, item := range todos {
    if item.ID == id {
      todos = append(todos[:i], todos[i+1:]...)
    }
  }
  json.NewEncoder(w).Encode(todos)
}

func main() {
  r := mux.NewRouter()

  todos = append(todos,
    Todo{ID: 1, Content: "砂糖を買う", Completed: false},
    Todo{ID: 2, Content: "課題を終わらせる", Completed: false},
    Todo{ID: 3, Content: "歯医者に行く", Completed: false},
  )

  r.HandleFunc("/api/v1/todos", getTodos).Methods("GET")
  r.HandleFunc("/api/v1/todos", addTodo).Methods("POST")
  r.HandleFunc("/api/v1/todos/{id}", removeTodo).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8080", r))
}

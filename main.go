package main

import (
  "database/sql"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strconv"

  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
)

type Todo struct {
  ID        int     `json:"id"`
  Content   string  `json:"content"`
  Completed bool    `json:"completed"`
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Add("Access-Control-Allow-Origin", "*")
  _ = json.NewEncoder(w).Encode(todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
  var todo Todo

  _ = json.NewDecoder(r.Body).Decode(&todo)
  fmt.Println("todo: ", todo)

  todos = append(todos, todo)

  _ = json.NewEncoder(w).Encode(todos)
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
  _ = json.NewEncoder(w).Encode(todos)
}

func main() {
  db, err := sql.Open("mysql", "root:includemath.h@tcp(127.0.0.1:3306)/react_go_todo")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  rows, err := db.Query("select * from todo")
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  for rows.Next() {
    var todo Todo
    err := rows.Scan(&todo.ID, &todo.Content, &todo.Completed)
    if err != nil {
      panic(err.Error())
    }
    todos = append(todos, todo)
  }

  err = rows.Err()
  if err != nil {
    panic(err.Error())
  }

  r := mux.NewRouter()

  r.HandleFunc("/api/v1/todos", getTodos).Methods("GET")
  r.HandleFunc("/api/v1/todos", addTodo).Methods("POST")
  r.HandleFunc("/api/v1/todos/{id}", removeTodo).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8080", r))
}

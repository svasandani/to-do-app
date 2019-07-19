package main

import (
  "net/http"
  "github.com/svasandani/to-do-app/handlers"
  "github.com/codegangsta/negroni"
)

func todoHandlers(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case http.MethodGet: handlers.GetTodoListHandler(w, r)
    case http.MethodPost: handlers.AddTodoHandler(w, r)
    case http.MethodPut: handlers.CompleteTodoHandler(w, r)
    case http.MethodDelete: handlers.DeleteTodoHandler(w, r)
    case http.MethodOptions: handlers.OptionsHandler(w, r)
  }
}

func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("/init/", handlers.IndexHandler)
  mux.HandleFunc("/todo", handlers.DeleteTableHandler)
  mux.HandleFunc("/todo/", todoHandlers)

  n := negroni.Classic()
  // n.Use(negroni.HandlerFunc(setupResponse))
  n.UseHandler(mux)
  n.Run(":2000")
}

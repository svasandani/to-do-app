package handlers

import (
  "fmt"
  "encoding/json"
  "io"
  "io/ioutil"
  "net/http"
  // "html/template"

  "github.com/svasandani/to-do-app/todo"
)

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	  (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
  encoder := json.NewEncoder(w)

  if err := encoder.Encode(todo.Get()); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
  defer r.Body.Close()

  todoItem, err := convertHTTPToTodo(r.Body)

  if err !=  nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  encoder := json.NewEncoder(w)

  if err = (encoder.Encode(todo.Add(todoItem.Contents))); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func CompleteTodoHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
  defer r.Body.Close()

  todoItem, err := convertHTTPToTodo(r.Body)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  if err = todo.Complete(todoItem.ID); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)

  if r.Method != http.MethodDelete {
    return
  }

  id := r.FormValue("id")
  fmt.Println("\n\n",id,"\n\n")

  if err := todo.Delete(id); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func convertHTTPToTodo(httpBody io.ReadCloser) (todo.Todo, error) {
  body, err := ioutil.ReadAll(httpBody)

  if err != nil {
    return todo.Todo{}, err
  }

  defer httpBody.Close()
  return convertJSONToTodo(body)
}

func convertJSONToTodo(jsonBody []byte) (todo.Todo, error) {
  var todoItem todo.Todo
  err := json.Unmarshal(jsonBody, &todoItem)

  if err != nil {
    return todo.Todo{}, err
  }

  return todoItem, nil
}

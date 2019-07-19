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
  fmt.Println("hello")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodOptions {
    OptionsHandler(w, r)
    return
  }
  setupResponse(&w, r)
  todo.InitializeList(r.FormValue("key"))
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
  encoder := json.NewEncoder(w)
  key := r.FormValue("key")

  if err := encoder.Encode(todo.Get(key)); err != nil {
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
  key := r.FormValue("key")

  if err = (encoder.Encode(todo.Add(key, todoItem.Contents))); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func CompleteTodoHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
  defer r.Body.Close()

  todoItem, _ := convertHTTPToTodo(r.Body)
  // id := r.FormValue("id")
  key := r.FormValue("key")

  if truth, _ := todo.IsIncomplete(key, todoItem.ID); (truth) {
    if err := todo.Complete(key, todoItem.ID); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  } else {
   if err := todo.Uncomplete(key, todoItem.ID); err != nil {
     http.Error(w, err.Error(), http.StatusInternalServerError)
   }
  }
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)

  if r.Method != http.MethodDelete {
    return
  }

  id := r.FormValue("id")
  key := r.FormValue("key")
  fmt.Println("\n\n",id,"\n\n")

  if err := todo.Delete(key, id); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func DeleteTableHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)

  if r.Method != http.MethodDelete {
    return
  }

  key := r.FormValue("key")
  fmt.Println("\n\n",key,"\n\n")

  if err := todo.DeleteTable(key); err != nil {
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

package todo

import(
  "errors"
  "sync"

  "github.com/rs/xid"
)

var (
  list []Todo
  mtx sync.RWMutex
  once sync.Once
)

// initialize todo list only once so it doesn't get erased
func init() {
  once.Do(initializeList)
}

func initializeList() {
  list = []Todo{}
}

// Todo data structure
type Todo struct {
  ID string `json:"id"`
  Contents string `json:"contents"`
  Completed bool `json:"completed"`
}

// return all todo items
func Get() []Todo {
  return list
}

//// public functions

// add new todo
func Add(contents string) string {
  t := newTodo(contents)
  mtx.Lock()
  list = append(list, t)
  mtx.Unlock()
  return t.ID
}

// delete todo based on its id
func Delete(id string) error {
  index, err := findTodoByID(id)
  if err != nil {
    return err
  }
  removeTodoByIndex(index)
  return nil
}

// mark todo as completed
func Complete(id string) error {
  index, err := findTodoByID(id)
  if err != nil {
    return err
  }
  markCompletedByIndex(index)
  return nil
}

// mark todo as not completed
func Uncomplete(id string) error {
  index, err := findTodoByID(id)
  if err != nil {
    return err
  }
  markUncompletedByIndex(index)
  return nil
}

// check if already completed
func IsIncomplete(id string) (bool, error) {
  index, err := findTodoByID(id)
  if err != nil {
    return false, err
  }
  if list[index].Completed == true {
    return false, nil
  } else {
    return true, nil
  }
}

//// private functions

// create new todo item
func newTodo(contents string) Todo {
  return Todo {ID: xid.New().String(), Contents: contents, Completed: false,}
}

// find todo based on its unique ID
func findTodoByID(id string) (int, error) {
  for i, t := range list {
    if isMatching(t.ID, id) {
      return i, nil
    }
  }

  return -1, errors.New("Couldn't find given Todo in the list.")
}

// remove todo based on its index
func removeTodoByIndex(index int) {
  mtx.Lock()
  list = append(list[:index], list[index+1:]...)
  mtx.Unlock()
}

// mark a todo completed based on its index
func markCompletedByIndex(index int) {
  mtx.Lock()
  list[index].Completed = true
  mtx.Unlock()
}

// mark a todo not completed based on its index
func markUncompletedByIndex(index int) {
  mtx.Lock()
  list[index].Completed = false
  mtx.Unlock()
}

// check if matching
func isMatching(id1 string, id2 string) bool {
  return id1 == id2
}

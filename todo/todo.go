package todo

import(
  "errors"
  "sync"
  "log"
  "fmt"

  "github.com/rs/xid"

  "database/sql"
  _ "github.com/lib/pq"
)

var (
  list []Todo
  mtx sync.RWMutex
  once sync.Once
  db *sql.DB
  savedKey string
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "12301101"
  dbname   = "todos"
)

// initialize todo list only once so it doesn't get erased
func init() {
  // once.Do(initializeList)
}

func InitializeList(key string) {
  list = []Todo{}
  savedKey = key
  var err error

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
  db, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    log.Fatalf(err.Error())
  }

  exString := fmt.Sprintf("SELECT 1 FROM %v LIMIT 1", "T" + key)
  if _, err := db.Query(exString); err == nil {
    exString2 := fmt.Sprintf("SELECT ID, CONTENTS, COMPLETED FROM %v", "T" + key)
    row, _ := db.Query(exString2)

    for row.Next() {

      var t Todo
      var complete int

      row.Scan(&t.ID, &t.Contents, &complete)

      if complete == 1 {
        t.Completed = true
      } else {
        t.Completed = false
      }

      if index, _ := findTodoByID(t.ID); index == -1 {
        list = append(list, t)
      }
    }
  } else {
    execString := fmt.Sprintf("create table %v(id text, contents text, completed integer);", "T" + key)
    db.Exec(execString)
    list = []Todo{}
  }
}

// Todo data structure
type Todo struct {
  ID string `json:"id"`
  Contents string `json:"contents"`
  Completed bool `json:"completed"`
}

// return all todo items
func Get() []Todo {
  fmt.Println(list)
  return list
}

//// public functions

// delete table
func DeleteTable(key string) error {
  execString := fmt.Sprintf("drop table %v;", "T" + key)
  _, err := db.Exec(execString)
  return err
}

// add new todo
func Add(contents string) string {
  t := newTodo(contents)
  mtx.Lock()
  list = append(list, t)
  var bit int
  if t.Completed {
    bit = 1
  } else {
    bit = 0
  }
  execString := fmt.Sprintf("INSERT INTO %v (ID, CONTENTS, COMPLETED) VALUES ('%v', '%v', %v)", "T" + savedKey, t.ID, t.Contents, bit)
  fmt.Println("\n\n\n", execString, "\n\n\n")
  result, err := db.Exec(execString)
  if err != nil {
    panic(err)
  }
  mtx.Unlock()
  id, _ := result.LastInsertId()
  return string(id)
}

// delete todo based on its id
func Delete(id string) error {
  index, err := findTodoByID(id)
  if err != nil {
    return err
  }
  removeTodoByIndex(index)
  execString := fmt.Sprintf("DELETE FROM %v WHERE ID = '%v'", "T" + savedKey, id)
  fmt.Println("\n\n\n", execString, "\n\n\n")
  if _, err := db.Exec(execString); err != nil {
    return err
  }
  return nil
}

// mark todo as completed
func Complete(id string) error {
  index, err := findTodoByID(id)
  if err != nil {
    return err
  }
  markCompletedByIndex(index)
  execString := fmt.Sprintf("UPDATE %v SET COMPLETED = 1 WHERE ID = '%v'", "T" + savedKey, id)
  fmt.Println("\n\n\n", execString, "\n\n\n")
  _, err = db.Exec(execString)
  if err != nil {
    return err
  }
  return nil
}

// mark todo as not completed
func Uncomplete(id string) error {
  index, err := findTodoByID(id)
  if err != nil {
    return err
  }
  markUncompletedByIndex(index)
  execString := fmt.Sprintf("UPDATE %v SET COMPLETED = 0 WHERE ID = '%v'", "T" + savedKey, id)
  fmt.Println("\n\n\n", execString, "\n\n\n")
  _, err = db.Exec(execString)
  if err != nil {
    return err
  }
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

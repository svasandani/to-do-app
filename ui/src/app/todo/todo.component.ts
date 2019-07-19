import { Component, OnInit, Output, EventEmitter, Input } from '@angular/core';
import { TodoService, Todo } from '../todo.service';

@Component({
  selector: 'app-todo',
  templateUrl: './todo.component.html',
  styleUrls: ['./todo.component.css'],
  providers: [TodoService]
})
export class TodoComponent implements OnInit {
  @Output() hideTodoEvent = new EventEmitter<boolean>();
  @Input('dbKey') key: string;

  displayedColumns = ['id', 'contents', 'completed']

  currentTodo: Todo;
  activeTodos: Todo[];
  completedTodos: Todo[];
  todoContents: string;

  hideTodos() {
    if (this.completedTodos.length == 0 && this.activeTodos.length == 0) {
      this.todoService.deleteTable(this.key).subscribe(() => {
        console.log(this.key);
      });;
    }
    this.hideTodoEvent.emit(false);
  }

  constructor(private todoService: TodoService) { }

  ngOnInit() {
    this.activeTodos = [];
    this.completedTodos = [];
    this.getAll();
  }

  getAll() {
    this.todoService.getTodoList(this.key).subscribe((data: Todo[]) => {
      if (data != null) {
        this.activeTodos = data.filter((a) => !a.completed);
        this.completedTodos = data.filter((a) => a.completed);
      } else {
        this.activeTodos = [];
        this.completedTodos = [];
      }
    });
  }

  addTodo() {
    console.log(this.todoContents.replace("'", "''"));
    var newTodo: Todo = {
      contents: this.todoContents.replace("'", "''"),
      id: '',
      completed: false
    };

    console.log(newTodo)

    this.todoService.addTodo(this.key, newTodo).subscribe(() => {
      this.getAll();
      this.todoContents = '';
    });
  }

  completeTodo(todo: Todo) {
    this.todoService.completeTodo(this.key, todo).subscribe(() => {
      this.getAll();
    });
  }

  deleteTodo(todo: Todo) {
    console.log(todo)

    this.todoService.deleteTodo(this.key, todo).subscribe(() => {
      this.getAll();
    })
  }

  doControlEvent($event) {
    console.log($event);
    if ($event == "") {
      this.currentTodo = undefined;
    } else {
      var newTodo: Todo = { id: $event, contents: "", completed: false }
      this.completeTodo(newTodo);
    }
  }

  doDeleteEvent($event) {
    this.currentTodo = undefined;
    var newTodo: Todo = { id: $event, contents: "", completed: false }
    this.deleteTodo(newTodo);
  }

}

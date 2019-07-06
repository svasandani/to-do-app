import { Component, OnInit } from '@angular/core';
import { TodoService, Todo } from '../todo.service';

@Component({
  selector: 'app-todo',
  templateUrl: './todo.component.html',
  styleUrls: ['./todo.component.css'],
  providers: [TodoService]
})
export class TodoComponent implements OnInit {

  displayedColumns = ['id', 'contents', 'completed']

  currentTodo: Todo;
  activeTodos: Todo[];
  completedTodos: Todo[];
  todoContents: string;

  constructor(private todoService: TodoService) { }

  ngOnInit() {
    this.getAll();
  }

  getAll() {
    this.todoService.getTodoList().subscribe((data: Todo[]) => {
      this.activeTodos = data.filter((a) => !a.completed);
      this.completedTodos = data.filter((a) => a.completed);
    });
  }

  addTodo() {
    var newTodo: Todo = {
      contents: this.todoContents,
      id: '',
      completed: false
    };

    this.todoService.addTodo(newTodo).subscribe(() => {
      this.getAll();
      this.todoContents = '';
    });
  }

  completeTodo(todo: Todo) {
    this.todoService.completeTodo(todo).subscribe(() => {
      this.getAll();
    });
  }

  deleteTodo(todo: Todo) {
    this.todoService.deleteTodo(todo).subscribe(() => {
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

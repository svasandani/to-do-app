import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class TodoService {
  constructor(private httpClient: HttpClient) { }

  getTodoList(key) {
    return this.httpClient.get(environment.gateway + '/todo/?key=' + key);
  }

  addTodo(key, todo: Todo) {
    return this.httpClient.post(environment.gateway + '/todo/?key=' + key, todo);
  }

  completeTodo(key, todo: Todo) {
    return this.httpClient.put(environment.gateway + '/todo/?key=' + key + '&id=', todo);
  }

  deleteTodo(key, todo: Todo) {
    return this.httpClient.delete(environment.gateway + '/todo/?key=' + key + '&id=' + todo.id);
  }

  deleteTable(key) {
    return this.httpClient.delete(environment.gateway + '/todo?key=' + key);
  }

  setKey(key) {
    return this.httpClient.get(environment.gateway + '/init/?key=' + key);
  }
}

export class Todo {
  id: string;
  contents: string;
  completed: boolean;
}

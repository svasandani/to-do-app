import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'ui';

  showTodo: boolean = false;
  databaseKey: string;

  receiveTodoKey($event) {
    console.log("hi");
    this.databaseKey = $event;
    this.showTodo = true;
  }

  hideTodos($event) {
    this.showTodo = $event;
  }
}

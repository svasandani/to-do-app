import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'ui';

  showTodo: boolean = false;

  receiveTodoKey($event) {
    console.log("hi")
    this.showTodo = $event;
  }
}

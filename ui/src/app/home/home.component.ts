import { Component, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  @Output() showTodoEvent = new EventEmitter<boolean>();

  hashString(str) {
    let hash = 0;
    for (let i = 0; i < str.length; i++) {
      hash += Math.pow(str.charCodeAt(i) * 31, str.length - i);
      hash = hash & hash; // Convert to 32bit integer
    }
    return hash;
  }

  showTodos(input, placeholder) {
    if (this.hashString(input.value) == 910670847) {
      this.showTodoEvent.emit(true);
    } else {
      placeholder.style.color = "red";
      placeholder.innerHTML = "Your key is incorrect."
      return;
    }
  }
}

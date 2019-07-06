import { Component, OnInit, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-todo-details',
  templateUrl: './todo-details.component.html',
  styleUrls: ['./todo-details.component.css'],
  inputs: ['todo']
})
export class TodoDetailsComponent implements OnInit {
  switch: boolean = false;

  @Output() doEvent = new EventEmitter<string>();
  @Output() deleteEvent = new EventEmitter<string>();

  doControlEvent(control) {
    if (control == "") {
      this.doEvent.emit("");
    } else {
      this.switch = !this.switch;
      this.doEvent.emit(control);
    }
  }

  deleteControlEvent(id) {
    this.deleteEvent.emit(id);
  }

  constructor() { }

  ngOnInit() {
  }

}

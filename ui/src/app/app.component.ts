import { Component, OnInit } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [CookieService]
})
export class AppComponent {
  title = 'ui';

  showTodo: boolean = false;
  databaseKey: string;
  cookieKey: string;

  constructor(private cookieService: CookieService) { }

  ngOnInit() {
    if (this.cookieService.check('key')) {
      this.databaseKey = this.cookieService.get('key');
      this.showTodo = true;
    }
  }

  receiveTodoKey($event) {
    console.log("hi");
    this.databaseKey = $event;
    this.cookieService.set('key', $event);
    this.showTodo = true;
  }

  hideTodos($event) {
    this.showTodo = $event;
    this.databaseKey = "";
    this.cookieService.delete('key');
  }
}

import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MaterialModule } from './material.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FlexLayoutModule } from '@angular/flex-layout';

import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { TodoComponent } from './todo/todo.component';
import { TodoDetailsComponent } from './todo-details/todo-details.component';

@NgModule({
  declarations: [
    AppComponent, HomeComponent, TodoComponent, TodoDetailsComponent
  ],
  imports: [
    BrowserModule, FormsModule, MaterialModule, BrowserAnimationsModule, FlexLayoutModule, HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

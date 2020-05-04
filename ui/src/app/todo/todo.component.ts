import { Component, OnInit } from '@angular/core';
import { TodoService, Todo } from '../todo.service';

@Component({
  selector: 'app-todo',
  templateUrl: './todo.component.html',
  styleUrls: ['./todo.component.css']
})
export class TodoComponent implements OnInit {

  activeTodos: Todo[];
  completedTodos: Todo[];
  todoMessage: string;

  constructor(private todoService: TodoService) { }

  ngOnInit() {
    this.getAll();
  }

  // getAll todo data call todoservice
  getAll() {
    this.todoService.getTodoList().subscribe((data: Todo[]) => {
      this.activeTodos = data.filter((a) => !a.complete);
      this.completedTodos = data.filter((a) => a.complete);
    });
  }

  // addTodo data call todoservice
  addTodo() {
    var newTodo : Todo = {
      message: this.todoMessage,
      id: '',
      complete: false
    };

    this.todoService.addTodo(newTodo).subscribe(() => {
      this.getAll();
      this.todoMessage = '';
    });
  }

  // completeTodo data call todoservice
  completeTodo(todo: Todo) {
    todo.complete = true
    this.todoService.completeTodo(todo).subscribe(() => {
      this.getAll();
    });
  }

  // deleteTodo data call todoservice
  deleteTodo(todo: Todo) {
    this.todoService.deleteTodo(todo).subscribe(() => {
      this.getAll();
    })
  }
}

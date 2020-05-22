import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";

@Injectable()
export class TodoService {
  // depedency injection for httpClient
  constructor(private httpClient: HttpClient) {}

  // call get todo api url `/api/todo
  getTodoList() {
    return this.httpClient.get("/api/todo");
  }

  // call create todo api url `/api/todo
  addTodo(todo: Todo) {
    return this.httpClient.post("/api/todo", todo);
  }

  // call put/update todo api url `/api/todo/id
  completeTodo(todo: Todo) {
    return this.httpClient.put("/api/todo/" + todo.id, todo);
  }

  // call delete/remove todo api url `/api/todo/id
  deleteTodo(todo: Todo) {
    return this.httpClient.delete("/api/todo/" + todo.id);
  }
}

// todo class
export class Todo {
  id: string;
  message: string;
  complete: boolean;
}

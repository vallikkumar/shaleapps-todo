import { async, ComponentFixture, TestBed, inject } from '@angular/core/testing';

import { TodoComponent } from './todo.component';
import { TodoService } from '../todo.service';
import { FormsModule } from '@angular/forms';
import { HttpClient, HttpHandler } from '@angular/common/http';

describe('TodoComponent', () => {
  let component: TodoComponent;
  let fixture: ComponentFixture<TodoComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TodoComponent ],
      imports: [ FormsModule ],
      providers: [ TodoService, HttpClient, HttpHandler ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TodoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create...', inject([TodoService], (service: TodoService) => {
    expect(service).toBeTruthy();
  }));

});

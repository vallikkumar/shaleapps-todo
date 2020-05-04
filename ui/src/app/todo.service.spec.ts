import { TestBed, inject } from '@angular/core/testing';

import { TodoService } from './todo.service';
import { HttpClient, HttpHandler } from '@angular/common/http';

describe('TodoService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [TodoService, HttpClient, HttpHandler ],
    });
  });

  it('should create...', inject([HttpClient], (httpClient: HttpClient) => {
    expect(httpClient).toBeTruthy();
  }));
});

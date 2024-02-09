import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { Task } from '../Task';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json'
  })
}

@Injectable({
  providedIn: 'root'
})
export class TaskService {
  private apiUrl = "http://localhost:8080/api/tasks/";
  private ngUnsubscribe = new Subject<void>();

  constructor(private http: HttpClient) { }

  getTasks(): Observable<Task[]> {
    return this.http.get<Task[]>(this.apiUrl).pipe(
      takeUntil(this.ngUnsubscribe)
    );
  }

  deleteTask(task: Task): Observable<Task> {
    const url = `${this.apiUrl}${task.task_id}`;
    return this.http.delete<Task>(url).pipe(
      takeUntil(this.ngUnsubscribe)
    );
  }

  updateTaskStatus(task: Task): Observable<Task> {
    const url = `${this.apiUrl}${task.task_id}`;
    return this.http.put<Task>(url, task, httpOptions).pipe(
      takeUntil(this.ngUnsubscribe)
    );
  }

  addTask(task: Task): Observable<Task> {
    return this.http.post<Task>(this.apiUrl, task, httpOptions).pipe(
      takeUntil(this.ngUnsubscribe)
    );
  }

  ngOnDestroy(): void {
    this.ngUnsubscribe.next();
    this.ngUnsubscribe.complete();
  }
}

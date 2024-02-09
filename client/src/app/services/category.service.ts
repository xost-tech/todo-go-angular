import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { Category } from '../Category';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {
  private apiUrl = "http://localhost:8080/api/categories/";
  private ngUnsubscribe = new Subject<void>();

  constructor(private http: HttpClient) { }

  getCategories(): Observable<Category[]> {
    return this.http.get<Category[]>(this.apiUrl).pipe(
      takeUntil(this.ngUnsubscribe)
    );
  }

  ngOnDestroy(): void {
    this.ngUnsubscribe.next();
    this.ngUnsubscribe.complete();
  }
}

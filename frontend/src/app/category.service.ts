/*
import { Injectable } from '@angular/core';

import { Observable, of } from 'rxjs';

import { Category } from './category';
import { CATEGORIES } from './mock-categories';
import { MessageService } from './message.service';

@Injectable({
  providedIn: 'root',
})
export class CategoryService {

  constructor(private messageService: MessageService) { }

  getCategories(): Observable<Category[]> {
    const categories = of(CATEGORIES);
    this.messageService.add('CategoryService: fetched categories');
    return categories;
  }

  getCategory(id: number): Observable<Category> {
    // For now, assume that a category with the specified `id` always exists.
    // Error handling will be added in the next step of the tutorial.
    const category = CATEGORIES.find(h => h.id === id)!;
    this.messageService.add(`CategoryService: fetched category id=${id}`);
    return of(category);
  }
}
*/

import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { Category } from './category';
import { MessageService } from './message.service';


@Injectable({ providedIn: 'root' })
export class CategoryService {

  private categoriesUrl = 'api/categories';  // URL to web api

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient,
    private messageService: MessageService) { }

  /** GET categories from the server */
  getCategories(): Observable<Category[]> {
    return this.http.get<Category[]>(this.categoriesUrl)
      .pipe(
        tap(_ => this.log('fetched categories')),
        catchError(this.handleError<Category[]>('getCategories', []))
      );
  }

  /** GET category by id. Return `undefined` when id not found */
  getCategoryNo404<Data>(id: number): Observable<Category> {
    const url = `${this.categoriesUrl}/?id=${id}`;
    return this.http.get<Category[]>(url)
      .pipe(
        map(categories => categories[0]), // returns a {0|1} element array
        tap(h => {
          const outcome = h ? `fetched` : `did not find`;
          this.log(`${outcome} category id=${id}`);
        }),
        catchError(this.handleError<Category>(`getCategory id=${id}`))
      );
  }

  /** GET category by id. Will 404 if id not found */
  getCategory(id: number): Observable<Category> {
    const url = `${this.categoriesUrl}/${id}`;
    return this.http.get<Category>(url).pipe(
      tap(_ => this.log(`fetched category id=${id}`)),
      catchError(this.handleError<Category>(`getCategory id=${id}`))
    );
  }

  /* GET categories whose name contains search term */
  searchCategories(term: string): Observable<Category[]> {
    if (!term.trim()) {
      // if not search term, return empty category array.
      return of([]);
    }
    return this.http.get<Category[]>(`${this.categoriesUrl}/?name=${term}`).pipe(
      tap(x => x.length ?
         this.log(`found categories matching "${term}"`) :
         this.log(`no categories matching "${term}"`)),
      catchError(this.handleError<Category[]>('searchCategories', []))
    );
  }

  //////// Save methods //////////

  /** POST: add a new category to the server */
  addCategory(category: Category): Observable<Category> {
    return this.http.post<Category>(this.categoriesUrl, category, this.httpOptions).pipe(
      tap((newCategory: Category) => this.log(`added category w/ id=${newCategory.id}`)),
      catchError(this.handleError<Category>('addCategory'))
    );
  }

  /** DELETE: delete the category from the server */
  deleteCategory(id: number): Observable<Category> {
    const url = `${this.categoriesUrl}/${id}`;

    return this.http.delete<Category>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted category id=${id}`)),
      catchError(this.handleError<Category>('deleteCategory'))
    );
  }

  /** PUT: update the category on the server */
  updateCategory(category: Category): Observable<any> {
    return this.http.put(this.categoriesUrl, category, this.httpOptions).pipe(
      tap(_ => this.log(`updated category id=${category.id}`)),
      catchError(this.handleError<any>('updateCategory'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  /** Log a CategoryService message with the MessageService */
  private log(message: string) {
    this.messageService.add(`CategoryService: ${message}`);
  }
}
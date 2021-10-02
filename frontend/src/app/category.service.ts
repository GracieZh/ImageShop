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
    // For now, assume that a hero with the specified `id` always exists.
    // Error handling will be added in the next step of the tutorial.
    const category = CATEGORIES.find(h => h.id === id)!;
    this.messageService.add(`CategoryService: fetched category id=${id}`);
    return of(category);
  }
}

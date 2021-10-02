import { Injectable } from '@angular/core';
import { InMemoryDbService } from 'angular-in-memory-web-api';
import { Category } from './category';

@Injectable({
  providedIn: 'root',
})
export class InMemoryDataService implements InMemoryDbService {
  createDb() {
    const categories = [
      { id: 11, name: 'Dr Nice' },
      { id: 12, name: 'Narco' },
      { id: 13, name: 'Bombasto' },
      { id: 14, name: 'Celeritas' },
      { id: 15, name: 'Magneta' },
      { id: 16, name: 'RubberMan' },
      { id: 17, name: 'Dynama' },
      { id: 18, name: 'Dr IQ' },
      { id: 19, name: 'Magma' },
      { id: 20, name: 'Tornado' }
    ];
    return {categories};
  }

  // Overrides the genId method to ensure that a category always has an id.
  // If the categories array is empty,
  // the method below returns the initial number (11).
  // if the categories array is not empty, the method below returns the highest
  // category id + 1.
  genId(categories: Category[]): number {
    return categories.length > 0 ? Math.max(...categories.map(category => category.id)) + 1 : 11;
  }
}
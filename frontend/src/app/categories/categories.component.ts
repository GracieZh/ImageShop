import { Component, OnInit } from '@angular/core';
import { Category } from '../category';
//import { CATEGORIES } from '../mock-categories';
import { CategoryService } from '../category.service';
import { MessageService } from '../message.service';

@Component({
  selector: 'app-categories',
  templateUrl: './categories.component.html',
  styleUrls: ['./categories.component.scss']
})
export class CategoriesComponent implements OnInit {
  //categories = CATEGORIES;
  selectedCategory?: Category;
  categories: Category[] = [];

  constructor(private categoryService: CategoryService, private messageService: MessageService) { }

  ngOnInit() {
    this.getCategories();
  }

  onSelect(category: Category): void {
    this.selectedCategory = category;
    this.messageService.add(`CategoriesComponent: Selected hero id=${category.id}`);
  }

  getCategories(): void {
    this.categoryService.getCategories()
        .subscribe(categories => this.categories = categories);
  }
}

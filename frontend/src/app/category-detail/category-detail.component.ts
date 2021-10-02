import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';

import { Category } from '../category';
import { CategoryService } from '../category.service';

@Component({
  selector: 'app-category-detail',
  templateUrl: './category-detail.component.html',
  styleUrls: [ './category-detail.component.scss' ]
})
export class CategoryDetailComponent implements OnInit {
  category: Category | undefined;

  constructor(
    private route: ActivatedRoute,
    private categoryService: CategoryService,
    private location: Location
  ) {}

  ngOnInit(): void {
    this.getCategory();
  }

  getCategory(): void {
    const id = parseInt(this.route.snapshot.paramMap.get('id')!, 10);
    this.categoryService.getCategory(id)
      .subscribe(category => this.category = category);
  }

  goBack(): void {
    this.location.back();
  }

  save(): void {
    if (this.category) {
      this.categoryService.updateCategory(this.category)
        .subscribe(() => this.goBack());
    }
  }
}
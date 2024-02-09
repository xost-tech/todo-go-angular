import { Component, Output, EventEmitter, OnInit, OnDestroy } from '@angular/core';
import { UiService } from '../../services/ui.service';
import { CategoryService } from '../../services/category.service';
import { Subscription } from 'rxjs';
import { Task } from '../../Task';
import { Category } from '../../Category';

@Component({
  selector: 'app-task-add',
  templateUrl: './task-add.component.html',
  styleUrl: './task-add.component.css'
})
export class TaskAddComponent implements OnInit, OnDestroy {
  @Output() onAddTask: EventEmitter<Task> = new EventEmitter();
  task_name!: string;
  task_description?: string;
  category_id?: number;
  due_date!: string;
  showAddTask: boolean = false;
  subscription!: Subscription;
  categories: Category[] = []

  constructor(private uiService: UiService, private categoryService: CategoryService) {
    this.subscription = this.uiService.onToggle().subscribe((value) => (this.showAddTask = value));
  }

  ngOnInit(): void {
    this.categoryService.getCategories().subscribe((categories) => {
      this.categories = categories;
      if (this.categories.length > 0) {
        this.category_id = this.categories[0].category_id;
      }
    });
  }

  ngOnDestroy(): void {
    this.categoryService.ngOnDestroy();
  }
  
  onSubmit() {
    if(!this.task_name) {
      alert('Please enter task name.');
      return;
    }

    const newTask = {
      user_id: 2,
      category_id: Number(this.category_id),
      task_name: this.task_name,
      task_description: this.task_description,
      task_status: 0,
      due_date: this.due_date
    }

    console.log('Request Payload:', newTask);

    this.onAddTask.emit(newTask);

    this.task_name = '';
    this.task_description = '';
    if (this.categories.length > 0) {
      this.category_id = this.categories[0].category_id;
    }
    this.due_date = '';
  }
}

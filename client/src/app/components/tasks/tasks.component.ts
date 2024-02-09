import { Component, OnDestroy, OnInit } from '@angular/core';
import { TaskService } from '../../services/task.service';
import { Task } from '../../Task';

@Component({
  selector: 'app-tasks',
  templateUrl: './tasks.component.html',
  styleUrl: './tasks.component.css'
})
export class TasksComponent implements OnInit, OnDestroy {
  tasks: Task[] = []

  constructor(private taskService: TaskService) {}

  ngOnInit(): void {
    this.taskService.getTasks().subscribe((tasks) => this.tasks = tasks);
  }

  ngOnDestroy(): void {
    this.taskService.ngOnDestroy();
  }

  deleteTask(task: Task) {
    this.taskService.deleteTask(task).subscribe(() => this.tasks = this.tasks.filter((t) => t.task_id !== task.task_id ));
  }

  updateTaskStatus(task: Task) {
    task.task_status = (task.task_status + 1) % 3;
    this.taskService.updateTaskStatus(task).subscribe();
  }

  addTask(task: Task) {
    this.taskService.addTask(task).subscribe((task) => this.tasks.push(task));
  }
}

export interface Task {
    task_id?: number;
    user_id: number;
    category_id: number;
    task_name: string;
    task_description?: string;
    task_status: number;
    due_date: string;
    created_timestamp?: string;
    modified_timestamp?: string;
  }
  
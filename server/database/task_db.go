package database

import (
	"todo-go-angular/server/models"
)

func CreateTask(task *models.Task) (int, error) {
	insertQuery := "INSERT INTO task (user_id, category_id, task_name, task_description, task_status, due_date) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := DB.Exec(insertQuery, task.UserID, task.CategoryID, task.TaskName, task.TaskDescription, models.Pending, task.DueDate)
	if err != nil {
		return 0, err
	}

	taskId, _ := result.LastInsertId()
	task.TaskID = int(taskId)

	return task.TaskID, nil
}

func GetTasks() ([]models.Task, error) {
	selectQuery := "SELECT * FROM task"
	rows, err := DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.TaskID,
			&task.UserID,
			&task.CategoryID,
			&task.TaskName,
			&task.TaskDescription,
			&task.TaskStatus,
			&task.DueDate,
			&task.CreatedTimestamp,
			&task.ModifiedTimestamp,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTasksByUserId(userId int) ([]models.Task, error) {
	selectQuery := "SELECT * FROM task WHERE user_id = ?"
	rows, err := DB.Query(selectQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.TaskID, &task.UserID, &task.CategoryID, &task.TaskName, &task.TaskDescription, &task.TaskStatus, &task.DueDate, &task.CreatedTimestamp, &task.ModifiedTimestamp)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTaskById(taskId int) (*models.Task, error) {
	selectQuery := "SELECT * FROM task WHERE task_id = ?"
	row := DB.QueryRow(selectQuery, taskId)

	var task models.Task
	err := row.Scan(&task.TaskID, &task.UserID, &task.CategoryID, &task.TaskName, &task.TaskDescription, &task.TaskStatus, &task.DueDate, &task.CreatedTimestamp, &task.ModifiedTimestamp)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func UpdateTask(taskId int, updatedTask *models.Task) error {
	updateQuery := "UPDATE task SET category_id = ?, task_name = ?, task_description = ?, task_status = ?, due_date = ? WHERE task_id = ?"
	_, err := DB.Exec(updateQuery, updatedTask.CategoryID, updatedTask.TaskName, updatedTask.TaskDescription, updatedTask.TaskStatus, updatedTask.DueDate, taskId)
	return err
}

func DeleteTask(taskId int) error {
	deleteQuery := "DELETE FROM task WHERE task_id = ?"
	_, err := DB.Exec(deleteQuery, taskId)
	return err
}

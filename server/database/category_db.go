package database

import (
	"todo-go-angular/server/models"
)

func CreateCategory(category *models.Category) (int, error) {
	insertQuery := "INSERT INTO category (user_id, category_name) VALUES (?, ?)"
	result, err := DB.Exec(insertQuery, category.UserID, category.CategoryName)
	if err != nil {
		return 0, err
	}

	categoryId, _ := result.LastInsertId()
	category.CategoryID = int(categoryId)

	return category.CategoryID, nil
}

func GetCategories() ([]models.Category, error) {
	selectQuery := "SELECT * FROM category"
	rows, err := DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(
			&category.CategoryID,
			&category.UserID,
			&category.CategoryName,
			&category.CreatedTimestamp,
			&category.ModifiedTimestamp,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func GetCategoriesByUserId(userId int) ([]models.Category, error) {
	selectQuery := "SELECT * FROM category WHERE user_id = ?"
	rows, err := DB.Query(selectQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.CategoryID, &category.UserID, &category.CategoryName, &category.CreatedTimestamp, &category.ModifiedTimestamp)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func GetCategoryById(categoryId int) (*models.Category, error) {
	selectQuery := "SELECT * FROM category WHERE category_id = ?"
	row := DB.QueryRow(selectQuery, categoryId)

	var category models.Category
	err := row.Scan(&category.CategoryID, &category.UserID, &category.CategoryName, &category.CreatedTimestamp, &category.ModifiedTimestamp)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func UpdateCategory(categoryId int, updatedCategory *models.Category) error {
	updateQuery := "UPDATE category SET category_name = ? WHERE category_id = ?"
	_, err := DB.Exec(updateQuery, updatedCategory.CategoryName, categoryId)
	return err
}

func DeleteCategory(categoryId int) error {
	deleteQuery := "DELETE FROM category WHERE category_id = ?"
	_, err := DB.Exec(deleteQuery, categoryId)
	return err
}

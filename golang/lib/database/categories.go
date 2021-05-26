package database

import (
	"strconv"
	"time"

	"github.com/askmuhammadamal/alta-store/models"
	"github.com/labstack/echo/v4"
)

func GetCategories() ([]models.Category, error) {
	var categories []models.Category

	if e := DB.Find(&categories).Error; e != nil {
		return nil, e
	}

	return categories, nil
}

func GetCategory(id int) ([]models.Category, error) {
	var category []models.Category

	if e := DB.Find(&category, id).Error; e != nil {
		return nil, e
	}

	return category, nil
}

func CreateCategory(c echo.Context) (interface{}, error) {
	category := models.Category{}
	c.Bind(&category)

	if e := DB.Save(&category).Error; e != nil {
		return nil, e
	}

	return category, nil
}

func UpdateCategory(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))

	category := models.Category{}
	c.Bind(&category)

	categoryDB := models.Category{}
	err := DB.Model(&category).Where("id = ?", id).Take(&categoryDB).UpdateColumns(
		map[string]interface{}{
			"name":        category.Name,
			"description": category.Description,
			"updated_at":  time.Now(),
		},
	).Error

	category.ID = categoryDB.ID
	category.CreatedAt = categoryDB.CreatedAt

	if err != nil {
		return nil, err
	}

	return category, nil
}

func DeleteCategory(id int) error {
	category := models.Category{}

	if err := DB.Delete(&category, id).Error; err != nil {
		return err
	}
	return nil
}

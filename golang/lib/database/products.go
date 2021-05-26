package database

import (
	"strconv"
	"time"

	"github.com/askmuhammadamal/alta-store/models"
	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) ([]models.Product, error) {
	var products []models.Product
	name := c.QueryParam("name")
	categoryName := c.QueryParam("categoryName")
	categoryId := c.QueryParam("categoryId")

	if name != "" {
		if e := DB.Model(&models.Product{}).Where("name LIKE ?", "%"+name+"%").Find(&products).Error; e != nil {
			return nil, e
		}
	} else if categoryName != "" {
		var category models.Category
		if e := DB.Model(&models.Category{}).Where("name LIKE ?", "%"+categoryName+"%").First(&category).Error; e != nil {
			return nil, e
		}
		if category.ID != 0 {
			if e := DB.Model(&models.Product{}).Where("category_id = ?", category.ID).Find(&products).Error; e != nil {
				return nil, e
			}
		}
	} else if categoryId != "" {
		var category models.Category
		if e := DB.Model(&models.Category{}).Where("id = ?", categoryId).First(&category).Error; e != nil {
			return nil, e
		}
		if category.ID != 0 {
			if e := DB.Model(&models.Product{}).Where("category_id = ?", category.ID).Find(&products).Error; e != nil {
				return nil, e
			}
		}
	} else {
		if e := DB.Find(&products).Error; e != nil {
			return nil, e
		}
	}

	return getCategoryData(products)
}

func GetProductDetail(id int) ([]models.Product, error) {
	var product []models.Product

	if err := DB.Find(&product, id).Error; err != nil {
		return nil, err
	}
	return getCategoryData(product)
}

func CreateProduct(c echo.Context) (interface{}, error) {

	product := models.Product{}
	c.Bind(&product)

	if e := DB.Save(&product).Error; e != nil {
		return nil, e
	}

	return getCategoryData([]models.Product{product})
}

func UpdateProduct(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))

	product := models.Product{}
	c.Bind(&product)

	productDB := models.Product{}
	err := DB.Model(&product).Where("id = ?", id).Take(&productDB).UpdateColumns(
		map[string]interface{}{
			"name":        product.Name,
			"description": product.Description,
			"stock":       product.Stock,
			"price":       product.Price,
			"category_id": product.CategoryID,
			"updated_at":  time.Now(),
		},
	).Error

	product.ID = productDB.ID
	product.CreatedAt = productDB.CreatedAt

	if err != nil {
		return nil, err
	}

	return getCategoryData([]models.Product{product})
}

func DeleteProduct(id int) error {
	product := models.Product{}

	if err := DB.Delete(&product, id).Error; err != nil {
		return err
	}
	return nil
}

func getCategoryData(products []models.Product) ([]models.Product, error) {
	if len(products) > 0 {
		for i := range products {
			err := DB.Model(&models.Category{}).Where("id = ?", products[i].CategoryID).Take(&products[i].Category).Error
			if err != nil {
				return nil, err
			}
		}
	}

	return products, nil
}

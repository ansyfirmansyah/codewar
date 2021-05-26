package controllers

import (
	"net/http"
	"strconv"

	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/labstack/echo/v4"
)

func GetCategoriesContoller(c echo.Context) error {
	categories, err := database.GetCategories()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get categories",
		"status":  "success",
		"data":    categories,
	})
}

func GetCategoryDetailContoller(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	category, err := database.GetCategory((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get category detail",
		"status":  "success",
		"data":    category,
	})
}

func CreateCategoryContoller(c echo.Context) error {
	category, e := database.CreateCategory(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": e.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create new category",
		"status":  "success",
		"data":    category,
	})
}

func UpdateCategoryContoller(c echo.Context) error {
	category, err := database.UpdateCategory(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success update category",
		"status":  "success",
		"data":    category,
	})
}

func DeleteCategoryContoller(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := database.DeleteCategory((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success delete category",
		"status":  "success",
	})
}

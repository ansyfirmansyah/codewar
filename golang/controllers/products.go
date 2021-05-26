package controllers

import (
	"net/http"
	"strconv"

	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/labstack/echo/v4"
)

func GetProductsContoller(c echo.Context) error {
	products, err := database.GetProducts(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get product",
		"status":  "success",
		"data":    products,
	})

}

func GetProductDetailContoller(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	product, err := database.GetProductDetail((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get product detail",
		"status":  "success",
		"data":    product,
	})
}

func CreateProductContoller(c echo.Context) error {
	product, e := database.CreateProduct(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": e.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create new product",
		"status":  "success",
		"data":    product,
	})
}

func UpdateProductContoller(c echo.Context) error {
	product, err := database.UpdateProduct(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success update product",
		"status":  "success",
		"data":    product,
	})
}

func DeleteProductContoller(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := database.DeleteProduct((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success delete product",
		"status":  "success",
	})
}

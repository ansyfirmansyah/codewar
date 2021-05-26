package controllers

import (
	"net/http"
	"strconv"

	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/labstack/echo/v4"
)

func CreateTransactionController(c echo.Context) error {
	transaction, e := database.CreateTransaction(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": e.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create new user",
		"status":  "success",
		"data":    transaction,
	})
}

func GetTransactionController(c echo.Context) error {
	transactions, err := database.GetTransactions(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get transactions",
		"status":  "success",
		"data":    transactions,
	})

}

func GetTransactionDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	transactions, err := database.GetTransaction((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get transaction detail",
		"status":  "success",
		"data":    transactions,
	})
}

func UpdateTransactionController(c echo.Context) error {
	transaction, err := database.EditTransaction(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success update transaction",
		"status":  "success",
		"data":    transaction,
	})
}

func DeleteTransactionController(c echo.Context) error {
	err := database.DeleteTransaction(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success delete transaction",
		"status":  "success",
	})
}

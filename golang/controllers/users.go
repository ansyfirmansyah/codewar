package controllers

import (
	"net/http"
	"strconv"

	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/askmuhammadamal/alta-store/models"
	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	token, e := database.LoginUsers(&user, user.Password)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": e.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "login success",
		"status":  "success",
		"data":    token,
	})
}

func GetUserController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get user",
		"status":  "success",
		"data":    users,
	})

}

func GetUserDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	users, err := database.GetUser((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get user detail",
		"status":  "success",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {
	user, e := database.CreateUsers(c)

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
		"data":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	user, err := database.EditUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success update user",
		"status":  "success",
		"data":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := database.DeleteUser((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success delete user",
		"status":  "success",
	})
}

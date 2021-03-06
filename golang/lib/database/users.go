package database

import (
	"errors"
	"strconv"
	"time"

	"github.com/askmuhammadamal/alta-store/middlewares"
	"github.com/askmuhammadamal/alta-store/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginUsers(user *models.User, password string) (interface{}, error) {
	var err error
	token := models.Token{}
	if err = DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		return nil, errors.New("incorrect email or password")
	}

	err = VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.New("incorrect email or password")
	}

	token.Data, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func validateEmail(email string, userId int) error {
	user := models.User{}
	if userId > 0 {
		if err := DB.Where("email = ? AND id <> ?", email, userId).First(&user).Error; err != nil {
			return nil
		}
	} else {
		if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
			return nil
		}
	}
	return errors.New("email already exists")
}

func GetUsers() ([]models.User, error) {
	var users []models.User

	if e := DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUser(userId int) ([]models.User, error) {
	var users []models.User

	if err := DB.Find(&users, userId).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUsers(c echo.Context) (interface{}, error) {

	user := models.User{}
	c.Bind(&user)
	emailExist := validateEmail(user.Email, 0)
	if emailExist != nil {
		return nil, emailExist
	}

	hashPassword, err := Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashPassword)

	if e := DB.Save(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func EditUser(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))

	// binding data
	user := models.User{}
	c.Bind(&user)

	emailExist := validateEmail(user.Email, id)
	if emailExist != nil {
		return nil, emailExist
	}

	hashPassword, errHash := Hash(user.Password)
	if errHash != nil {
		return nil, errHash
	}
	user.Password = string(hashPassword)

	userDB := models.User{}
	err := DB.Model(&user).Where("id = ?", id).Take(&userDB).UpdateColumns(
		map[string]interface{}{
			"full_name":     user.FullName,
			"phone_number":  user.PhoneNumber,
			"email":         user.Email,
			"password":      user.Password,
			"gender":        user.Gender,
			"date_of_birth": user.DateOfBirth,
			"district":      user.District,
			"sub_district":  user.SubDistrict,
			"address":       user.Address,
			"updated_at":    time.Now(),
		},
	).Error
	user.ID = userDB.ID
	user.CreatedAt = userDB.CreatedAt
	user.Role = userDB.Role

	if err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(userId int) error {
	// binding data
	user := models.User{}

	if err := DB.Delete(&user, userId).Error; err != nil {
		return err
	}
	return nil
}

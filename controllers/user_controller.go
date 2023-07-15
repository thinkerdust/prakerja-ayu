package controllers

import (
	"net/http"
	"prakerja6/config"
	"prakerja6/models"

	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal memasukkan data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success memasukkan data", Status: true, Data: user,
	})
}

func DetailUserController(c echo.Context) error {
	id := c.Param("id")

	db := config.DB

	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "User not found", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success", Status: true, Data: &user,
	})
}

func UserController(c echo.Context) error {
	var users []models.User

	result := config.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal mendapatkan data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success mendapatkan data", Status: true, Data: users,
	})
}

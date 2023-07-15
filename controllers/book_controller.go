package controllers

import (
	"net/http"
	"prakerja6/config"
	"prakerja6/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	result := config.DB.Create(&book)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Failed to save data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success", Status: true, Data: &book,
	})
}

func UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Message: "Invalid Book ID", Status: false, Data: nil})
	}

	// Check if the book exists
	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{Message: "Book not found", Status: false, Data: nil})
	}

	// Parse the request body into a Book object
	updatedBook := new(models.Book)
	if err := c.Bind(updatedBook); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Message: "Invalid data format", Status: false, Data: nil})
	}

	// Update the book data
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author

	// Save the updated book back to the database
	result = config.DB.Save(&book)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Message: "Failed to update Book", Status: false, Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Message: "Success", Status: false, Data: book})
}

func GetBooks(c echo.Context) error {
	var books []models.Book

	result := config.DB.Find(&books)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error while fetching data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success", Status: true, Data: books,
	})

}

func GetBooksById(c echo.Context) error {
	id := c.Param("id")
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Message: "Data not found", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success", Status: true, Data: &book,
	})
}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: "Invalid book ID", Status: false, Data: nil,
		})
	}

	// Check if the book exists
	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Message: "Book not found", Status: false, Data: nil,
		})
	}

	// Delete the book from the database
	result = config.DB.Delete(&book)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Failed to delete book", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Book deleted successfully", Status: true, Data: book,
	})
}

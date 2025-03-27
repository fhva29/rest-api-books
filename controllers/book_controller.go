package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"rest_api_books/models"
	"rest_api_books/requests"
	"rest_api_books/responses"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBook(ctx *gin.Context) {
	var request requests.CrateBookRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		responses.SendError(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := request.Validate(); err != nil {
		responses.SendError(ctx, http.StatusBadRequest, "Request validation error", err)
		return
	}

	var author models.Author
	if err := db.First(&author, request.AuthorID).Error; err != nil {
		responses.SendError(ctx, http.StatusBadRequest, "Author not found")
		return
	}

	book := models.Book{
		Title:       request.Title,
		Genre:       request.Genre,
		Year:        request.Year,
		Description: request.Description,
		AuthorID:    request.AuthorID,
	}

	if err := db.Create(&book).Error; err != nil {
		responses.SendError(ctx, http.StatusInternalServerError, "Database error")
		return
	}

	if err := db.Preload("Author").First(&book, book.ID).Error; err != nil {
		responses.SendError(ctx, http.StatusInternalServerError, "Error loading author")
		return
	}

	responses.SendSuccess(ctx, responses.ToBookResponse(book))
}

func ListBooks(ctx *gin.Context) {
	var books []models.Book

	if err := db.Preload("Author").Find(&books).Error; err != nil {
		responses.SendError(ctx, http.StatusInternalServerError, "Error finding books.")
		return
	}

	responses.SendSuccess(ctx, responses.ToBookResponseList(books))
}

func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	bookID, err := strconv.Atoi(id)
	if err != nil {
		responses.SendError(ctx, http.StatusBadRequest, "id must be a valid integer.")
		return
	}

	var book models.Book
	if err := db.Preload("Author").First(&book, bookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Book with id %d not found", bookID))
		} else {
			responses.SendError(ctx, http.StatusInternalServerError, "Database error")
		}
		return
	}

	responses.SendSuccess(ctx, responses.ToBookResponse(book))
}

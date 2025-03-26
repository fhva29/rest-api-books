package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"rest_api_books/models"
	"rest_api_books/requests"
	"rest_api_books/responses"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAuthor(ctx *gin.Context) {
	var request requests.CreateAuthorRequest

	// Faz o parse do request
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
			Code:  http.StatusBadRequest,
		})
		return
	}

	// Valida o request
	if err := request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "Request validation error",
			Code:    http.StatusBadRequest,
			Details: err,
		})
		return
	}

	// Parser para a data de nascimento
	dataNascimento, err := time.Parse("02/01/2006", request.DataNascimento)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "Invalid date format. Expected: DD/MM/YYYY",
			Code:  http.StatusBadRequest,
		})
		return
	}

	// Cria o author
	author := models.Author{
		Name:           request.Name,
		Nacionality:    request.Nacionality,
		DataNascimento: dataNascimento,
	}

	// Salva o author no banco de dados
	if err := db.Create(&author).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
			Code:  http.StatusInternalServerError,
		})
		return
	}

	responses.SendSuccess(ctx, responses.ToAuthorResponse(author))
}

func ListAuthors(ctx *gin.Context) {
	var authors []models.Author

	if err := db.Find(&authors).Error; err != nil {
		responses.SendError(ctx, http.StatusInternalServerError, "Error finding  authors.")
		return
	}

	responses.SendSuccess(ctx, responses.ToAuthorResponseList(authors))
}

func GetAuthor(ctx *gin.Context) {
	id := ctx.Param("id")

	authorID, err := strconv.Atoi(id)
	if err != nil {
		responses.SendError(ctx, http.StatusBadRequest, "ID must be a valid integer.")
		return
	}

	var author models.Author

	if err := db.First(&author, authorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Author with id %d not found.", authorID))
		} else {
			responses.SendError(ctx, http.StatusInternalServerError, "Database error.")
		}
		return
	}

	responses.SendSuccess(ctx, responses.ToAuthorResponse(author))
}

func UpdateAuthor(ctx *gin.Context) {
	id := ctx.Param("id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		responses.SendError(ctx, http.StatusBadRequest, "ID must be a valid integer")
		return
	}

	var request requests.CreateAuthorRequest
	// Faz o parse do request
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
			Code:  http.StatusBadRequest,
		})
		return
	}

	// Valida o request
	if err := request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "Request validation error",
			Code:    http.StatusBadRequest,
			Details: err,
		})
		return
	}

	// Parser para a data de nascimento
	dataNascimento, err := time.Parse("02/01/2006", request.DataNascimento)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "Invalid date format. Expected: DD/MM/YYYY",
			Code:  http.StatusBadRequest,
		})
		return
	}

	// Recupera o autor
	author := models.Author{}

	if err := db.First(&author, authorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Author with id %d not found.", authorID))
		} else {
			responses.SendError(ctx, http.StatusInternalServerError, "Database error.")
		}
		return
	}

	newAuthor := models.Author{
		Name:           request.Name,
		Nacionality:    request.Nacionality,
		DataNascimento: dataNascimento,
	}

	// Update no autor
	if err := db.Model(&author).Updates(newAuthor).Error; err != nil {
		responses.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error updating book with id: %d", authorID))
		return
	}

	responses.SendSuccess(ctx, responses.ToAuthorResponse(author))
}

func DeleteAuthor(ctx *gin.Context) {
	id := ctx.Param("id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		responses.SendError(ctx, http.StatusBadRequest, "Id must be a valid integer.")
		return
	}

	author := models.Author{}
	if err := db.First(&author, authorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Author with id %d not found", authorID))
		} else {
			responses.SendError(ctx, http.StatusInternalServerError, "Database error.")
		}
		return
	}

	if err := db.Delete(&author).Error; err != nil {
		responses.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting author with id %d", authorID))
		return
	}

	responses.SendSuccess(ctx, responses.ToAuthorResponse(author))
}

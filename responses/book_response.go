package responses

import "rest_api_books/models"

type BookResponse struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
	Genre       string         `json:"genre"`
	Year        int            `json:"year"`
	Description *string        `json:"description,omitempty"`
	Author      AuthorResponse `json:"author"`
}

func ToBookResponse(book models.Book) BookResponse {
	return BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Genre:       book.Genre,
		Year:        book.Year,
		Description: book.Description,
		Author:      ToAuthorResponse(book.Author),
	}
}

func ToBookResponseList(books []models.Book) []BookResponse {
	var result []BookResponse
	for _, book := range books {
		result = append(result, ToBookResponse(book))
	}
	return result
}

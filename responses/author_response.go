package responses

import "rest_api_books/models"

type AuthorResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Nacionality    string `json:"nacionality"`
	DataNascimento string `json:"data_nascimento"`
}

func ToAuthorResponse(author models.Author) AuthorResponse {
	return AuthorResponse{
		ID:             author.ID,
		Name:           author.Name,
		Nacionality:    author.Nacionality,
		DataNascimento: author.DataNascimento.Format("02/01/2006"),
	}
}

func ToAuthorResponseList(authors []models.Author) []AuthorResponse {
	var result []AuthorResponse
	for _, a := range authors {
		result = append(result, ToAuthorResponse(a))
	}
	return result
}

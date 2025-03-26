package requests

import "strings"

type CreateAuthorRequest struct {
	Name           string `json:"name"`
	Nacionality    string `json:"nacionality"`
	DataNascimento string `json:"data_nascimento"`
}

func (r *CreateAuthorRequest) Validate() map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(r.Name) == "" {
		errors["name"] = "Name is required"
	} else if len(r.Name) > 50 {
		errors["name"] = "Name must be less than 50 characters"
	} else if len(r.Name) < 3 {
		errors["name"] = "Name must be greater than 3 characters"
	}

	if strings.TrimSpace(r.Nacionality) == "" {
		errors["nacionality"] = "Nacionality is required"
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

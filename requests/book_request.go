package requests

import (
	"fmt"
	"strings"
	"time"
)

type CrateBookRequest struct {
	Title       string  `json:"title"`
	Genre       string  `json:"genre"`
	Year        int     `json:"year"`
	Description *string `json:"description"`
	AuthorID    uint    `json:"author_id"`
}

func (r *CrateBookRequest) Validate() map[string]string {
	errors := make(map[string]string)

	if err := validateStringField(r.Title, "title", 3, 50); err != "" {
		errors["title"] = err
	}

	if err := validateStringField(r.Genre, "genre", 3, 50); err != "" {
		errors["genre"] = err
	}

	currentYear := time.Now().Year()
	if r.Year < 1000 || r.Year > currentYear {
		errors["year"] = fmt.Sprintf("Year must be between 1000 and %d", currentYear)
	}

	if r.Description != nil {
		if err := validateStringField(*r.Description, "description", 10, 500); err != "" {
			errors["description"] = err
		}
	}

	if r.AuthorID == 0 {
		errors["author_id"] = "Author ID is required"
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func validateStringField(fieldValue string, fieldName string, min int, max int) string {
	value := strings.TrimSpace(fieldValue)
	if value == "" {
		return fmt.Sprintf("%s is required", fieldName)
	}
	if len(value) > max {
		return fmt.Sprintf("%s must be less than %d characters", fieldName, max)
	}
	if len(value) < min {
		return fmt.Sprintf("%s must be at least %d characters", fieldName, min)
	}
	return ""
}

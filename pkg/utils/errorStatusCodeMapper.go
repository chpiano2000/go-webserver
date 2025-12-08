package utils

import (
	"net/http"

	"github.com/go-webserver/internal/domains"
)

func MapErrorToStatus(err domains.Error) int {
	switch err.Code {
	case domains.ErrRecipeNotFound.Code:
		return http.StatusNotFound
	case domains.ErrRecipeInvalidData.Code,
		domains.ErrInvalidInput.Code,
		domains.ErrMissingField.Code:
		return http.StatusBadRequest
	case domains.ErrDatabaseConnection.Code:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

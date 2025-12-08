package domains

import "fmt"

// Error represents a domain error with code and message
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// NewError creates a new domain error
func NewError(code, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}

// Common errors
var (
	ErrNone = Error{Code: "", Message: ""}

	// Recipe errors
	ErrRecipeNotFound = NewError(
		"Recipe.NotFound",
		"The recipe with the specified identifier was not found",
	)
	ErrRecipeAlreadyExists = NewError(
		"Recipe.AlreadyExists",
		"A recipe with the same name already exists",
	)
	ErrRecipeInvalidData = NewError(
		"Recipe.InvalidData",
		"The recipe data is invalid",
	)

	// Validation errors
	ErrInvalidInput = NewError(
		"Validation.InvalidInput",
		"The provided input is invalid",
	)
	ErrMissingField = NewError(
		"Validation.MissingField",
		"Required field is missing",
	)

	// Infrastructure errors
	ErrDatabaseConnection = NewError(
		"Infrastructure.DatabaseConnection",
		"Database connection failed",
	)
)

// Recipe-specific errors constructor
type RecipeErrors struct{}

var RecipeError = RecipeErrors{}

func (RecipeErrors) NotFound() Error {
	return ErrRecipeNotFound
}

func (RecipeErrors) InvalidName(name string) Error {
	return NewError(
		"Recipe.InvalidName",
		fmt.Sprintf("Recipe name '%s' is invalid", name),
	)
}

func (RecipeErrors) InvalidPrepTime() Error {
	return NewError(
		"Recipe.InvalidPrepTime",
		"Preparation time must be positive",
	)
}

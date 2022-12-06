// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package v1

// Error defines model for Error.
type Error struct {
	Error *string `json:"error,omitempty"`
}

// Errors defines model for Errors.
type Errors struct {
	Errors *map[string][]string `json:"errors,omitempty"`
}

// User defines model for User.
type User struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// UserCreate defines model for UserCreate.
type UserCreate struct {
	Email    *string `json:"email,omitempty" validate:"required,email"`
	Name     *string `json:"name,omitempty" validate:"required"`
	Password *string `json:"password,omitempty" validate:"required,alphanum,len=8"`
}

// UserID defines model for UserID.
type UserID = string

// ErrorNotFound defines model for ErrorNotFound.
type ErrorNotFound = Error

// ErrorUnexpected defines model for ErrorUnexpected.
type ErrorUnexpected = Error

// ErrorValidation defines model for ErrorValidation.
type ErrorValidation = Errors

// UserResponse defines model for UserResponse.
type UserResponse struct {
	Data *User `json:"data,omitempty"`
}

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = UserCreate

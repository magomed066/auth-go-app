package users

import "errors"

var ErrUserAlreadyExists = errors.New("user already exists")
var UserFieldsValidation = errors.New("Fields: Fist name, Last name, Email and Password cannot be empty")
var ErrInvalidCredentials = errors.New("Invalid email or password")
var ErrPasswordTooLong = errors.New("Password is too long")
var ErrUserPassword = errors.New("Password is too long")


var FieldErrorMessages = map[string]string{
    "Email.required":    "Email is required",
    "Email.email":       "Must be a valid email",
    "Password.min":      "Password must be at least 6 characters",
    "FirstName.required":           "First name cannot be empty",
    "LastName.required":           "Last name cannot be empty",
}

var LoginValidationErrors = map[string]string{
    "Email.required":    "Email is required",
    "Email.email":       "Must be a valid email",
    "Password.required": "Password is required",
}

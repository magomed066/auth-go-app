package request

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var bodyValidator = validator.New()

// Validate checks a request struct and returns readable field-level errors.
// Add or change messages here to customize validation feedback.
func Validate(data any, messages map[string]string) error {
    err := bodyValidator.Struct(data)
    if err == nil {
        return nil
    }

    var validationErrors validator.ValidationErrors
    if !errors.As(err, &validationErrors) {
        return err
    }

    fieldMessages := make([]string, 0, len(validationErrors))
    for _, fieldErr := range validationErrors {
        key := fieldErr.StructField() + "." + fieldErr.Tag()
        message, ok := messages[key]
        if !ok {
            message = fieldErr.Field() + " is invalid"
        }
        fieldMessages = append(fieldMessages, message)
    }

    return errors.New(strings.Join(fieldMessages, "; "))
}

package main

import (
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func Error(e error) string {
	return e.Error()
}
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func ValidateAge(age int) error {
	if age < 0 || age > 150 {
		return &ValidationError{Field: "age", Message: "Invalid. Must be between 0 and 150"}
	}
	return nil
}

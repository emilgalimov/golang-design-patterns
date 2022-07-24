package Errors

import "fmt"

type PhoneTypeNotFoundError struct {
	typeID int
}

func (e *PhoneTypeNotFoundError) Error() string {
	return fmt.Sprintf("Error type with ID = %v not found", e.typeID)
}

func NewPhoneTypeNotFoundError(typeID int) *PhoneTypeNotFoundError {
	return &PhoneTypeNotFoundError{
		typeID: typeID,
	}
}

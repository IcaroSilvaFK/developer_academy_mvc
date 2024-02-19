package utils

import "github.com/google/uuid"

func IsValidId(id string) bool {

	err := uuid.Validate(id)

	return err == nil
}

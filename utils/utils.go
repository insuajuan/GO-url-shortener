package utils

import "github.com/google/uuid"

func GenerateShortKey() string {
	// Generate a random UUID
	u, err := uuid.NewRandom()
	if err != nil {
		// Handle the error, e.g., by returning a default key or an error message
		return "default_key"
	}

	// Convert the UUID to a string and take the first 6 characters
	shortKey := u.String()[:6]
	return shortKey
}

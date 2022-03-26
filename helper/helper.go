package helper

import "strings"

func ValidateUserInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// User Input Validation
	validatedName := len(firstName) >= 2 && len(lastName) >= 2
	validatedEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	validatedTickets := userTickets > 0 && userTickets <= remainingTickets

	return validatedName, validatedEmail, validatedTickets
}

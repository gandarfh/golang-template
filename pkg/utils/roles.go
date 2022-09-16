package utils

import (
	"fmt"

	"goapi/pkg/permissions"
)

// VerifyRole func for verifying a given role.
func VerifyRole(role string) (string, error) {
	// Switch given role.
	switch role {
	case permissions.AdminRole:
		// Nothing to do, verified successfully.
	case permissions.ModeratorRole:
		// Nothing to do, verified successfully.
	case permissions.UserRole:
		// Nothing to do, verified successfully.
	case permissions.ReadOnlyRole:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}

package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

type success struct {
	Status int `json:"status"`
}

// @Description Health check test.
// @Summary Health check validation
// @Tags Healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} success
// @Router /v1/healthcheck [get]
func GetHealthcheck(c *fiber.Ctx) error {
	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK})
}

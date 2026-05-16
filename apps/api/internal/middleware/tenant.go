package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func TenantMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// In a real app, you'd get this from JWT claims
		// or a custom header for API keys.
		tenantID := c.Get("X-Tenant-ID")

		if tenantID == "" {
			// For demonstration, we'll allow it but in production
			// you should return 401/403 if required.
			// return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Tenant context missing"})
		}

		c.Locals("tenantID", tenantID)
		return c.Next()
	}
}

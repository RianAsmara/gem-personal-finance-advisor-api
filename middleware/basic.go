package middleware

import (
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// BasicAuthMiddleware checks for valid basic authentication credentials.
func BasicAuthMiddleware(username, password string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		auth := c.Get("Authorization")
		if auth == "" {
			// Request for basic auth if no Authorization header is provided
			c.Set("WWW-Authenticate", `Basic realm="Restricted"`)
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		// Decode the base64 encoded credentials
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Basic" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		decoded, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		// Check if the credentials match the expected username and password
		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 || credentials[0] != username || credentials[1] != password {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		// If everything is correct, allow access to the route
		return c.Next()
	}
}

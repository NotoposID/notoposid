package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notopos/api/internal/common"
	"github.com/notopos/api/internal/config"
	"github.com/notopos/api/internal/helpers"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{cfg: cfg}
}

type LoginRequest struct {
	Email    string `json:"email" example:"admin@notopos.com"`
	Password string `json:"password" example:"password123"`
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} common.Response
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request", nil)
	}

	// Placeholder logic (In real app, check DB)
	// For demo: admin@notopos.com / password123
	if req.Email != "admin@notopos.com" || req.Password != "password123" {
		return common.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials", nil)
	}

	token, err := helpers.GenerateToken("user-uuid", "tenant-uuid", "admin", h.cfg.JWTSecret)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Token generation failed", nil)
	}

	return common.SuccessResponse(c, "Login successful", fiber.Map{
		"token": token,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

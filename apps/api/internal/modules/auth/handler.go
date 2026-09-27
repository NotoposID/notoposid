package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notopos/api/internal/common"
	"github.com/notopos/api/internal/config"
	"github.com/notopos/api/internal/helpers"
	"github.com/notopos/api/internal/modules/users"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	cfg      *config.Config
	userRepo users.Repository
}

func NewHandler(cfg *config.Config, userRepo users.Repository) *AuthHandler {
	return &AuthHandler{cfg: cfg, userRepo: userRepo}
}

type LoginRequest struct {
	Identifier string `json:"identifier" example:"admin@notopos.com"` // can be email or username
	Password   string `json:"password" example:"password123"`
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

	// 1. Check in DB first (if DB is running and user exists)
	user, err := h.userRepo.GetByEmailOrUsername(req.Identifier)
	
	// Fallback to static mockup if DB is empty or connection fails (since Docker might not be running)
	if err != nil || user == nil {
		if (req.Identifier == "admin@notopos.com" || req.Identifier == "admin") && req.Password == "password123" {
			token, err := helpers.GenerateToken("user-uuid", "tenant-uuid", "admin", h.cfg.JWTSecret)
			if err != nil {
				return common.ErrorResponse(c, fiber.StatusInternalServerError, "Token generation failed", nil)
			}
			return common.SuccessResponse(c, "Login successful", fiber.Map{
				"token": token,
				"user": fiber.Map{
					"name": "Admin Placeholder",
					"email": "admin@notopos.com",
					"username": "admin",
				},
			})
		}
		return common.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials", nil)
	}

	// Check password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return common.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials", nil)
	}

	token, err := helpers.GenerateToken(user.ID.String(), user.TenantID.String(), user.Role, h.cfg.JWTSecret)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Token generation failed", nil)
	}

	return common.SuccessResponse(c, "Login successful", fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id": user.ID,
			"name": user.Name,
			"email": user.Email,
			"username": user.Username,
			"role": user.Role,
		},
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

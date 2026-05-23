package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/notopos/api/internal/common"
)

type Handler struct {
	repo    Repository
	service Service
}

func NewHandler(repo Repository, service Service) *Handler {
	return &Handler{
		repo:    repo,
		service: service,
	}
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"john@example.com"`
	Password string `json:"password" validate:"required,min=6" example:"secret123"`
	Role     string `json:"role" validate:"required" example:"staff"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required" example:"John Updated"`
	Role  string `json:"role" validate:"required" example:"admin"`
	Email string `json:"email" validate:"required,email" example:"john@example.com"`
}

// Create godoc
// @Summary Create a new user
// @Description Create a new user for the current tenant
// @Tags Users
// @Accept json
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param request body CreateUserRequest true "User details"
// @Success 201 {object} common.Response
// @Router /users [post]
func (h *Handler) Create(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)

	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Format tipe data tidak valid", err.Error())
	}

	if errs := common.ValidateStruct(&req); errs != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Validasi gagal", errs)
	}

	// validation unique email from request body
	if _, err := h.repo.GetByEmail(req.Email, tenantID); err == nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Email sudah terdaftar", nil)
	}

	passwordHash, err := h.service.HashPassword(req.Password)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to process password", nil)
	}

	user := &User{
		TenantID:     tenantID,
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
		Role:         req.Role,
	}

	if err := h.repo.Create(user); err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
	}

	return common.SuccessResponse(c, "User created successfully", user)
}

// GetAll godoc
// @Summary List all users
// @Description Get all users for the current tenant
// @Tags Users
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Success 200 {object} common.Response
// @Router /users [get]
func (h *Handler) GetAll(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)

	users, err := h.repo.GetAll(tenantID)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch users", err.Error())
	}

	return common.SuccessResponse(c, "Users retrieved successfully", users)
}

// GetByID godoc
// @Summary Get user by ID
// @Description Get a specific user by their UUID
// @Tags Users
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param id path string true "User ID"
// @Success 200 {object} common.Response
// @Router /users/{id} [get]
func (h *Handler) GetByID(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)
	id, _ := uuid.Parse(c.Params("id"))

	user, err := h.repo.GetByID(id, tenantID)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusNotFound, "User not found", err.Error())
	}

	return common.SuccessResponse(c, "User retrieved successfully", user)
}

// Update godoc
// @Summary Update user
// @Description Update user details
// @Tags Users
// @Accept json
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param id path string true "User ID"
// @Param request body UpdateUserRequest true "Updated user details"
// @Success 200 {object} common.Response
// @Router /users/{id} [put]
func (h *Handler) Update(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)
	id, _ := uuid.Parse(c.Params("id"))

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Format tipe data tidak valid", err.Error())
	}

	if errs := common.ValidateStruct(&req); errs != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Validasi gagal", errs)
	}

	user, err := h.repo.GetByID(id, tenantID)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusNotFound, "User not found", err.Error())
	}

	user.Name = req.Name
	user.Role = req.Role
	user.Email = req.Email

	if err := h.repo.Update(user); err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update user", err.Error())
	}

	return common.SuccessResponse(c, "User updated successfully", user)
}

// Delete godoc
// @Summary Delete user
// @Description Remove a user from the tenant
// @Tags Users
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param id path string true "User ID"
// @Success 200 {object} common.Response
// @Router /users/{id} [delete]
func (h *Handler) Delete(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)
	id, _ := uuid.Parse(c.Params("id"))

	if err := h.repo.Delete(id, tenantID); err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete user", err.Error())
	}

	return common.SuccessResponse(c, "User deleted successfully", nil)
}

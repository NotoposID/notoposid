package categories

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

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required" example:"Beverages"`
	Description string `json:"description" example:"All kinds of drinks"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name" validate:"required" example:"Beverages Updated"`
	Description string `json:"description" example:"Updated description"`
}

// Create godoc
// @Summary Create a new category
// @Description Create a new product category for the current tenant
// @Tags Categories
// @Accept json
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param request body CreateCategoryRequest true "Category details"
// @Success 201 {object} common.Response
// @Router /categories [post]
func (h *Handler) Create(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)

	var req CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Format tipe data tidak valid", err.Error())
	}

	if errs := common.ValidateStruct(&req); errs != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Validation failed", errs)
	}

	category := &Category{
		TenantID:    tenantID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.repo.Create(category); err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create category", err.Error())
	}

	return common.SuccessResponse(c, "Category created successfully", category)
}

// GetAll godoc
// @Summary List all categories
// @Description Get all categories for the current tenant
// @Tags Categories
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Success 200 {object} common.Response
// @Router /categories [get]
func (h *Handler) GetAll(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)

	categories, err := h.repo.GetAll(tenantID)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch categories", err.Error())
	}

	return common.SuccessResponse(c, "Categories retrieved successfully", categories)
}

// GetByID godoc
// @Summary Get category by ID
// @Description Get a specific category by its UUID
// @Tags Categories
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param id path string true "Category ID"
// @Success 200 {object} common.Response
// @Router /categories/{id} [get]
func (h *Handler) GetByID(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Invalid category ID", err.Error())
	}

	category, err := h.repo.GetByID(id, tenantID)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusNotFound, "Category not found", err.Error())
	}

	return common.SuccessResponse(c, "Category retrieved successfully", category)
}

// Update godoc
// @Summary Update category
// @Description Update category details
// @Tags Categories
// @Accept json
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param id path string true "Category ID"
// @Param request body UpdateCategoryRequest true "Updated category details"
// @Success 200 {object} common.Response
// @Router /categories/{id} [put]
func (h *Handler) Update(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Invalid category ID", err.Error())
	}

	var req UpdateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Format tipe data tidak valid", err.Error())
	}

	if errs := common.ValidateStruct(&req); errs != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Validation failed", errs)
	}

	category, err := h.repo.GetByID(id, tenantID)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusNotFound, "Category not found", err.Error())
	}

	category.Name = req.Name
	category.Description = req.Description

	if err := h.repo.Update(category); err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update category", err.Error())
	}

	return common.SuccessResponse(c, "Category updated successfully", category)
}

// Delete godoc
// @Summary Delete category
// @Description Remove a category from the tenant
// @Tags Categories
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param id path string true "Category ID"
// @Success 200 {object} common.Response
// @Router /categories/{id} [delete]
func (h *Handler) Delete(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Invalid category ID", err.Error())
	}

	if err := h.repo.Delete(id, tenantID); err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete category", err.Error())
	}

	return common.SuccessResponse(c, "Category deleted successfully", nil)
}

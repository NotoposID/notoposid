package products

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

type CreateProductRequest struct {
	CategoryID    *string `json:"category_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name          string  `json:"name" validate:"required" example:"Cappuccino"`
	SKU           string  `json:"sku" example:"CAP-01"`
	Barcode       string  `json:"barcode" example:"123456789012"`
	Description   string  `json:"description" example:"Hot cappuccino with latte art"`
	Price         float64 `json:"price" validate:"required,gt=0" example:"35000.00"`
	CostPrice     float64 `json:"cost_price" example:"15000.00"`
	StockQuantity int     `json:"stock_quantity" example:"100"`
	MinStockLevel int     `json:"min_stock_level" example:"20"`
	ImageURL      string  `json:"image_url" example:"https://example.com/cappuccino.jpg"`
}

type UpdateProductRequest struct {
	CategoryID    *string `json:"category_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name          string  `json:"name" validate:"required" example:"Cappuccino Large"`
	SKU           string  `json:"sku" example:"CAP-02"`
	Barcode       string  `json:"barcode" example:"123456789013"`
	Description   string  `json:"description" example:"Large hot cappuccino"`
	Price         float64 `json:"price" validate:"required,gt=0" example:"45000.00"`
	CostPrice     float64 `json:"cost_price" example:"20000.00"`
	StockQuantity int     `json:"stock_quantity" example:"50"`
	MinStockLevel int     `json:"min_stock_level" example:"10"`
	ImageURL      string  `json:"image_url" example:"https://example.com/cappuccino-lg.jpg"`
}

// Create godoc
// @Summary Create a new product
// @Description Create a new product for the current tenant
// @Tags Products
// @Accept json
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param request body CreateProductRequest true "Product details"
// @Success 201 {object} common.Response
// @Router /products [post]
func (h *Handler) Create(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)

	var req CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Format tipe data tidak valid", err.Error())
	}

	if errs := common.ValidateStruct(&req); errs != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Validation failed", errs)
	}

	product := &Product{
		TenantID:      tenantID,
		Name:          req.Name,
		SKU:           req.SKU,
		Barcode:       req.Barcode,
		Description:   req.Description,
		Price:         req.Price,
		CostPrice:     req.CostPrice,
		StockQuantity: req.StockQuantity,
		MinStockLevel: req.MinStockLevel,
		ImageURL:      req.ImageURL,
	}

	if req.CategoryID != nil && *req.CategoryID != "" {
		categoryID, err := uuid.Parse(*req.CategoryID)
		if err == nil {
			product.CategoryID = &categoryID
		}
	}

	if err := h.service.CreateProduct(product); err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return common.SuccessResponse(c, "Product created successfully", product)
}

// GetAll godoc
// @Summary List all products
// @Description Get all products for the current tenant
// @Tags Products
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Success 200 {object} common.Response
// @Router /products [get]
func (h *Handler) GetAll(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)

	products, err := h.repo.GetAll(tenantID)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch products", err.Error())
	}

	return common.SuccessResponse(c, "Products retrieved successfully", products)
}

// GetByID godoc
// @Summary Get product by ID
// @Description Get a specific product by its UUID
// @Tags Products
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param id path string true "Product ID"
// @Success 200 {object} common.Response
// @Router /products/{id} [get]
func (h *Handler) GetByID(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID", err.Error())
	}

	product, err := h.repo.GetByID(id, tenantID)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusNotFound, "Product not found", err.Error())
	}

	return common.SuccessResponse(c, "Product retrieved successfully", product)
}

// Update godoc
// @Summary Update product
// @Description Update product details
// @Tags Products
// @Accept json
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param id path string true "Product ID"
// @Param request body UpdateProductRequest true "Updated product details"
// @Success 200 {object} common.Response
// @Router /products/{id} [put]
func (h *Handler) Update(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID", err.Error())
	}

	var req UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Format tipe data tidak valid", err.Error())
	}

	if errs := common.ValidateStruct(&req); errs != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Validation failed", errs)
	}

	product, err := h.repo.GetByID(id, tenantID)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusNotFound, "Product not found", err.Error())
	}

	product.Name = req.Name
	product.SKU = req.SKU
	product.Barcode = req.Barcode
	product.Description = req.Description
	product.Price = req.Price
	product.CostPrice = req.CostPrice
	product.StockQuantity = req.StockQuantity
	product.MinStockLevel = req.MinStockLevel
	product.ImageURL = req.ImageURL

	if req.CategoryID != nil && *req.CategoryID != "" {
		categoryID, err := uuid.Parse(*req.CategoryID)
		if err == nil {
			product.CategoryID = &categoryID
		}
	} else {
		product.CategoryID = nil
	}

	if err := h.repo.Update(product); err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update product", err.Error())
	}

	return common.SuccessResponse(c, "Product updated successfully", product)
}

// Delete godoc
// @Summary Delete product
// @Description Remove a product from the tenant
// @Tags Products
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param id path string true "Product ID"
// @Success 200 {object} common.Response
// @Router /products/{id} [delete]
func (h *Handler) Delete(c *fiber.Ctx) error {
	tenantIDStr := c.Locals("tenantID").(string)
	tenantID, _ := uuid.Parse(tenantIDStr)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID", err.Error())
	}

	if err := h.repo.Delete(id, tenantID); err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete product", err.Error())
	}

	return common.SuccessResponse(c, "Product deleted successfully", nil)
}

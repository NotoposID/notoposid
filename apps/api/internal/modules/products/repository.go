package products

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `json:"name"`
}

func (Category) TableName() string {
	return "categories"
}

type Tenant struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `json:"name"`
}

func (Tenant) TableName() string {
	return "tenants"
}

type Product struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TenantID      uuid.UUID  `gorm:"type:uuid;not null;index"`
	Tenant        Tenant     `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
	CategoryID    *uuid.UUID `gorm:"type:uuid"`
	Category      *Category  `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Name          string     `gorm:"size:255;not null"`
	SKU           string     `gorm:"size:100"`
	Barcode       string     `gorm:"size:100"`
	Description   string     `gorm:"type:text"`
	Price         float64    `gorm:"type:decimal(15,2);not null"`
	CostPrice     float64    `gorm:"type:decimal(15,2)"`
	StockQuantity int        `gorm:"default:0"`
	MinStockLevel int        `gorm:"default:10"`
	ImageURL      string     `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Repository interface {
	Create(product *Product) error
	GetAll(tenantID uuid.UUID) ([]Product, error)
	GetByID(id uuid.UUID, tenantID uuid.UUID) (*Product, error)
	Update(product *Product) error
	Delete(id uuid.UUID, tenantID uuid.UUID) error
	GetBySKU(sku string, tenantID uuid.UUID) (*Product, error)
	GetByBarcode(barcode string, tenantID uuid.UUID) (*Product, error)
	GetCategoryByID(id uuid.UUID) (*Category, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetAll(tenantID uuid.UUID) ([]Product, error) {
	var products []Product
	err := r.db.Preload("Tenant").Preload("Category").Where("tenant_id = ?", tenantID).Find(&products).Error
	return products, err
}

func (r *productRepository) GetBySKU(sku string, tenantID uuid.UUID) (*Product, error) {
	var product Product
	err := r.db.Where("sku = ? AND tenant_id = ?", sku, tenantID).First(&product).Error
	return &product, err
}

func (r *productRepository) GetByBarcode(barcode string, tenantID uuid.UUID) (*Product, error) {
	var product Product
	err := r.db.Where("barcode = ? AND tenant_id = ?", barcode, tenantID).First(&product).Error
	return &product, err
}

func (r *productRepository) GetCategoryByID(id uuid.UUID) (*Category, error) {
	var category Category
	err := r.db.Where("id = ?", id).First(&category).Error
	return &category, err
}
func (r *productRepository) GetByID(id uuid.UUID, tenantID uuid.UUID) (*Product, error) {
	var product Product
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&product).Error
	return &product, err
}

func (r *productRepository) Update(product *Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uuid.UUID, tenantID uuid.UUID) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&Product{}).Error
}

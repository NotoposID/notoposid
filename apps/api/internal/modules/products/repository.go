package products

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TenantID      uuid.UUID `gorm:"type:uuid;not null;index"`
	CategoryID    *uuid.UUID `gorm:"type:uuid"`
	Name          string    `gorm:"size:255;not null"`
	SKU           string    `gorm:"size:100"`
	Barcode       string    `gorm:"size:100"`
	Description   string    `gorm:"type:text"`
	Price         float64   `gorm:"type:decimal(15,2);not null"`
	CostPrice     float64   `gorm:"type:decimal(15,2)"`
	StockQuantity int       `gorm:"default:0"`
	MinStockLevel int       `gorm:"default:10"`
	ImageURL      string    `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Repository interface {
	Create(product *Product) error
	GetByTenant(tenantID uuid.UUID) ([]Product, error)
	GetByID(id uuid.UUID, tenantID uuid.UUID) (*Product, error)
	Update(product *Product) error
	Delete(id uuid.UUID, tenantID uuid.UUID) error
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

func (r *productRepository) GetByTenant(tenantID uuid.UUID) ([]Product, error) {
	var products []Product
	err := r.db.Where("tenant_id = ?", tenantID).Find(&products).Error
	return products, err
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

package categories

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TenantID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Name        string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository interface {
	Create(category *Category) error
	GetAll(tenantID uuid.UUID) ([]Category, error)
	GetByID(id uuid.UUID, tenantID uuid.UUID) (*Category, error)
	Update(category *Category) error
	Delete(id uuid.UUID, tenantID uuid.UUID) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetAll(tenantID uuid.UUID) ([]Category, error) {
	var categories []Category
	err := r.db.Where("tenant_id = ?", tenantID).Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetByID(id uuid.UUID, tenantID uuid.UUID) (*Category, error) {
	var category Category
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&category).Error
	return &category, err
}

func (r *categoryRepository) Update(category *Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uuid.UUID, tenantID uuid.UUID) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&Category{}).Error
}

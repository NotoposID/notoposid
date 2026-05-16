package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TenantID     uuid.UUID `gorm:"type:uuid;not null;index"`
	Name         string    `gorm:"size:255;not null"`
	Email        string    `gorm:"size:255;not null"`
	PasswordHash string    `gorm:"type:text"`
	Role         string    `gorm:"size:50;default:'staff'"` // owner, admin, manager, staff
	AvatarURL    string    `gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Repository interface {
	Create(user *User) error
	GetAll(tenantID uuid.UUID) ([]User, error)
	GetByID(id uuid.UUID, tenantID uuid.UUID) (*User, error)
	GetByEmail(email string, tenantID uuid.UUID) (*User, error)
	Update(user *User) error
	Delete(id uuid.UUID, tenantID uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetAll(tenantID uuid.UUID) ([]User, error) {
	var users []User
	err := r.db.Where("tenant_id = ?", tenantID).Find(&users).Error
	return users, err
}

func (r *userRepository) GetByID(id uuid.UUID, tenantID uuid.UUID) (*User, error) {
	var user User
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&user).Error
	return &user, err
}

func (r *userRepository) GetByEmail(email string, tenantID uuid.UUID) (*User, error) {
	var user User
	err := r.db.Where("email = ? AND tenant_id = ?", email, tenantID).First(&user).Error
	return &user, err
}

func (r *userRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uuid.UUID, tenantID uuid.UUID) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&User{}).Error
}

package tenants

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name             string    `gorm:"size:255;not null"`
	Domain           string    `gorm:"size:255;unique"`
	LogoURL          string    `gorm:"type:text"`
	SubscriptionPlan string    `gorm:"size:50;default:'free'"`
	IsActive         bool      `gorm:"default:true"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

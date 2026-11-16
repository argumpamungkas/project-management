package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// struct User represents a user in the system.
type User struct {
	InternalID int64          `json:"internal_id" db:"internal_id" gorm:"primaryKey"`
	PublicID   uuid.UUID      `json:"public_id" db:"public_id"`
	Name       string         `json:"name" db:"name"`
	Email      string         `json:"email" db:"email" gorm:"unique"`
	Password   string         `json:"password,omitempty" db:"password" gorm:"column:password"`
	Role       string         `json:"role" db:"role"` // admin, user
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" db:"deleted_at" gorm:"index"` //soft delete, dibuat indexing
}

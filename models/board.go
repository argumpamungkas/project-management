package models

import (
	"time"

	"github.com/google/uuid"
)

type Board struct {
	InternalID    int64      `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID      uuid.UUID  `json:"public_id" db:"public_id"`
	Title         string     `json:"title" db:"title"`
	Description   string     `json:"description" db:"description"`
	OwnerID       int64      `json:"owner_internal_id" db:"owner_internal_id"` // user id
	OwnerPublicID uuid.UUID  `json:"owner_public_id" db:"owner_public_id"`     // user public id
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	DueDate       *time.Time `json:"due_date,omitempty" db:"due_date"` // use pointer to time.Time for nullable field, user address from other variable address
}

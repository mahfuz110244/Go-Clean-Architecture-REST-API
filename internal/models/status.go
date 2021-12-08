package models

import (
	"time"

	"github.com/google/uuid"
)

// Status base model
type Status struct {
	ID          uuid.UUID  `json:"id" db:"id" validate:"omitempty,uuid"`
	CreatedBy   uuid.UUID  `json:"created_by" db:"created_by" validate:"omitempty"`
	UpdatedBy   uuid.UUID  `json:"updated_by" db:"updated_by" validate:"omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at" validate:"omitempty"`
	Name        string     `json:"name" db:"name" validate:"required,lte=36"`
	Description string     `json:"description" db:"description" validate:"required,lte=255"`
	Active      bool       `json:"active" db:"active" validate:"omitempty"`
	OrderNumber int        `json:"order_number" db:"order_number" validate:"omitempty"`
}

// All Status response
type StatusList struct {
	TotalCount int       `json:"total_count"`
	TotalPages int       `json:"total_pages"`
	Page       int       `json:"page"`
	Size       int       `json:"size"`
	HasMore    bool      `json:"has_more"`
	Status     []*Status `json:"data"`
}

// Status base
type StatusBase struct {
	ID          uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	Name        string    `json:"name" db:"name" validate:"required,gte=10"`
	Description string    `json:"description" db:"description" validate:"required,gte=255"`
	Active      bool      `json:"active" db:"active" validate:"omitempty"`
	OrderNumber int       `json:"order_number" db:"order_number" validate:"omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by" validate:"omitempty,uuid"`
	UpdatedBy   uuid.UUID `json:"updated_by" db:"updated_by" validate:"omitempty,uuid"`
}

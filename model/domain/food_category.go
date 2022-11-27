package domain

import (
	"time"

	"github.com/google/uuid"
)

type FoodCategory struct {
	ID        uuid.UUID
	ParentID  uuid.UUID
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
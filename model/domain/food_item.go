package domain

import (
	"time"

	"github.com/google/uuid"
)

type FoodItem struct {
	ID         uuid.UUID
	Title      string
	Size       Food_Size
	Price      float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Food_Size uint8

const (
	Food_Size_Unset  Food_Size = iota
	Food_Size_Small
	Food_Size_Meduim
	Food_Size_Large 
)
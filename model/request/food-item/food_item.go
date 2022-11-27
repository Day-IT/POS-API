package fooditem

import (
	"github.com/Delaram-Gholampoor-Sagha/POS-API/model/domain"
	"github.com/google/uuid"
)

type Create struct {
	Title string    `json:"title"`
	Size  domain.Food_Size `json:"size"`
	Price float64   `json:"price"`
}

type Get struct {
	ID uuid.UUID `json:"id"`
}

type FindByID struct {
	FoodID uuid.UUID `json:"food_id"`
}

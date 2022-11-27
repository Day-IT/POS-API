package food

import (
	"context"

	"github.com/Delaram-Gholampoor-Sagha/POS-API/model/domain"
	"github.com/google/uuid"
)

type foodItemstorage struct {
}

func NewfoodItemstorage() {
	return
}

func (s *foodItemstorage) Save(ctx context.Context, fooditem domain.FoodItem) error {
	return nil
}

func (s *foodItemstorage) FindByID(ctx context.Context, id uuid.UUID) (*domain.FoodItem, error) {
	return &domain.FoodItem{} , nil
}

func (s *foodItemstorage) Get(ctx context.Context, id uuid.UUID) (*domain.FoodItem, error) {
	return &domain.FoodItem{} , nil
}

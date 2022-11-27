package fooditem

import (
	"context"

	"github.com/Delaram-Gholampoor-Sagha/POS-API/model/domain"
	"github.com/google/uuid"
)

type Food_Item_Repo interface {
	Save(ctx context.Context , fooditem domain.FoodItem) error
	FindByID(ctx context.Context ,id uuid.UUID) (*domain.FoodItem, error)
	Get(ctx context.Context ,id uuid.UUID) (*domain.FoodItem , error )
}

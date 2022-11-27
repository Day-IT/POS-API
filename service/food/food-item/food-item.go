package fooditem

import (
	"context"

	req "github.com/Delaram-Gholampoor-Sagha/POS-API/model/request/food-item"
	res "github.com/Delaram-Gholampoor-Sagha/POS-API/model/response/food-item"
	fooditem "github.com/Delaram-Gholampoor-Sagha/POS-API/repository/food-item"
)

type foodItemService struct {
	repo fooditem.Food_Item_Repo
}

func NewfoodItemService(repo fooditem.Food_Item_Repo) Service {
	return &foodItemService{
		repo: repo,
	}
}

func (s *foodItemService) Create(ctx context.Context, req req.Create) (*res.Create, error) {
	return &res.Create{}, nil
}

func (s *foodItemService) FindByID(ctx context.Context, req req.FindByID) (*res.FindByID, error) {
	return &res.FindByID{}, nil
}

func (s *foodItemService) Get(ctx context.Context, req req.Get) (*res.Get, error) {
	return &res.Get{}, nil
}

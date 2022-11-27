package fooditem

import (
	"context"

	req "github.com/Delaram-Gholampoor-Sagha/POS-API/model/request/food-item"
	res "github.com/Delaram-Gholampoor-Sagha/POS-API/model/response/food-item"
)

type Service interface {
	Create(ctx context.Context, req req.Create) (*res.Create, error)
	FindByID(ctx context.Context, req req.FindByID) (*res.FindByID , error)
	Get(ctx context.Context , req req.Get) (*res.Get , error)
}

// type foodService struct {
// 	repo food.Food_Item_Repo
// }

// // NewUserService func
// func NewfoodService(repo food.Food_Item_Repo) Service {
// 	return &foodService{
// 		repo: repo,

// 	}
// }

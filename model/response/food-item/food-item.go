package fooditem

import "github.com/Delaram-Gholampoor-Sagha/POS-API/model/domain"

type Create struct {
	Food *domain.FoodItem `json:"food_item"`
}


type FindByID struct {
	Food *domain.FoodItem `json:"food_item"`
}

type Get struct {
	Food *domain.FoodItem `json:"food_item"`
}
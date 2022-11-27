package repository

import fooditem "github.com/Delaram-Gholampoor-Sagha/POS-API/repository/food-item"

type Repository struct {
	FoodItem fooditem.Food_Item_Repo
}
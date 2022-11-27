package http

import (
	"net/http"

	fooditem "github.com/Delaram-Gholampoor-Sagha/POS-API/service/food/food-item"
)

type foodItemHandler struct {
	s fooditem.Service
}

func NewfoodItemHandler(serivce fooditem.Service) *foodItemHandler {
	return &foodItemHandler{
		s: serivce,
	}
}

func (h *foodItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

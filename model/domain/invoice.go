package domain

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID            uuid.UUID
	Number        int
	Type          Order_Type
	Item_Quantity int
	Total_Amount  int
	Created_At    time.Time
	Updated_At    time.Time
}

type Order_Type uint8

const (
	Order_Type_Unset Order_Type = iota
	Order_Type_TakeAway
	Order_Type_Collection
	Order_Type_DineIn
	Order_Type_Reservation
)

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Food_Order struct {
	ID              uuid.UUID
	Cook_Inst       Cook_Inst_Level
	Special_Request string
	No              string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}


type Cook_Inst_Level uint8

const (
	Cook_Inst_Level_Unset Cook_Inst_Level = iota
	Cook_Inst_Level_WellDone 
	Cook_Inst_Level_Vinegar
	Cook_Inst_Level_Salt
	Cook_Inst_Level_Medium
	Cook_Inst_Level_Rare
)

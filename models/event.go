package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

//Event document !
type Event struct {
	uadmin.Model
	Name              string
	Counterparty      Counterparty
	CounterpartyID    uint
	Contract          Contract
	ContractID        uint
	Places            Places
	PlacesID          uint
	TypesOfDiscount   TypesOfDiscount
	TypesOfDiscountID uint
	Discount          int
	// File              string `uadmin:"file"`
	// Description       string `uadmin:"html"`
	// RawText           string `uadmin:"list_exclude"`
	Folder            Folder `uadmin:"filter"`
	FolderID          uint
	CreatedDate       time.Time
	Channel           Channel `uadmin:"list_exclude"`
	ChannelID         uint
	CreatedBy         string
}

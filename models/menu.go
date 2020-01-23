package models

import (
	"github.com/uadmin/uadmin"
)

//Menu for event document !
type Menu struct {
	uadmin.Model
	Name           string
	Event          Event
	EventID        uint
	Nomenclature   Nomenclature
	nomenclatureID uint
	Number         int `uadmin:"min:1"`
	Cost           int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field.`
	Sum            int
	Kit            Kit
	KitID          uint
}

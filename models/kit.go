package models

import (
	"github.com/uadmin/uadmin"
)

//Kit for event document !
type Kit struct {
	uadmin.Model
	Name           string
	Nomenclature   Nomenclature
	nomenclatureID uint
	NumberInKit    int `uadmin:"min:1"`
	CostInKit      int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field.`
}

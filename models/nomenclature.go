package models

import (
	"github.com/uadmin/uadmin"
)

//Nomenclature for event document !
type Nomenclature struct {
	uadmin.Model
	Name                 string
	CostInNomenclature   int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field.`
	Service              bool
	Measure              Measure
	MeasureID            uint
	FrequencyOfService   FrequencyOfService
	FrequencyOfServiceID uint
	tax                  int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field.`
}

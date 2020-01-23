package models

import (
	"github.com/uadmin/uadmin"
)

//Contract for event document !
type Contract struct {
	uadmin.Model
	Name string
	// ContractCategories   ContractCategories
	// ContractCategoriesID uint
	// Event   Event
	// EventID uint
}

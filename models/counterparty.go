package models

import (
	"github.com/uadmin/uadmin"
)

//Counterparty for event document !
type Counterparty struct {
	uadmin.Model
	Name    string
	Contract   Contract
	ContractID uint
}

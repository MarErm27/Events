package models

import (
	"github.com/uadmin/uadmin"
)

// TypesOfDiscount !
type TypesOfDiscount struct {
	uadmin.Model
	Name    string `uadmin:"required"`
}

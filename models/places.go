package models

import (
	"github.com/uadmin/uadmin"
)

// Places
type Places struct {
	uadmin.Model
	Name string `uadmin:"required"`
}

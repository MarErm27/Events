package models

import (
	"github.com/uadmin/uadmin"
)

//FrequencyOfService for event document !
type FrequencyOfService struct {
	uadmin.Model
	Name string
}

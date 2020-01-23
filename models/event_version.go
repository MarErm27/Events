package models

import (
	"fmt"
	"time"

	"github.com/uadmin/uadmin"
)

// EventDocumentVersion !
type EventVersion struct {
	uadmin.Model
	Event   Event
	EventID uint
	File    string `uadmin:"file"`
	Number  int    `uadmin:"help:version number"`
	Date    time.Time
}

// Returns the version number
func (e EventVersion) String() string {
	return fmt.Sprint(e.Number)
}

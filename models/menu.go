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
	NomenclatureID uint
	Number         int `uadmin:"min:1"`
	Cost           int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field."`
	Sum            int `uadmin:"read_only"`
	Kit            Kit
	KitID          uint
	Description    string `uadmin:"html"`
}

// Save !
func (m *Menu) Save() {
	// Prints the result
	uadmin.Trail(uadmin.DEBUG, "The document has changed.")
	m.Sum = m.Number * m.Cost
	// Save the document
	uadmin.Save(m)
	// Sets the document value to the DocumentVersion

}

// func (m Menu) Validate() (ret map[string]string) {
// 	// ret = map[string]string{}
// 	// if v.Name != "test" {
// 	//     ret["Name"] = "Error name not found"
// 	// }
// 	m.Sum = m.Number * m.Cost
//     errMsg = map[string]string{}
//     // Get any records from the database that maches the name of
//     // this record and make sure the record is not the record we are
//     // editing right now
//     menu := Menu{}
//     if uadmin.Count(&menu, "name = ? AND id <> ?", "Summary", m.ID) != 0 {
//         errMsg["Name"] = "This todo name is already in the system"
//     }else{
// 		summary := Menu{}
// 		summary.Name = "Summary"
// 		summary.Event = m.Event
// 		summary.Number = m.Number
// 		summary.Cost = m.Cost
// 		summary.Sum = m.Sum

// 		// Save the document version
// 		uadmin.Save(&summary)
// 	}
//     return
// }

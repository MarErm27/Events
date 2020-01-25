package main

import (
	"github.com/MarErm27/events/models"
	"github.com/uadmin/uadmin"
)

func main() {
	// Sets the name of the website that shows on title and dashboard
	uadmin.SiteName = "Events and excursions"
	uadmin.Register(
		models.Folder{},
		models.FolderGroup{},        // place it here
		models.FolderUser{},         // place it here
		models.Channel{},            // place it here
		models.Event{},              // place it here
		models.EventVersion{},       // place it here
		models.Menu{},               // place it here
		models.Kit{},                // place it here
		models.EventVersion{},       // place it here
		models.FrequencyOfService{}, // place it here
		models.Measure{},            // place it here
		models.EventVersion{},       // place it here
		models.Counterparty{},
		models.Contract{},
		models.ContractCategories{},
		models.TypesOfDiscount{},
		models.Nomenclature{},
		models.Places{},
	)
	uadmin.RegisterInlines(
		models.Event{},
		map[string]string{
			"menu":               "EventID",
			"counterparty":       "EventID",
			"contract":           "EventID",
			"typesOfDiscount":    "EventID",
			"frequencyOfService": "EventID",
		},
	)
	uadmin.RegisterInlines(
		models.Kit{},
		map[string]string{
			"nomenclature": "KitID",
		},
	)
	// uadmin.RegisterInlines(
	// 	models.ContractCategories{},
	// 	map[string]string{
	// 		"contract": "contractCategoriesID",
	// 	},
	// )
	// uadmin.RegisterInlines(
	// 	models.Menu{},
	// 	map[string]string{
	// 		"kit":          "MenuID",
	// 		"nomenclature": "MenuID",
	// 	},
	// )

	// uadmin.RegisterInlines(
	// 	models.Document{},
	// 	map[string]string{
	// 		"kit":                "NomenclatureID",
	// 		"measure":            "NomenclatureID",
	// 		"frequencyOfService": "NomenclatureID",
	// 	},
	// )

	// Call the ModelName schema of "menu" model
	menu := uadmin.Schema["menu"]

	// Include Javascript file for the form
	menu.IncludeFormJS = []string{"/static/js/menuForm.js"}

	// Pass back the menu variable to apply changes
	uadmin.Schema["menu"] = menu

	// Activates a uAdmin server
	uadmin.StartServer()
}

package main

import (
	"net/http"

	"github.com/MarErm27/Events/api"
	"github.com/MarErm27/Events/models"
	"github.com/uadmin/uadmin"
)

func redirectToAdminPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}

func main() {
	// Sets the name of the website that shows on title and dashboard
	//uadmin.Port = 4500
	//uadmin.RootURL = "/admin/"
	//uadmin.SiteName = "Events and excursions"

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
	// Initialize the Setting
	//setting := uadmin.Setting{}

	// Use update function to apply the assigned value in Settings
	//uadmin.Update(&setting, "Value", strconv.Itoa(uadmin.Port), "code = ?", "uAdmin.Port")
	//uadmin.Update(&setting, "Value", uadmin.RootURL, "code = ?", "uAdmin.RootURL")
	//uadmin.Update(&setting, "Value", uadmin.SiteName, "code = ?", "uAdmin.SiteName")
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

	// Call the ModelName schema of "menu" model
	event := uadmin.Schema["event"]

	// Include Javascript file for the form
	event.IncludeFormJS = []string{"/static/js/eventForm.js"}

	// Pass back the menu variable to apply changes
	uadmin.Schema["event"] = event

	http.HandleFunc("/api/", api.APIHandler)
	//http.HandleFunc("/table/", views.HTTPHandler)
	// Activates a uAdmin server
	http.HandleFunc("/", http.HandlerFunc(redirectToAdminPage))
	uadmin.StartServer()
}

package api

import (
	"fmt"
	"net/http"
	"strings"

	// Specify the username that you used inside github.com folder
	"github.com/MarErm27/Events/models"
	//"github.com/MarErm27/Events/models"
	"github.com/uadmin/uadmin"
)

// TableHandler !
func TableHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /todo_list
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/table")

	switch r.Method {
	case "GET":
		//http.ServeFile(w, r, "form.html")
		host := string(r.Host)
		referer := r.Header["Referer"][0]
		hostIndex := strings.Index(referer, host)
		modelNameIndex := hostIndex + len(host) + len(uadmin.RootURL)
		refererSubStringAfterModelNameIndex := referer[modelNameIndex:]
		modelNameLastIndex := strings.Index(refererSubStringAfterModelNameIndex, "/")
		modelName := refererSubStringAfterModelNameIndex[:modelNameLastIndex]
		switch modelName {
		case "event":
			//menu := []models.["modelName"]{}
			// Assigns a map as a string of interface to store any types of values
			results := []map[string]interface{}{}
			// Fetches all object in the database
			menu := []models.Menu{}
			eventIDToGet := refererSubStringAfterModelNameIndex[modelNameLastIndex+1:]
			uadmin.Filter(&menu, "event_id = ?", eventIDToGet)

			// Accesses and fetches data from another model
			for i := range menu {
				uadmin.Preload(&menu[i])
				// Assigns the string of interface in each Todo fields
				results = append(results, map[string]interface{}{
					// This returns only the name of the Category model, not the
					// other fields
					"Nomenclature": menu[i].Nomenclature.Name,
					// This returns only the name of the Friend model, not the
					// other fields
					"Number": menu[i].Number,
					// This returns only the name of the Item model, not the other
					// fields
					"Cost":        menu[i].Cost,
					"Sum":         menu[i].Sum,
					"Kit":         menu[i].Kit.Name,
					"Description": menu[i].Description,
				})
			}

			// Prints the todo in JSON format
			uadmin.ReturnJSON(w, r, results)
		default:
			fmt.Fprintf(w, "Sorry, there is no table for this model specified.")
		}
	case "POST":
		//Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)

		//case method is get nomenclature list - get nomenclature for dropdown menu in a table
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

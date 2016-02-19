package render

import (
	"../system/assets"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templates []string, name string, data interface{}, app string) error {

	// Load Template Routines
	finalTemplate := template.New(name)

	//get parse template files and load it from the Assets
	for _, file := range templates {
		asset, err := assets.Asset(file)
		if err != nil {
			panic(err)
		}
		finalTemplate.Parse(string(asset))
	}

	// Execute the template
	err := finalTemplate.ExecuteTemplate(w, name, data)
	if err != nil {
		return err
	}

	return nil
}

func GetBaseTemplates() []string {
	templates := []string{"views/base/footer.html", "views/base/header.html", "views/base.html"}
	return templates
}

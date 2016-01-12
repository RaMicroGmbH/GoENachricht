package controllers

import (
	"GoENachricht/render"
	"encoding/json"
	"log"
	"net/http"
)

// Default struct for title
type Content struct {
	Title        string
	Version      string
	ProtPort     string
	CurrentJobNr string
}

type Result struct {
	Code    string
	Message string
}

type UserDefaultData struct {
	Code    string
	IniData InitData
}

var (
	StartUpContent Content
	searchquery    string
)

func InitContent(title, version, protPort string) {
	StartUpContent.Title = title
	StartUpContent.Version = version
	StartUpContent.ProtPort = protPort
}

/*-----------------------------------------------------------------------------
				Passing user default data to Frontend
 -----------------------------------------------------------------------------*/
func GetUserDefaultData(w http.ResponseWriter, req *http.Request) {

	result := UserDefaultData{"OK", InitData{UserSettings, AllSubjects, AllUsers}}
	js, err := json.Marshal(result)	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}

/*-----------------------------------------------------------------------------
	Rendering index.html page
------------------------------------------------------------------------------*/
func ConfigPage(w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates()
	templates = append(templates, "views/index.html")
	err := render.RenderTemplate(w, templates, "base", &StartUpContent, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*-----------------------------------------------------------------------------
	Rendering auswahlliste.html page
------------------------------------------------------------------------------*/

func AuswahllistePage(w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates()
	templates = append(templates, "views/auswahlliste.html")
	err := render.RenderTemplate(w, templates, "base", &StartUpContent, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*-----------------------------------------------------------------------------
	Updating ENachricht Data
-------------------------------------------------------------------------------*/
func UpdateENachrichtData(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	var enachricht = req.FormValue("enachricht")

	log.Println("ENACHRICHT:", enachricht)

	result := Result{"OK", "ENachricht Successfully received !"}
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

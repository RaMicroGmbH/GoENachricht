package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "os"
	"path/filepath"
	"strings"

	"../radata"
	"../render"
	"../utils"
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
	IniData radata.InitData
}

var (
	StartUpContent         Content
	searchquery            string
	relativefilepath       string
	relativeAttachmentpath string
)

func InitContent(title, version string) {
	StartUpContent.Title = title
	StartUpContent.Version = version
}

/*-----------------------------------------------------------------------------
				Passing user default data to Frontend
 ------------------------------------------------------------------------------*/
func GetUserDefaultData(w http.ResponseWriter, req *http.Request) {

	result := UserDefaultData{"OK", radata.InitData{radata.UserSettings, radata.AllSubjects, radata.AllUsers}}
	js, err := json.Marshal(result)
	Checkerror(err)
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
	Checkerror(err)
}

/*-----------------------------------------------------------------------------
	Rendering auswahlliste.html page
------------------------------------------------------------------------------*/

func AuswahllistePage(w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates()
	templates = append(templates, "views/auswahlliste.html")
	err := render.RenderTemplate(w, templates, "base", &StartUpContent, "")
	Checkerror(err)
}

/*-----------------------------------------------------------------------------
	Updating E-Nachricht Data (along with Attachement if any !)
-------------------------------------------------------------------------------*/

func UpdateENachrichtData(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	var b64data = req.FormValue("B64Data")
	//var filename = req.FormValue("filename")
	var attachmentname = req.FormValue("attachmentname")
	var enachricht = req.FormValue("enachricht")

	if b64data != "undefined" && b64data != "" {

		dataFirstPart := strings.Split(b64data, "base64,")
		encdata := strings.Split(dataFirstPart[1], "\">")

		data, err := base64.StdEncoding.DecodeString(encdata[0])

		var fapath = radata.GetENachrichtPath()

		relativeAttachmentpath = filepath.Join(fapath, attachmentname)
		ioutil.WriteFile(relativeAttachmentpath, []byte(data), 0x777)

		if err != nil {
			fmt.Println("error:", err)
			return
		}
	}

	var fname = utils.GetDate() + ".txt"
	var fpath = radata.GetENachrichtPath()
	relativefilepath = filepath.Join(fpath, fname)

	ioutil.WriteFile(relativefilepath, []byte(enachricht), 0x777)

	result := Result{"OK", "E-Nachricht Successfully received !"}
	js, err := json.Marshal(result)
	Checkerror(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

/* Check error */
func Checkerror(err error) {
	if err != nil {
		var w http.ResponseWriter
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package controllers

import (
	"../render"
	_ "encoding/json"
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

var (
	StartUpContent Content
	searchquery    string
)

func InitContent(title, version, protPort string) {
	StartUpContent.Title = title
	StartUpContent.Version = version
	StartUpContent.ProtPort = protPort
}

func ConfigPage(w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates()
	templates = append(templates, "views/index.html")
	err := render.RenderTemplate(w, templates, "base", &StartUpContent, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

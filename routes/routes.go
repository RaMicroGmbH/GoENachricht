package routes

import (
	"GoENachricht/controllers"
	"net/http"
)

func Include(m *http.ServeMux) {

	m.HandleFunc("/getUserDefaultData", controllers.GetUserDefaultData)
	m.HandleFunc("/enachricht/ui", controllers.ConfigPage)

	m.HandleFunc("/enachricht/auswahlliste", controllers.AuswahllistePage)
	m.HandleFunc("/updateENachrichtData", controllers.UpdateENachrichtData)
}

package routes

import (
	"../controllers"
	"net/http"
)

func Include(m *http.ServeMux) {

	m.HandleFunc("/enachricht/ui", controllers.ConfigPage)

}

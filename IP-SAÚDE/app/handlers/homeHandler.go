package handlers

import (
	"html/template"
	"net/http"

	"servidorHTTP/app/utils"
)

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	dashboardData, err := utils.GetDashboardData()
	if err != nil {
		http.Error(response, "Erro ao carregar os dados da UTI", http.StatusInternalServerError)
		return
	}

	templateFile, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(response, "Erro ao carregar a página inicial", http.StatusInternalServerError)
		return
	}

	err = templateFile.Execute(response, dashboardData)
	if err != nil {
		http.Error(response, "Erro ao renderizar a página inicial", http.StatusInternalServerError)
		return
	}
}

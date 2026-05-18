package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"servidorHTTP/app/utils"
)

func UpdatePatientHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}
	idTexto := strings.TrimSpace(request.FormValue("id"))
	prioridade := strings.TrimSpace(request.FormValue("prioridade"))

	if idTexto == "" || prioridade == "" {
		http.Error(response, "ID e prioridade são obrigatórios", http.StatusBadRequest)
		return
	}

	idPaciente, err := strconv.Atoi(idTexto)
	if err != nil || idPaciente <= 0 {
		http.Error(response, "ID inválido", http.StatusBadRequest)
		return
	}

	err = utils.UpdatePatientPriority(idPaciente, prioridade)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}

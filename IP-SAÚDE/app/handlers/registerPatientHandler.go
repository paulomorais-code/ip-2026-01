package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"servidorHTTP/app/utils"
)

func RegisterPatientHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	nomeCompleto := strings.TrimSpace(request.FormValue("nome"))
	idadeTexto := strings.TrimSpace(request.FormValue("idade"))
	leito := strings.TrimSpace(request.FormValue("leito"))
	diagnosticoInicial := strings.TrimSpace(request.FormValue("diagnostico"))
	prioridade := strings.TrimSpace(request.FormValue("prioridade"))
	origem := strings.TrimSpace(request.FormValue("origem"))
	observacoesClinicas := strings.TrimSpace(request.FormValue("observacoes"))

	if nomeCompleto == "" || idadeTexto == "" || leito == "" || diagnosticoInicial == "" || prioridade == "" || origem == "" {
		http.Error(response, "Preencha os campos obrigatórios", http.StatusBadRequest)
		return
	}

	idade, err := strconv.Atoi(idadeTexto)
	if err != nil || idade <= 0 {
		http.Error(response, "Idade inválida", http.StatusBadRequest)
		return
	}

	err = utils.InsertPatient(nomeCompleto, idade, leito, diagnosticoInicial, prioridade, origem, observacoesClinicas)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}

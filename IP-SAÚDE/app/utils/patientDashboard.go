package utils

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Patient struct {
	ID                  int
	NomeCompleto        string
	Idade               int
	Leito               string
	DiagnosticoInicial  string
	Prioridade          string
	Origem              string
	ObservacoesClinicas string
	Status              string
	DataCadastro        string
	DataAtualizacao     string
	StatusClass         string
	PriorityClass       string
}

type DashboardData struct {
	Patients          []Patient
	TotalCount        int
	ActiveCount       int
	DischargedCount   int
	TransferCount     int
	DeathCount        int
	CriticalCount     int
	HighRiskCount     int
	StableCount       int
	LastUpdated       string
	EmptyStateMessage string
}

func GetDashboardData() (*DashboardData, error) {
	patientRows, err := DB.Query(`
		SELECT
			id,
			nome_completo,
			idade,
			leito,
			diagnostico_inicial,
			prioridade,
			origem,
			COALESCE(observacoes_clinicas, ''),
			COALESCE(status, 'Ativo'),
			COALESCE(TO_CHAR(data_cadastro, 'DD/MM/YYYY HH24:MI'), '-'),
			COALESCE(TO_CHAR(data_atualizacao, 'DD/MM/YYYY HH24:MI'), '-')
		FROM pacientes_uti
		WHERE COALESCE(status, 'Ativo') = 'Ativo'
		ORDER BY data_cadastro DESC, id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer patientRows.Close()

	patients := make([]Patient, 0)
	for patientRows.Next() {
		var patient Patient
		if err := patientRows.Scan(
			&patient.ID,
			&patient.NomeCompleto,
			&patient.Idade,
			&patient.Leito,
			&patient.DiagnosticoInicial,
			&patient.Prioridade,
			&patient.Origem,
			&patient.ObservacoesClinicas,
			&patient.Status,
			&patient.DataCadastro,
			&patient.DataAtualizacao,
		); err != nil {
			return nil, err
		}
		patient.StatusClass = statusClass(patient.Status)
		patient.PriorityClass = priorityClass(patient.Prioridade)
		patients = append(patients, patient)
	}
	if err := patientRows.Err(); err != nil {
		return nil, err
	}

	var totalCount int
	var activeCount int
	var dischargedCount int
	var transferCount int
	var deathCount int
	var criticalCount int
	var highRiskCount int
	var stableCount int
	var lastUpdated sql.NullString

	err = DB.QueryRow(`
		SELECT
			COUNT(*) FILTER (WHERE COALESCE(status, 'Ativo') = 'Ativo') AS total_count,
			COUNT(*) FILTER (WHERE COALESCE(status, 'Ativo') = 'Ativo') AS active_count,
			0 AS discharged_count,
			0 AS transfer_count,
			0 AS death_count,
			COUNT(*) FILTER (WHERE COALESCE(status, 'Ativo') = 'Ativo' AND prioridade = 'Crítico') AS critical_count,
			COUNT(*) FILTER (WHERE COALESCE(status, 'Ativo') = 'Ativo' AND prioridade = 'Alto risco') AS highrisk_count,
			COUNT(*) FILTER (WHERE COALESCE(status, 'Ativo') = 'Ativo' AND prioridade = 'Estável') AS stable_count,
			TO_CHAR(MAX(data_atualizacao), 'DD/MM/YYYY HH24:MI') AS last_updated
		FROM pacientes_uti
	`).Scan(&totalCount, &activeCount, &dischargedCount, &transferCount, &deathCount, &criticalCount, &highRiskCount, &stableCount, &lastUpdated)
	if err != nil {
		return nil, err
	}

	formattedLastUpdated := "-"
	if lastUpdated.Valid && strings.TrimSpace(lastUpdated.String) != "" {
		formattedLastUpdated = lastUpdated.String
	}

	return &DashboardData{
		Patients:          patients,
		TotalCount:        totalCount,
		ActiveCount:       activeCount,
		DischargedCount:   dischargedCount,
		TransferCount:     transferCount,
		DeathCount:        deathCount,
		CriticalCount:     criticalCount,
		HighRiskCount:     highRiskCount,
		StableCount:       stableCount,
		LastUpdated:       formattedLastUpdated,
		EmptyStateMessage: "Nenhum paciente ativo cadastrado no momento.",
	}, nil
}

func InsertPatient(nomeCompleto string, idade int, leito string, diagnosticoInicial string, prioridade string, origem string, observacoesClinicas string) error {
	if strings.TrimSpace(nomeCompleto) == "" || strings.TrimSpace(leito) == "" {
		return fmt.Errorf("nome e leito são obrigatórios")
	}

	var existingPatientID int
	err := DB.QueryRow(`
		SELECT id
		FROM pacientes_uti
		WHERE leito = $1 AND COALESCE(status, 'Ativo') = 'Ativo'
		LIMIT 1
	`, leito).Scan(&existingPatientID)
	if err == nil {
		return fmt.Errorf("já existe um paciente ativo no leito %s", leito)
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	query := `INSERT INTO pacientes_uti (nome_completo, idade, leito, diagnostico_inicial, prioridade, origem, observacoes_clinicas, status, data_cadastro, data_atualizacao) VALUES ($1, $2, $3, $4, $5, $6, $7, 'Ativo', NOW(), NOW())`
	_, err = DB.Exec(query, nomeCompleto, idade, leito, diagnosticoInicial, prioridade, origem, observacoesClinicas)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePatient(idPaciente int, nomeCompleto string, idade int, leito string, diagnosticoInicial string, prioridade string, origem string, observacoesClinicas string) error {
	if idPaciente <= 0 {
		return fmt.Errorf("id inválido")
	}

	query := `UPDATE pacientes_uti SET nome_completo = $1, idade = $2, leito = $3, diagnostico_inicial = $4, prioridade = $5, origem = $6, observacoes_clinicas = $7, data_atualizacao = NOW() WHERE id = $8 AND COALESCE(status, 'Ativo') = 'Ativo'`
	result, err := DB.Exec(query, nomeCompleto, idade, leito, diagnosticoInicial, prioridade, origem, observacoesClinicas, idPaciente)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("paciente ativo não encontrado para o id %d", idPaciente)
	}

	return nil
}

func UpdatePatientStatus(idPaciente int, status string) error {
	if idPaciente <= 0 {
		return fmt.Errorf("id inválido")
	}
	if strings.TrimSpace(status) == "" {
		return fmt.Errorf("status inválido")
	}

	query := `UPDATE pacientes_uti SET status = $1, data_atualizacao = NOW() WHERE id = $2 AND COALESCE(status, 'Ativo') = 'Ativo'`
	result, err := DB.Exec(query, status, idPaciente)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("paciente ativo não encontrado para o id %d", idPaciente)
	}

	return nil
}

func UpdatePatientPriority(idPaciente int, priority string) error {
	if idPaciente <= 0 {
		return fmt.Errorf("id inválido")
	}
	if strings.TrimSpace(priority) == "" {
		return fmt.Errorf("prioridade inválida")
	}

	query := `UPDATE pacientes_uti SET prioridade = $1, data_atualizacao = NOW() WHERE id = $2 AND COALESCE(status, 'Ativo') = 'Ativo'`
	result, err := DB.Exec(query, priority, idPaciente)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("paciente ativo não encontrado para o id %d", idPaciente)
	}

	return nil
}

func statusClass(status string) string {
	switch status {
	case "Alta":
		return "success"
	case "Transferido":
		return "warning"
	case "Óbito":
		return "critical"
	default:
		return "attention"
	}
}

func priorityClass(priority string) string {
	switch priority {
	case "Crítico":
		return "critical"
	case "Alto risco":
		return "warning"
	case "Estável":
		return "success"
	default:
		return "attention"
	}
}

func formatTime(value time.Time) string {
	if value.IsZero() {
		return "-"
	}
	return value.Format("02/01/2006 15:04")
}

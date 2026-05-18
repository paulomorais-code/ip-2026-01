package utils

import (
	"fmt"
	"log"
)

func RemovePatient(idPaciente int) error {
	if idPaciente <= 0 {
		return fmt.Errorf("id inválido")
	}

	result, err := DB.Exec(`DELETE FROM pacientes_uti WHERE id = $1 AND COALESCE(status, 'Ativo') = 'Ativo'`, idPaciente)
	if err != nil {
		log.Printf("Erro ao remover paciente do banco de dados: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("paciente ativo não encontrado para o id %d", idPaciente)
	}

	log.Println("Paciente removido com sucesso!")
	return nil
}

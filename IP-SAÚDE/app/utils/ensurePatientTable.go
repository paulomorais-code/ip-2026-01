package utils

func EnsurePatientTable() error {
	query := `
CREATE TABLE IF NOT EXISTS pacientes_uti (
	id BIGSERIAL PRIMARY KEY,
	nome_completo VARCHAR(150) NOT NULL,
	idade INTEGER NOT NULL CHECK (idade > 0 AND idade < 130),
	leito VARCHAR(20) NOT NULL,
	diagnostico_inicial TEXT NOT NULL,
	prioridade VARCHAR(20) NOT NULL CHECK (prioridade IN ('Crítico', 'Alto risco', 'Estável')),
	origem VARCHAR(40) NOT NULL CHECK (origem IN ('Pronto-socorro', 'Centro cirúrgico', 'Transferência externa')),
	observacoes_clinicas TEXT,
	status VARCHAR(20) NOT NULL DEFAULT 'Ativo' CHECK (status IN ('Ativo', 'Alta', 'Transferido', 'Óbito')),
	data_cadastro TIMESTAMP NOT NULL DEFAULT NOW(),
	data_atualizacao TIMESTAMP NOT NULL DEFAULT NOW()
)
`
	_, err := DB.Exec(query)
	return err
}

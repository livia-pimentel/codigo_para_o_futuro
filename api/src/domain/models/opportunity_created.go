package models

import "github.com/google/uuid"

// OpportunityCreated representa a estrutura de dados retornada como resposta após o cadastro bem-sucedido de uma nova oportunidade, contendo o identificador único gerado.
type OpportunityCreated struct {
	ID uuid.UUID `json:"id"`
}
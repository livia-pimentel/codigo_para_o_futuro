package models

// OpportunityUpdate representa a estrutura de dados necessária para receber e validar a requisição de atualização dos dados de uma oportunidade existente.
type OpportunityUpdate struct {
	Title 		string `json:"title" validate:"required,min=2,max=120"`
	Company 	string `json:"company" validate:"required"`
	Category 	string `json:"category" validate:"required"`
}
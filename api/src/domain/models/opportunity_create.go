package models

// OpportunityCreate representa a estrutura de dados necessária para receber e validar a requisição de criação de uma nova oportunidade no sistema.
type OpportunityCreate struct {
	Title 		string `json:"title" validate:"required,min=2,max=120"`
	Company 	string `json:"company" validate:"required"`
	Category 	string `json:"category" validate:"required"`
}
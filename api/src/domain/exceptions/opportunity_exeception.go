package exceptions

const (
	// erros de regras de negocio
	ErrOpportunityNotFound      string = "errOpportunityNotFound"
	ErrOpportunityAlreadyExists string = "errOpportunityAlreadyExists"

	// erros de infra
	ErrOnFindAllOpportunity       string = "errOnFindAllOpportunity"
	ErrOnFindOpportunityByTitle   string = "errOnFindOpportunityByTitle"
	ErrOnExistsOpportunityById    string = "errOnExistsOpportunityById"
	ErrOnExistsOpportunityByTitle string = "errOnExistsOpportunityByTitle"
	ErrOnInsertOpportunity        string = "errOnInsertOpportunity"
	ErrOnUpdateOpportunity        string = "errOnUpdateOpportunity"
	ErrOnToogleActiveOpportunity  string = "errOnToogleActiveOpportunity"
	ErrOnDeleteOpportunity        string = "errOnDeleteOpportunity"
	ErrOnFindOpportunityById   string = "errOnFindOpportunityById"
)
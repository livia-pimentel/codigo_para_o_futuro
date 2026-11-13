package models

// OpportunityListParams representa os parâmetros de consulta (query parameters) que podem ser enviados na URL para filtrar e refinar as listagens de oportunidades.
type OpportunityListParams struct {
	Title string `form:"title"`
}
package controllers

import (
	"net/http"

	"github.com/colibriproject-dev/colibri-sdk-go/pkg/web/restserver"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/domain/exceptions"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/domain/models"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/domain/usecases"
	"github.com/google/uuid"
)

// OpportunityV1Controller gerencia os endpoints HTTP da versão 1 para o recurso de oportunidades,atuando como ponte entre as requisições externas e os casos de uso.
type OpportunityV1Controller struct {
	GetAllOpportunityUsecase       usecases.IGetAllOpportunityUsecase
	CreateOpportunityUsecase       usecases.ICreateOpportunityUsecase
	UpdateOpportunityUsecase       usecases.IUpdateOpportunityUsecase
	ToogleActiveOpportunityUsecase usecases.IToogleActiveOpportunityUsecase
	DeleteOpportunityUsecase       usecases.IDeleteOpportunityUsecase
	GetOpportunityByIdUsecase       usecases.IGetOpportunityByIdUsecase
}

// NewOpportunityV1Controller instancia o controller injetando todos os casos de uso necessários.
func NewOpportunityV1Controller() *OpportunityV1Controller {
	return &OpportunityV1Controller{
		GetAllOpportunityUsecase:       usecases.NewGetAllOpportunityUsecase(),
		CreateOpportunityUsecase:       usecases.NewCreateOpportunityUsecase(),
		UpdateOpportunityUsecase:       usecases.NewUpdateOpportunityUsecase(),
		ToogleActiveOpportunityUsecase: usecases.NewToogleActiveOpportunityUsecase(),
		DeleteOpportunityUsecase:       usecases.NewDeleteOpportunityUsecase(),
		GetOpportunityByIdUsecase:      usecases.NewGetOpportunityByIdUsecase(),
	}
}

// Routes mapeia todas as rotas HTTP e seus respectivos métodos vinculados às funções do controller.
func (c *OpportunityV1Controller) Routes() []restserver.Route {
	const baseRoute string = "v1/opportunity"
	const baseRouteWithId string = baseRoute + "/{id}"

	return []restserver.Route{
		{
			URI:      baseRoute,
			Method:   http.MethodGet,
			Function: c.GetAllOpportunity,
			Prefix:   restserver.NoPrefix,
		},
		{
			URI:      baseRoute,
			Method:   http.MethodPost,
			Function: c.CreateOpportunity,
			Prefix:   restserver.NoPrefix,
		},
		{
			URI:      baseRouteWithId,
			Method:   http.MethodPut,
			Function: c.UpdateOpportunity,
			Prefix:   restserver.NoPrefix,
		},
		{
			URI:      baseRouteWithId,
			Method:   http.MethodPatch,
			Function: c.ToogleActiveOpportunity,
			Prefix:   restserver.NoPrefix,
		},
		{
			URI:      baseRouteWithId,
			Method:   http.MethodDelete,
			Function: c.DeleteOpportunity,
			Prefix:   restserver.NoPrefix,
		},
		{
			URI:      baseRouteWithId,
			Method:   http.MethodGet,
			Function: c.GetOpportunityById,
			Prefix:   restserver.NoPrefix,
		},
	}
}

// GetAllOpportunity recebe a requisição HTTP GET de listagem, decodifica os parâmetros de consulta (query params) e aciona o caso de uso correspondente e retorna a lista de oportunidades em formato JSON.
func (c *OpportunityV1Controller) GetAllOpportunity(wctx restserver.WebContext) {
	var params models.OpportunityListParams
	if err := wctx.DecodeQueryParams(&params); err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	result, err := c.GetAllOpportunityUsecase.Execute(wctx.Context(), &params)
	if err != nil {
		wctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	wctx.JsonResponse(http.StatusOK, result)
}

// CreateOpportunity recebe a requisição HTTP POST, decodifica o payload JSON, aciona o caso de uso de criação e devolve a resposta apropriada ao cliente.
func (c *OpportunityV1Controller) CreateOpportunity(wctx restserver.WebContext) {
	var body models.OpportunityCreate
	if err := wctx.DecodeBody(&body); err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	result, err := c.CreateOpportunityUsecase.Execute(wctx.Context(), &body)
	if err != nil {
		switch err.Error() {
		case exceptions.ErrOpportunityAlreadyExists:
			wctx.ErrorResponse(http.StatusConflict, err)
		default:
			wctx.ErrorResponse(http.StatusInternalServerError, err)
		}
		return
	}

	wctx.JsonResponse(http.StatusCreated, result)
}

// UpdateOpportunity extrai o ID da URL, decodifica o corpo da requisição PUT, aciona o caso de uso de atualização e retorna o status de sucesso sem conteúdo (NoContent).
func (c *OpportunityV1Controller) UpdateOpportunity(wctx restserver.WebContext) {
	paramId, err := uuid.Parse(wctx.PathParam("id"))
	if err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	var body models.OpportunityUpdate
	if err := wctx.DecodeBody(&body); err != nil {
		wctx.ErrorResponse(http.StatusUnprocessableEntity, err)
		return
	}

	if err := c.UpdateOpportunityUsecase.Execute(wctx.Context(), paramId, &body); err != nil {
		switch err.Error() {
		case exceptions.ErrOpportunityAlreadyExists:
			wctx.ErrorResponse(http.StatusConflict, err)
		case exceptions.ErrOpportunityNotFound:
			wctx.ErrorResponse(http.StatusNotFound, err)
		default:
			wctx.ErrorResponse(http.StatusInternalServerError, err)
		}
		return
	}

	wctx.EmptyResponse(http.StatusNoContent)
}

// ToogleActiveOpportunity extrai o ID da URL e aciona o caso de uso responsável por inverter o status de ativo/inativo da oportunidade (PATCH).
func (c *OpportunityV1Controller) ToogleActiveOpportunity(wctx restserver.WebContext) {
	paramId, err := uuid.Parse(wctx.PathParam("id"))
	if err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	if err := c.ToogleActiveOpportunityUsecase.Execute(wctx.Context(), paramId); err != nil {
		switch err.Error() {
		case exceptions.ErrOpportunityNotFound:
			wctx.ErrorResponse(http.StatusNotFound, err)
		default:
			wctx.ErrorResponse(http.StatusInternalServerError, err)
		}
		return
	}

	wctx.EmptyResponse(http.StatusNoContent)
}

// DeleteOpportunity extrai o ID da URL e aciona o caso de uso para remover permanentemente a oportunidade (DELETE).
func (c *OpportunityV1Controller) DeleteOpportunity(wctx restserver.WebContext) {
	paramId, err := uuid.Parse(wctx.PathParam("id"))
	if err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	if err := c.DeleteOpportunityUsecase.Execute(wctx.Context(), paramId); err != nil {
		switch err.Error() {
		case exceptions.ErrOpportunityNotFound:
			wctx.ErrorResponse(http.StatusNotFound, err)
		default:
			wctx.ErrorResponse(http.StatusInternalServerError, err)
		}
		return
	}

	wctx.EmptyResponse(http.StatusNoContent)
}

// GetOpportunityById extrai o ID da URL, busca os detalhes da oportunidade específica pelo identificador único e retorna o resultado em JSON (GET por ID).
func (c *OpportunityV1Controller) GetOpportunityById(wctx restserver.WebContext) {
	paramId, err := uuid.Parse(wctx.PathParam("id"))
	
	if err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	result, err := c.GetOpportunityByIdUsecase.Execute(wctx.Context(), paramId)
	if err != nil {
		switch err.Error() {
		case exceptions.ErrOpportunityNotFound:
			wctx.ErrorResponse(http.StatusNotFound, err)
		default:
			wctx.ErrorResponse(http.StatusInternalServerError, err)
		}
		return
	}

	wctx.JsonResponse(http.StatusOK, result)
}

package usecases

import (
	"context"
	"errors"

	"github.com/colibriproject-dev/colibri-sdk-go/pkg/base/logging"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/domain/exceptions"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/domain/models"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/infra/repositories"
)

const (
	// errAnErrorOccurredInGetAllOpportunityUsecaseMsg padroniza a mensagem de log para erros ocorridos neste usecase.
	errAnErrorOccurredInGetAllOpportunityUsecaseMsg string = "an error occurred in GetAllOpportunityUsecase"
)

// IGetAllOpportunityUsecase define o contrato para o caso de uso de listagem e busca de oportunidades.
type IGetAllOpportunityUsecase interface {
	Execute(ctx context.Context, params *models.OpportunityListParams) ([]models.Opportunity, error)
}

// GetAllOpportunityUsecase implementa a lógica de negócio para buscar e listar oportunidades.
type GetAllOpportunityUsecase struct {
	OpportunityRepository repositories.IOpportunityRepository
}

// NewGetAllOpportunityUsecase instancia uma nova estruturação do caso de uso de listagem.
func NewGetAllOpportunityUsecase() *GetAllOpportunityUsecase {
	return &GetAllOpportunityUsecase{
		OpportunityRepository: repositories.NewOpportunityDBRepository(),
	}
}

// Execute aciona o repositório para buscar todas as oportunidades (filtradas ou não pelos parâmetros) e trata eventuais erros de infraestrutura ou log de falhas.
func (u *GetAllOpportunityUsecase) Execute(ctx context.Context, params *models.OpportunityListParams) ([]models.Opportunity, error) {
	result, err := u.OpportunityRepository.FindAll(ctx, params)
	if err != nil {
		logging.Error(ctx).
			Err(err).
			AddParam("step", "OpportunityRepository.FindAll").
			AddParam("params", params).
			Msg(errAnErrorOccurredInGetAllOpportunityUsecaseMsg)

		return nil, errors.New(exceptions.ErrOnFindAllOpportunity)
	}

	return result, nil
}
package usecases

import (
	"context"
	"errors"

	"github.com/colibriproject-dev/colibri-sdk-go/pkg/base/logging"
	"github.com/google/uuid"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/domain/exceptions"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/domain/models"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/infra/repositories"
)

const (
	// errAnErrorOccurredInGetOpportunityByIdUsecaseMsg padroniza a mensagem de log para erros ocorridos neste usecase.
	errAnErrorOccurredInGetOpportunityByIdUsecaseMsg string = "an error occurred in IGetOpportunityByIdUsecase"
)

// IGetOpportunityByIdUsecase define o contrato para o caso de uso de busca de uma oportunidade específica pelo ID.
type IGetOpportunityByIdUsecase interface {
	Execute(ctx context.Context, id uuid.UUID) (*models.Opportunity, error)
}

// GetOpportunityByIdUsecase implementa a lógica de negócio para buscar uma oportunidade com base em seu ID.
type GetOpportunityByIdUsecase struct {
	OpportunityRepository repositories.IOpportunityRepository
}

// NewGetOpportunityByIdUsecase instancia uma nova estruturação do caso de uso de busca por UUID, retornando a interface correspondente.
func NewGetOpportunityByIdUsecase() IGetOpportunityByIdUsecase {
	return &GetOpportunityByIdUsecase{
		OpportunityRepository: repositories.NewOpportunityDBRepository(),
	}
}

// Execute aciona o repositório para localizar a oportunidade pelo UUID informado e trata eventuais erros de infraestrutura ou falhas de consulta.
func (u *GetOpportunityByIdUsecase) Execute(ctx context.Context, id uuid.UUID) (*models.Opportunity, error) {
	result, err := u.OpportunityRepository.FindById(ctx, id)
	if err != nil {
		logging.Error(ctx).
			Err(err).
			AddParam("step", "OpportunityRepository.FindById").
			AddParam("id", id).
			Msg(errAnErrorOccurredInGetOpportunityByIdUsecaseMsg)

		return nil, errors.New(exceptions.ErrOnFindAllOpportunity)
	}

	return result, nil
}
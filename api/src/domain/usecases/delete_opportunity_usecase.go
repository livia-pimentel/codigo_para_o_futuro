package usecases

import (
	"context"
	"errors"

	"github.com/colibriproject-dev/colibri-sdk-go/pkg/base/logging"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/domain/exceptions"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/infra/repositories"
	"github.com/google/uuid"
)

const (
	// errAnErrorOccurredInDeleteOpportunityUsecaseMsg padroniza a mensagem de log para erros ocorridos neste usecase.
	errAnErrorOccurredInDeleteOpportunityUsecaseMsg string = "an error occurred in DeleteOpportunityUsecase"
)

// IDeleteOpportunityUsecase define o contrato para o caso de uso de exclusão de oportunidades.
type IDeleteOpportunityUsecase interface {
	Execute(ctx context.Context, id uuid.UUID) error
}

// DeleteOpportunityUsecase implementa a lógica de negócio para remover permanentemente uma oportunidade do sistema.
type DeleteOpportunityUsecase struct {
	OpportunityRepository repositories.IOpportunityRepository
}
// NewDeleteOpportunityUsecase instancia uma nova estruturação do caso de uso de exclusão.
func NewDeleteOpportunityUsecase() *DeleteOpportunityUsecase {
	return &DeleteOpportunityUsecase{
		OpportunityRepository: repositories.NewOpportunityDBRepository(),
	}
}

// Execute valida se a oportunidade realmente existe pelo ID informado e, caso positivo, faz a exclusão definitiva no banco de dados.
func (u *DeleteOpportunityUsecase) Execute(ctx context.Context, id uuid.UUID) error {
	if err := u.existsOpportunityById(ctx, id); err != nil {
		return err
	}

	return u.DeleteOpportunity(ctx, id)
}

// existsOpportunityById verifica se existe uma oportunidade cadastrada com o UUID fornecido.
func (u *DeleteOpportunityUsecase) existsOpportunityById(ctx context.Context, id uuid.UUID) error {
	exists, err := u.OpportunityRepository.ExistsById(ctx, id)
	if err != nil {
		logging.Error(ctx).
			Err(err).
			AddParam("step", "OpportunityRepository.ExistsById").
			AddParam("id", id).
			Msg(errAnErrorOccurredInDeleteOpportunityUsecaseMsg)

		return errors.New(exceptions.ErrOnExistsOpportunityById)
	}

	if (exists == nil) || !*exists {
		return errors.New(exceptions.ErrOpportunityNotFound)
	}

	return nil
}

// DeleteOpportunity executa a remoção física ou lógica da oportunidade no repositório.
func (u *DeleteOpportunityUsecase) DeleteOpportunity(ctx context.Context, id uuid.UUID) error {
	if err := u.OpportunityRepository.Delete(ctx, id); err != nil {
		logging.Error(ctx).
			Err(err).
			AddParam("step", "OpportunityRepository.Delete").
			AddParam("id", id).
			Msg(errAnErrorOccurredInDeleteOpportunityUsecaseMsg)
		return errors.New(exceptions.ErrOnDeleteOpportunity)
	}

	return nil
}
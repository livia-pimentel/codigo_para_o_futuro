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
	// errAnErrorOccurredInToogleActiveOpportunityUsecaseMsg padroniza a mensagem de log para erros ocorridos neste usecase.
	errAnErrorOccurredInToogleActiveOpportunityUsecaseMsg string = "an error occurred in ToogleActiveOpportunityUsecase"
)

// IToogleActiveOpportunityUsecase define o contrato para o caso de uso de alternância de status (ativo/inativo) de uma oportunidade.
type IToogleActiveOpportunityUsecase interface {
	Execute(ctx context.Context, id uuid.UUID) error
}

// ToogleActiveOpportunityUsecase implementa a lógica de negócio para inverter o estado de ativação de uma oportunidade.
type ToogleActiveOpportunityUsecase struct {
	OpportunityRepository repositories.IOpportunityRepository
}

// NewToogleActiveOpportunityUsecase instancia uma nova estruturação do caso de uso de alternância de status.
func NewToogleActiveOpportunityUsecase() *ToogleActiveOpportunityUsecase {
	return &ToogleActiveOpportunityUsecase{
		OpportunityRepository: repositories.NewOpportunityDBRepository(),
	}
}

// Execute valida se a oportunidade existe pelo ID fornecido e, caso exista, prossegue com a inversão do seu status ativo/inativo no banco de dados.
func (u *ToogleActiveOpportunityUsecase) Execute(ctx context.Context, id uuid.UUID) error {
	if err := u.existsOpportunityById(ctx, id); err != nil {
		return err
	}

	return u.toogleActiveOpportunity(ctx, id)
}

// existsOpportunityById verifica se existe uma oportunidade cadastrada com o UUID fornecido.
func (u *ToogleActiveOpportunityUsecase) existsOpportunityById(ctx context.Context, id uuid.UUID) error {
	exists, err := u.OpportunityRepository.ExistsById(ctx, id)
	if err != nil {
		logging.Error(ctx).
			Err(err).
			AddParam("step", "OpportunityRepository.ExistsById").
			AddParam("id", id).
			Msg(errAnErrorOccurredInToogleActiveOpportunityUsecaseMsg)

		return errors.New(exceptions.ErrOnExistsOpportunityById)
	}

	if (exists == nil) || !*exists {
		return errors.New(exceptions.ErrOpportunityNotFound)
	}

	return nil
}

// toogleActiveOpportunity aciona o repositório para alterar o estado do registro (de ativo para inativo ou vice-versa).
func (u *ToogleActiveOpportunityUsecase) toogleActiveOpportunity(ctx context.Context, id uuid.UUID) error {
	if err := u.OpportunityRepository.ToogleActive(ctx, id); err != nil {
		logging.Error(ctx).
			Err(err).
			AddParam("step", "OpportunityRepository.ToogleActive").
			AddParam("id", id).
			Msg(errAnErrorOccurredInToogleActiveOpportunityUsecaseMsg)
		return errors.New(exceptions.ErrOnToogleActiveOpportunity)
	}

	return nil
}
package usecases

import (
	"context"
	"errors"

	"github.com/colibriproject-dev/colibri-sdk-go/pkg/base/logging"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/domain/exceptions"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/domain/models"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/infra/repositories"
	"github.com/google/uuid"
)

const (
	// errAnErrorOccurredInUpdateOpportunityUsecaseMsg padroniza a mensagem de log para erros ocorridos neste usecase.
	errAnErrorOccurredInUpdateOpportunityUsecaseMsg string = "an error occurred in UpdateOpportunityUsecase"
)

// IUpdateOpportunityUsecase define o contrato para o caso de uso de atualização de oportunidades.
type IUpdateOpportunityUsecase interface {
	Execute(ctx context.Context, id uuid.UUID, model *models.OpportunityUpdate) error
}

// UpdateOpportunityUsecase implementa a lógica de negócio para modificar os dados de uma oportunidade existente.
type UpdateOpportunityUsecase struct {
	OpportunityRepository repositories.IOpportunityRepository
}

// NewUpdateOpportunityUsecase instancia uma nova estruturação do caso de uso de atualização.
func NewUpdateOpportunityUsecase() *UpdateOpportunityUsecase {
	return &UpdateOpportunityUsecase{
		OpportunityRepository: repositories.NewOpportunityDBRepository(),
	}
}

// Execute realiza a sanitização dos campos de texto, valida se a oportunidade existe pelo ID, verifica se o novo título já não pertence a outro registro (evitando duplicidade) e executa a atualização.
func (u *UpdateOpportunityUsecase) Execute(ctx context.Context, id uuid.UUID, model *models.OpportunityUpdate) error {

	// Sanitiza os dados
	model.Title = models.CleanString(model.Title)
	model.Company = models.CleanString(model.Company)
	model.Category = models.CleanString(model.Category)

	if err := u.existsOpportunityById(ctx, id); err != nil {
		return err
	}

	if err := u.findOpportunityByTitle(ctx, id, model); err != nil {
		return err
	}

	return u.updateOpportunity(ctx, id, model)
}

// existsOpportunityById verifica se existe uma oportunidade cadastrada com o UUID fornecido.
func (u *UpdateOpportunityUsecase) existsOpportunityById(ctx context.Context, id uuid.UUID) error {
	exists, err := u.OpportunityRepository.ExistsById(ctx, id)
	if err != nil {
		logging.Error(ctx).
			Err(err).
			AddParam("step", "OpportunityRepository.ExistsById").
			AddParam("id", id).
			Msg(errAnErrorOccurredInUpdateOpportunityUsecaseMsg)

		return errors.New(exceptions.ErrOnExistsOpportunityById)
	}

	if (exists == nil) || !*exists {
		return errors.New(exceptions.ErrOpportunityNotFound)
	}

	return nil
}

// findOpportunityByTitle valida se já existe outra oportunidade com o mesmo título para garantir que não ocorra conflito de unicidade com um ID diferente.
func (u *UpdateOpportunityUsecase) findOpportunityByTitle(ctx context.Context, id uuid.UUID, model *models.OpportunityUpdate) error {
	result, err := u.OpportunityRepository.FindByTitle(ctx, model.Title)
	if err != nil {
		logging.Error(ctx).
			Err(err).
			AddParam("step", "OpportunityRepository.FindByTitle").
			AddParam("id", id).
			AddParam("model", model).
			Msg(errAnErrorOccurredInUpdateOpportunityUsecaseMsg)
		return errors.New(exceptions.ErrOnFindOpportunityByTitle)
	}

	if (result != nil) && (result.Title == model.Title) && (result.ID != id) {
		return errors.New(exceptions.ErrOpportunityAlreadyExists)
	}

	return nil
}

// updateOpportunity executa a persistência das alterações da oportunidade no repositório.
func (u *UpdateOpportunityUsecase) updateOpportunity(ctx context.Context, id uuid.UUID, model *models.OpportunityUpdate) error {
	if err := u.OpportunityRepository.Update(ctx, id, model); err != nil {
		logging.Error(ctx).
			Err(err).
			AddParam("step", "OpportunityRepository.Update").
			AddParam("id", id).
			AddParam("model", model).
			Msg(errAnErrorOccurredInUpdateOpportunityUsecaseMsg)
		return errors.New(exceptions.ErrOnUpdateOpportunity)
	}

	return nil
}
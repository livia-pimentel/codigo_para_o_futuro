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
	// errAnErrorOccurredInCreateOpportunityUsecaseMsg padroniza a mensagem de log para erros ocorridos neste usecase.
	errAnErrorOccurredInCreateOpportunityUsecaseMsg string = "an error occurred in CreateOpportunityUsecase"
)

// ICreateOpportunityUsecase define o contrato para o caso de uso de criação de oportunidades.
type ICreateOpportunityUsecase interface {
	Execute(ctx context.Context, model *models.OpportunityCreate) (*models.OpportunityCreated, error)
}

// CreateOpportunityUsecase implementa a lógica de negócio para cadastrar uma nova oportunidade.
type CreateOpportunityUsecase struct {
	OpportunityRepository repositories.IOpportunityRepository
}

// NewCreateOpportunityUsecase instancia uma nova estruturação do caso de uso de criação.
func NewCreateOpportunityUsecase() *CreateOpportunityUsecase{
	return &CreateOpportunityUsecase{
		OpportunityRepository: repositories.NewOpportunityDBRepository(),
	}
}

// Execute realiza a sanitização dos dados, valida se a oportunidade já existe para a mesma empresa
// e faz a inserção no banco de dados.
func (u *CreateOpportunityUsecase) Execute(ctx context.Context, model *models.OpportunityCreate) (*models.OpportunityCreated, error) {
	
	// Sanitiza os dados
	model.Title = models.CleanString(model.Title)
	model.Company = models.CleanString(model.Company)
	model.Category = models.CleanString(model.Category)

	if err := u.existsOpportunityByTitle(ctx, model); err != nil {
		return nil, err
	}

	return u.insertOpportunity(ctx, model)
}

// existsOpportunityByTitle verifica se já existe uma oportunidade com o mesmo título vinculada à mesma empresa.
func (u *CreateOpportunityUsecase) existsOpportunityByTitle(ctx context.Context, model *models.OpportunityCreate) error {
	exists, err := u.OpportunityRepository.ExistsByTitleAndCompany(ctx, model.Title, model.Company)
	if err != nil {
		logging.Error(ctx).
			Err(err).
			AddParam("step", "OpportunityRepository.ExistsByTitleAndCompany").
			AddParam("model", model).
			Msg(errAnErrorOccurredInCreateOpportunityUsecaseMsg)

		return errors.New(exceptions.ErrOnExistsOpportunityByTitle)
	}

	if exists != nil && *exists {
		return errors.New(exceptions.ErrOpportunityAlreadyExists)
	}

	return nil
}

// insertOpportunity executa a persistência da nova oportunidade no repositório.
func (u *CreateOpportunityUsecase) insertOpportunity(ctx context.Context, model *models.OpportunityCreate) (*models.OpportunityCreated, error) {
	result, err := u.OpportunityRepository.Insert(ctx, model)
	if err != nil {
		logging.Error(ctx).
			Err(err).
			AddParam("step", "OpportunityRepository.Insert").
			AddParam("model", model).
			Msg(errAnErrorOccurredInCreateOpportunityUsecaseMsg)
		return nil, errors.New(exceptions.ErrOnInsertOpportunity)
	}

	return result, nil
}
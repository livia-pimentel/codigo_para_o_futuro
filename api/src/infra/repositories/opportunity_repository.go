package repositories

import (
	"context"

	"github.com/colibriproject-dev/colibri-sdk-go/pkg/database/sqlDB"
	"github.com/google/uuid"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/domain/models"
)

const (
	opportunityBaseQuery          = `SELECT id, title, company, category, active, created_at, updated_at FROM opportunity`
	findAllOpportunityQuery       = opportunityBaseQuery + ` WHERE ($1 = '' OR (title ILIKE CONCAT('%', $1, '%'))) `
	findOpportunityByTitleQuery   = opportunityBaseQuery + ` WHERE title = $1`
	existsOpportunityByIdQuery    = `SELECT EXISTS(SELECT 1 FROM opportunity WHERE id = $1)`
	existsByTitleAndCompanyQuery = `SELECT EXISTS(SELECT 1 FROM opportunity WHERE title = $1 AND company = $2)`
	insertOpportunityQuery        = `INSERT INTO opportunity (title, company, category) VALUES ($1, $2, $3) RETURNING id`
	updateOpportunityQuery        = `UPDATE opportunity SET title = $1, company = $2, category = $3, updated_at = NOW() WHERE id = $4`
	toogleActiveOpportunityQuery  = `UPDATE opportunity SET active = NOT active, updated_at = NOW() WHERE id = $1`
	deleteOpportunityQuery        = `DELETE FROM opportunity WHERE id = $1`
	findOpportunityById 		  = `SELECT id, title, company, category, active, created_at, updated_at FROM opportunity WHERE id = $1`
)

// IOpportunityRepository define o contrato (interface) com todas as operações de persistência e consultas ao banco de dados para a entidade Opportunity.
type IOpportunityRepository interface {
	FindAll(ctx context.Context, params *models.OpportunityListParams) ([]models.Opportunity, error)
	FindByTitle(ctx context.Context, title string) (*models.Opportunity, error)
	ExistsById(ctx context.Context, id uuid.UUID) (*bool, error)
	// Substituir o que era existsOpportunityByTitleQuery para ExistsByTitleAndCompany para verificar o title e company antes de cadastrar
	// existsOpportunityByTitleQuery = `SELECT EXISTS(SELECT 1 FROM opportunity WHERE title = $1)`
	ExistsByTitleAndCompany(ctx context.Context, title, company string) (*bool, error)
	Insert(ctx context.Context, model *models.OpportunityCreate) (*models.OpportunityCreated, error)
	Update(ctx context.Context, id uuid.UUID, model *models.OpportunityUpdate) error
	ToogleActive(ctx context.Context, id uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (*models.Opportunity, error)
}

// OpportunityDBRepository implementa a interface IOpportunityRepository utilizando o driver de banco de dados SQL.
type OpportunityDBRepository struct{}

// NewOpportunityDBRepository instancia um novo repositório de oportunidades conectado ao banco de dados.
func NewOpportunityDBRepository() *OpportunityDBRepository {
	return &OpportunityDBRepository{}
}

// FindAll busca todas as oportunidades cadastradas, aplicando opcionalmente um filtro por título se fornecido.
func (r *OpportunityDBRepository) FindAll(ctx context.Context, params *models.OpportunityListParams) ([]models.Opportunity, error) {
	return sqlDB.NewQuery[models.Opportunity](ctx, findAllOpportunityQuery, params.Title).Many()
}

// FindByTitle busca uma oportunidade específica correspondente ao título informado.
func (r *OpportunityDBRepository) FindByTitle(ctx context.Context, title string) (*models.Opportunity, error) {
	return sqlDB.NewQuery[models.Opportunity](ctx, findOpportunityByTitleQuery, title).One()
}

// ExistsById verifica no banco de dados se existe algum registro associado ao UUID informado.
func (r *OpportunityDBRepository) ExistsById(ctx context.Context, id uuid.UUID) (*bool, error) {
	return sqlDB.NewQuery[bool](ctx, existsOpportunityByIdQuery, id).One()
}

// ExistsByTitleAndCompany verifica se já existe uma oportunidade cadastrada com o mesmo título vinculada à mesma empresa.
func (r *OpportunityDBRepository) ExistsByTitleAndCompany(ctx context.Context, title, company string) (*bool, error) {
	return sqlDB.NewQuery[bool](ctx, existsByTitleAndCompanyQuery, title, company).One()
}

// Insert insere uma nova oportunidade no banco de dados utilizando os dados informados e retorna o ID gerado.
func (r *OpportunityDBRepository) Insert(ctx context.Context, model *models.OpportunityCreate) (*models.OpportunityCreated, error) {
	return sqlDB.NewQuery[models.OpportunityCreated](ctx, insertOpportunityQuery,
		model.Title,
		model.Company,
		model.Category,
	).One()
}

// Update atualiza as informações de uma oportunidade existente com base no ID e nos novos dados fornecidos.
func (r *OpportunityDBRepository) Update(ctx context.Context, id uuid.UUID, model *models.OpportunityUpdate) error {
	return sqlDB.NewStatement(ctx, updateOpportunityQuery,
		model.Title,
		model.Company,
		model.Category,
		id,
	).Execute()
}

// ToogleActive inverte o status atual da oportunidade (de ativa para inativa ou vice-versa) com base no ID.
func (r *OpportunityDBRepository) ToogleActive(ctx context.Context, id uuid.UUID) error {
	return sqlDB.NewStatement(ctx, toogleActiveOpportunityQuery, id).Execute()
}

// Delete remove permanentemente do banco de dados o registro da oportunidade correspondente ao ID.
func (r *OpportunityDBRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return sqlDB.NewStatement(ctx, deleteOpportunityQuery, id).Execute()
}

// FindById busca e retorna os detalhes completos de uma oportunidade específica a partir do seu identificador único.
func (r *OpportunityDBRepository) FindById(ctx context.Context, id uuid.UUID) (*models.Opportunity, error) {
    return sqlDB.NewQuery[models.Opportunity](ctx, findOpportunityById, id).One()
}
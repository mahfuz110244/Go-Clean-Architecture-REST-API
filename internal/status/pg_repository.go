//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package status

import (
	"context"

	"github.com/google/uuid"

	"github.com/mahfuz110244/api-mc/internal/models"
	"github.com/mahfuz110244/api-mc/pkg/utils"
)

// Status Repository
type Repository interface {
	Create(ctx context.Context, status *models.Status) (*models.Status, error)
	Update(ctx context.Context, status *models.Status) (*models.Status, error)
	GetStatusByID(ctx context.Context, statusID uuid.UUID) (*models.Status, error)
	Delete(ctx context.Context, statusID uuid.UUID) error
	GetStatus(ctx context.Context, pq *utils.PaginationQuery) (*models.StatusList, error)
	SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.StatusList, error)
}

//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package status

import (
	"context"

	"github.com/google/uuid"

	"github.com/mahfuz110244/api-mc/internal/models"
	"github.com/mahfuz110244/api-mc/pkg/utils"
)

// Status use case
type UseCase interface {
	Create(ctx context.Context, status *models.Status) (*models.Status, error)
	Update(ctx context.Context, status *models.Status) (*models.Status, error)
	GetStatusByID(ctx context.Context, statusID uuid.UUID) (*models.StatusBase, error)
	Delete(ctx context.Context, statusID uuid.UUID) error
	GetStatus(ctx context.Context, pq *utils.PaginationQuery) (*models.StatusList, error)
	SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.StatusList, error)
}

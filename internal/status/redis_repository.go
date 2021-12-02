//go:generate mockgen -source redis_repository.go -destination mock/redis_repository_mock.go -package mock
package status

import (
	"context"

	"github.com/mahfuz110244/api-mc/internal/models"
)

// Status redis repository
type RedisRepository interface {
	GetStatusByIDCtx(ctx context.Context, key string) (*models.StatusBase, error)
	SetStatusCtx(ctx context.Context, key string, seconds int, status *models.StatusBase) error
	DeleteStatusCtx(ctx context.Context, key string) error
}

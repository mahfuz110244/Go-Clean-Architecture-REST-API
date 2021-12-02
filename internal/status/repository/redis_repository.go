package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/mahfuz110244/api-mc/internal/models"
	"github.com/mahfuz110244/api-mc/internal/status"
)

// Status redis repository
type statusRedisRepo struct {
	redisClient *redis.Client
}

// Status redis repository constructor
func NewStatusRedisRepo(redisClient *redis.Client) status.RedisRepository {
	return &statusRedisRepo{redisClient: redisClient}
}

// Get new by id
func (n *statusRedisRepo) GetStatusByIDCtx(ctx context.Context, key string) (*models.StatusBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRedisRepo.GetStatusByIDCtx")
	defer span.Finish()

	statusBytes, err := n.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "statusRedisRepo.GetStatusByIDCtx.redisClient.Get")
	}
	statusBase := &models.StatusBase{}
	if err = json.Unmarshal(statusBytes, statusBase); err != nil {
		return nil, errors.Wrap(err, "statusRedisRepo.GetStatusByIDCtx.json.Unmarshal")
	}

	return statusBase, nil
}

// Cache status item
func (n *statusRedisRepo) SetStatusCtx(ctx context.Context, key string, seconds int, status *models.StatusBase) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRedisRepo.SetStatusCtx")
	defer span.Finish()

	statusBytes, err := json.Marshal(status)
	if err != nil {
		return errors.Wrap(err, "statusRedisRepo.SetStatusCtx.json.Marshal")
	}
	if err = n.redisClient.Set(ctx, key, statusBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "statusRedisRepo.SetStatusCtx.redisClient.Set")
	}
	return nil
}

// Delete new item from cache
func (n *statusRedisRepo) DeleteStatusCtx(ctx context.Context, key string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRedisRepo.DeleteStatusCtx")
	defer span.Finish()

	if err := n.redisClient.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "statusRedisRepo.DeleteStatusCtx.redisClient.Del")
	}
	return nil
}

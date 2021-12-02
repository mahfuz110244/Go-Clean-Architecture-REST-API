package repository

import (
	"context"
	"log"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/mahfuz110244/api-mc/internal/models"
	"github.com/mahfuz110244/api-mc/internal/status"
)

func SetupRedis() status.RedisRepository {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	statusRedisRepo := NewStatusRedisRepo(client)
	return statusRedisRepo
}

func TestStatusRedisRepo_SetStatusCtx(t *testing.T) {
	t.Parallel()

	statusRedisRepo := SetupRedis()

	t.Run("SetStatusCtx", func(t *testing.T) {
		statusUID := uuid.New()
		key := "key"
		n := &models.StatusBase{
			StatusID: statusUID,
			Title:    "Title",
			Content:  "Content",
		}

		err := statusRedisRepo.SetStatusCtx(context.Background(), key, 10, n)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}

func TestStatusRedisRepo_GetStatusByIDCtx(t *testing.T) {
	t.Parallel()

	statusRedisRepo := SetupRedis()

	t.Run("GetStatusByIDCtx", func(t *testing.T) {
		statusUID := uuid.New()
		key := "key"
		n := &models.StatusBase{
			StatusID: statusUID,
			Title:    "Title",
			Content:  "Content",
		}

		statusBase, err := statusRedisRepo.GetStatusByIDCtx(context.Background(), key)
		require.Nil(t, statusBase)
		require.NotNil(t, err)

		err = statusRedisRepo.SetStatusCtx(context.Background(), key, 10, n)
		require.NoError(t, err)
		require.Nil(t, err)

		statusBase, err = statusRedisRepo.GetStatusByIDCtx(context.Background(), key)
		require.NoError(t, err)
		require.Nil(t, err)
		require.NotNil(t, statusBase)
	})
}

func TestStatusRedisRepo_DeleteStatusCtx(t *testing.T) {
	t.Parallel()

	statusRedisRepo := SetupRedis()

	t.Run("SetStatusCtx", func(t *testing.T) {
		key := "key"

		err := statusRedisRepo.DeleteStatusCtx(context.Background(), key)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}

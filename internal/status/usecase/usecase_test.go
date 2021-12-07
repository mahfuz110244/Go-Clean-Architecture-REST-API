package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"

	"github.com/mahfuz110244/api-mc/internal/models"
	"github.com/mahfuz110244/api-mc/internal/status/mock"
	"github.com/mahfuz110244/api-mc/pkg/logger"
	"github.com/mahfuz110244/api-mc/pkg/utils"
)

func TestStatusUC_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockStatusRepo := mock.NewMockRepository(ctrl)
	statusUC := NewStatusUseCase(nil, mockStatusRepo, nil, apiLogger)

	userUID := uuid.New()

	status := &models.Status{
		Name:        "estimatee",
		Description: "estimate",
	}

	user := &models.User{
		UserID: userUID,
	}

	ctx := context.WithValue(context.Background(), utils.UserCtxKey{}, user)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "statusUC.Create")
	defer span.Finish()

	mockStatusRepo.EXPECT().Create(ctxWithTrace, gomock.Eq(status)).Return(status, nil)

	createdStatus, err := statusUC.Create(ctx, status)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, createdStatus)
}

func TestStatusUC_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockStatusRepo := mock.NewMockRepository(ctrl)
	mockRedisRepo := mock.NewMockRedisRepository(ctrl)
	statusUC := NewStatusUseCase(nil, mockStatusRepo, mockRedisRepo, apiLogger)

	userUID := uuid.New()
	statusUID := uuid.New()
	status := &models.Status{
		ID:          statusUID,
		Name:        "estimate",
		Description: "estimate",
		CreatedBy:   userUID,
		UpdatedBy:   userUID,
	}

	statusBase := &models.StatusBase{
		ID:          statusUID,
		Name:        "estimate",
		Description: "estimate",
		CreatedBy:   userUID,
		UpdatedBy:   userUID,
	}

	user := &models.User{
		UserID: userUID,
	}

	cacheKey := fmt.Sprintf("%s: %s", basePrefix, status.ID)

	ctx := context.WithValue(context.Background(), utils.UserCtxKey{}, user)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "statusUC.Update")
	defer span.Finish()

	mockStatusRepo.EXPECT().GetStatusByID(ctxWithTrace, gomock.Eq(status.ID)).Return(statusBase, nil)
	mockStatusRepo.EXPECT().Update(ctxWithTrace, gomock.Eq(status)).Return(status, nil)
	mockRedisRepo.EXPECT().DeleteStatusCtx(ctxWithTrace, gomock.Eq(cacheKey)).Return(nil)

	updatedStatus, err := statusUC.Update(ctx, status)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, updatedStatus)
}

func TestStatusUC_GetStatusByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockStatusRepo := mock.NewMockRepository(ctrl)
	mockRedisRepo := mock.NewMockRedisRepository(ctrl)
	statusUC := NewStatusUseCase(nil, mockStatusRepo, mockRedisRepo, apiLogger)

	statusUID := uuid.New()
	statusBase := &models.StatusBase{
		ID: statusUID,
	}
	ctx := context.Background()
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "statusUC.GetStatusByID")
	defer span.Finish()
	cacheKey := fmt.Sprintf("%s: %s", basePrefix, statusUID)

	mockRedisRepo.EXPECT().GetStatusByIDCtx(ctxWithTrace, gomock.Eq(cacheKey)).Return(nil, nil)
	mockStatusRepo.EXPECT().GetStatusByID(ctxWithTrace, gomock.Eq(statusUID)).Return(statusBase, nil)
	mockRedisRepo.EXPECT().SetStatusCtx(ctxWithTrace, cacheKey, cacheDuration, statusBase).Return(nil)

	statusByID, err := statusUC.GetStatusByID(ctx, statusBase.ID)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, statusByID)
}

func TestStatusUC_Delete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockStatusRepo := mock.NewMockRepository(ctrl)
	mockRedisRepo := mock.NewMockRedisRepository(ctrl)
	statusUC := NewStatusUseCase(nil, mockStatusRepo, mockRedisRepo, apiLogger)

	statusUID := uuid.New()
	userUID := uuid.New()
	statusBase := &models.StatusBase{
		ID: statusUID,
	}
	cacheKey := fmt.Sprintf("%s: %s", basePrefix, statusUID)

	user := &models.User{
		UserID: userUID,
	}

	ctx := context.WithValue(context.Background(), utils.UserCtxKey{}, user)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "statusUC.Delete")
	defer span.Finish()

	mockStatusRepo.EXPECT().GetStatusByID(ctxWithTrace, gomock.Eq(statusBase.ID)).Return(statusBase, nil)
	mockStatusRepo.EXPECT().Delete(ctxWithTrace, gomock.Eq(statusUID)).Return(nil)
	mockRedisRepo.EXPECT().DeleteStatusCtx(ctxWithTrace, gomock.Eq(cacheKey)).Return(nil)

	err := statusUC.Delete(ctx, statusBase.ID)
	require.NoError(t, err)
	require.Nil(t, err)
}

func TestStatusUC_GetStatus(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockStatusRepo := mock.NewMockRepository(ctrl)
	mockRedisRepo := mock.NewMockRedisRepository(ctrl)
	statusUC := NewStatusUseCase(nil, mockStatusRepo, mockRedisRepo, apiLogger)

	ctx := context.Background()
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "statusUC.GetStatus")
	defer span.Finish()

	query := &utils.PaginationQuery{
		Size:    10,
		Page:    1,
		OrderBy: "",
	}

	statusList := &models.StatusList{}

	mockStatusRepo.EXPECT().GetStatus(ctxWithTrace, query).Return(statusList, nil)

	status, err := statusUC.GetStatus(ctx, query)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, status)
}

func TestStatusUC_SearchByTitle(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockStatusRepo := mock.NewMockRepository(ctrl)
	mockRedisRepo := mock.NewMockRedisRepository(ctrl)
	statusUC := NewStatusUseCase(nil, mockStatusRepo, mockRedisRepo, apiLogger)

	ctx := context.Background()
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "statusUC.SearchByTitle")
	defer span.Finish()
	query := &utils.PaginationQuery{
		Size:    10,
		Page:    1,
		OrderBy: "",
	}

	statusList := &models.StatusList{}
	title := "title"

	mockStatusRepo.EXPECT().SearchByTitle(ctxWithTrace, title, query).Return(statusList, nil)

	status, err := statusUC.SearchByTitle(ctx, title, query)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, status)
}

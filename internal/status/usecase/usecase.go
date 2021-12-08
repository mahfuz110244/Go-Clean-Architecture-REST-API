package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/mahfuz110244/api-mc/config"
	"github.com/mahfuz110244/api-mc/internal/models"
	"github.com/mahfuz110244/api-mc/internal/status"
	"github.com/mahfuz110244/api-mc/pkg/httpErrors"
	"github.com/mahfuz110244/api-mc/pkg/logger"
	"github.com/mahfuz110244/api-mc/pkg/utils"
)

const (
	basePrefix    = "api-status:"
	cacheDuration = 3600
)

// Status UseCase
type statusUC struct {
	cfg        *config.Config
	statusRepo status.Repository
	redisRepo  status.RedisRepository
	logger     logger.Logger
}

// Status UseCase constructor
func NewStatusUseCase(cfg *config.Config, statusRepo status.Repository, redisRepo status.RedisRepository, logger logger.Logger) status.UseCase {
	return &statusUC{cfg: cfg, statusRepo: statusRepo, redisRepo: redisRepo, logger: logger}
}

// Create status
func (u *statusUC) Create(ctx context.Context, status *models.Status) (*models.Status, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusUC.Create")
	defer span.Finish()

	user, err := utils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, httpErrors.NewUnauthorizedError(errors.WithMessage(err, "statusUC.Create.GetUserFromCtx"))
	}

	status.CreatedBy = user.UserID
	status.UpdatedBy = user.UserID

	if err = utils.ValidateStruct(ctx, status); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.WithMessage(err, "statusUC.Create.ValidateStruct"))
	}

	n, err := u.statusRepo.Create(ctx, status)
	if err != nil {
		return nil, err
	}

	return n, err
}

// Update status item
func (u *statusUC) Update(ctx context.Context, status *models.Status) (*models.Status, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusUC.Update")
	defer span.Finish()

	statusByID, err := u.statusRepo.GetStatusByID(ctx, status.ID)
	if err != nil {
		return nil, err
	}
	fmt.Println(statusByID)
	// if statusByID == nil {

	// }

	// if err = utils.ValidateIsOwner(ctx, statusByID.AuthorID.String(), u.logger); err != nil {
	// 	return nil, httpErrors.NewRestError(http.StatusForbidden, "Forbidden", errors.Wrap(err, "statusUC.Update.ValidateIsOwner"))
	// }

	user, err := utils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, httpErrors.NewUnauthorizedError(errors.WithMessage(err, "statusUC.Update.GetUserFromCtx"))
	}
	status.UpdatedBy = user.UserID

	updatedUser, err := u.statusRepo.Update(ctx, status)
	if err != nil {
		return nil, err
	}

	if err = u.redisRepo.DeleteStatusCtx(ctx, u.getKeyWithPrefix(status.ID.String())); err != nil {
		u.logger.Errorf("statusUC.Update.DeleteStatusCtx: %v", err)
	}

	return updatedUser, nil
}

// Get status by id
func (u *statusUC) GetStatusByID(ctx context.Context, statusID uuid.UUID) (*models.Status, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusUC.GetStatusByID")
	defer span.Finish()

	status, err := u.redisRepo.GetStatusByIDCtx(ctx, u.getKeyWithPrefix(statusID.String()))
	if err != nil {
		u.logger.Errorf("statusUC.GetStatusByID.GetStatusByIDCtx: %v", err)
	}
	if status != nil {
		return status, nil
	}

	n, err := u.statusRepo.GetStatusByID(ctx, statusID)
	if err != nil {
		return nil, err
	}

	if err = u.redisRepo.SetStatusCtx(ctx, u.getKeyWithPrefix(statusID.String()), cacheDuration, n); err != nil {
		u.logger.Errorf("statusUC.GetStatusByID.SetStatusCtx: %s", err)
	}

	return n, nil
}

// Soft Delete status
func (u *statusUC) Delete(ctx context.Context, statusID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusUC.Delete")
	defer span.Finish()

	statusByID, err := u.statusRepo.GetStatusByID(ctx, statusID)
	if err != nil {
		return err
	}
	fmt.Println(statusByID)

	// if err = utils.ValidateIsOwner(ctx, statusByID.CreatedBy.String(), u.logger); err != nil {
	// 	return httpErrors.NewRestError(http.StatusForbidden, "Forbidden", errors.Wrap(err, "statusUC.Delete.ValidateIsOwner"))
	// }

	if err = u.statusRepo.Delete(ctx, statusID); err != nil {
		return err
	}

	if err = u.redisRepo.DeleteStatusCtx(ctx, u.getKeyWithPrefix(statusID.String())); err != nil {
		u.logger.Errorf("statusUC.Delete.DeleteStatusCtx: %v", err)
	}

	return nil
}

// Get status
func (u *statusUC) GetStatus(ctx context.Context, pq *utils.PaginationQuery) (*models.StatusList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusUC.GetStatus")
	defer span.Finish()

	return u.statusRepo.GetStatus(ctx, pq)
}

// Find nes by title
func (u *statusUC) SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.StatusList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusUC.SearchByTitle")
	defer span.Finish()

	return u.statusRepo.SearchByTitle(ctx, title, query)
}

func (u *statusUC) getKeyWithPrefix(statusID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, statusID)
}

package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/mahfuz110244/api-mc/internal/models"
	"github.com/mahfuz110244/api-mc/internal/status"
	"github.com/mahfuz110244/api-mc/pkg/utils"
)

// Status Repository
type statusRepo struct {
	db *sqlx.DB
}

// Status repository constructor
func NewStatusRepository(db *sqlx.DB) status.Repository {
	return &statusRepo{db: db}
}

// Create status
func (r *statusRepo) Create(ctx context.Context, status *models.Status) (*models.Status, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRepo.Create")
	defer span.Finish()

	var n models.Status
	if err := r.db.QueryRowxContext(
		ctx,
		createStatus,
		&status.Name,
		&status.Description,
		&status.CreatedBy,
		&status.UpdatedBy,
	).StructScan(&n); err != nil {
		return nil, errors.Wrap(err, "statusRepo.Create.QueryRowxContext")
	}

	return &n, nil
}

// Update status item
func (r *statusRepo) Update(ctx context.Context, status *models.Status) (*models.Status, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRepo.Update")
	defer span.Finish()

	var n models.Status
	if err := r.db.QueryRowxContext(
		ctx,
		updateStatus,
		&status.Name,
		&status.Description,
		&status.Active,
		&status.OrderNumber,
		&status.UpdatedBy,
		&status.ID,
	).StructScan(&n); err != nil {
		return nil, errors.Wrap(err, "statusRepo.Update.QueryRowxContext")
	}

	return &n, nil
}

// Get single status by id
func (r *statusRepo) GetStatusByID(ctx context.Context, statusID uuid.UUID) (*models.StatusBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRepo.GetStatusByID")
	defer span.Finish()

	n := &models.StatusBase{}
	if err := r.db.GetContext(ctx, n, getStatusByID, statusID); err != nil {
		return nil, errors.Wrap(err, "statusRepo.GetStatusByID.GetContext")
	}

	return n, nil
}

// Delete status by id
func (r *statusRepo) Delete(ctx context.Context, statusID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRepo.Delete")
	defer span.Finish()

	result, err := r.db.ExecContext(ctx, deleteStatus, statusID)
	if err != nil {
		return errors.Wrap(err, "statusRepo.Delete.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "statusRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "statusRepo.Delete.rowsAffected")
	}

	return nil
}

// Get status
func (r *statusRepo) GetStatus(ctx context.Context, pq *utils.PaginationQuery) (*models.StatusList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRepo.GetStatus")
	defer span.Finish()

	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getTotalCount); err != nil {
		return nil, errors.Wrap(err, "statusRepo.GetStatus.GetContext.totalCount")
	}

	if totalCount == 0 {
		return &models.StatusList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
			Status:     make([]*models.Status, 0),
		}, nil
	}

	var statusList = make([]*models.Status, 0, pq.GetSize())
	rows, err := r.db.QueryxContext(ctx, getStatus, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "statusRepo.GetStatus.QueryxContext")
	}
	defer rows.Close()

	for rows.Next() {
		n := &models.Status{}
		if err = rows.StructScan(n); err != nil {
			return nil, errors.Wrap(err, "statusRepo.GetStatus.StructScan")
		}
		statusList = append(statusList, n)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "statusRepo.GetStatus.rows.Err")
	}

	return &models.StatusList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
		Status:     statusList,
	}, nil
}

// Find status by title
func (r *statusRepo) SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.StatusList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRepo.SearchByTitle")
	defer span.Finish()

	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, findByTitleCount, title); err != nil {
		return nil, errors.Wrap(err, "statusRepo.SearchByTitle.GetContext")
	}
	if totalCount == 0 {
		return &models.StatusList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Status:     make([]*models.Status, 0),
		}, nil
	}

	var statusList = make([]*models.Status, 0, query.GetSize())
	rows, err := r.db.QueryxContext(ctx, findByTitle, title, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "statusRepo.SearchByTitle.QueryxContext")
	}
	defer rows.Close()

	for rows.Next() {
		n := &models.Status{}
		if err = rows.StructScan(n); err != nil {
			return nil, errors.Wrap(err, "statusRepo.SearchByTitle.StructScan")
		}
		statusList = append(statusList, n)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "statusRepo.SearchByTitle.rows.Err")
	}

	return &models.StatusList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Status:     statusList,
	}, nil
}

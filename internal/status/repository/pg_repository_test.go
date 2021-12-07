package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"

	"github.com/mahfuz110244/api-mc/internal/models"
)

func TestStatusRepo_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	statusRepo := NewStatusRepository(sqlxDB)

	t.Run("Create", func(t *testing.T) {
		userUID := uuid.New()
		name := "estmate"
		description := "Description"
		created_by := userUID
		updated_by := userUID

		rows := sqlmock.NewRows([]string{"name", "description", "created_by", "updated_by"}).AddRow(name, description, created_by, updated_by)

		status := &models.Status{
			Name:        name,
			Description: description,
			CreatedBy:   created_by,
			UpdatedBy:   updated_by,
		}

		mock.ExpectQuery(createStatus).WithArgs(status.Name, status.Description).WillReturnRows(rows)

		createdStatus, err := statusRepo.Create(context.Background(), status)

		require.NoError(t, err)
		require.NotNil(t, createdStatus)
		require.Equal(t, status.Name, createdStatus.Name)
	})
}

func TestStatusRepo_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	statusRepo := NewStatusRepository(sqlxDB)

	t.Run("Update", func(t *testing.T) {
		statusUID := uuid.New()
		name := "estmate"
		description := "Description"

		rows := sqlmock.NewRows([]string{"id", "name", "description"}).AddRow(statusUID, name, description)

		status := &models.Status{
			ID:          statusUID,
			Name:        name,
			Description: description,
		}

		mock.ExpectQuery(updateStatus).WithArgs(status.Name,
			status.Description,
			status.ID,
		).WillReturnRows(rows)

		updatedStatus, err := statusRepo.Update(context.Background(), status)

		require.NoError(t, err)
		require.NotNil(t, updateStatus)
		require.Equal(t, updatedStatus, status)
	})
}

func TestStatusRepo_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	statusRepo := NewStatusRepository(sqlxDB)

	t.Run("Delete", func(t *testing.T) {
		statusUID := uuid.New()
		mock.ExpectExec(deleteStatus).WithArgs(statusUID).WillReturnResult(sqlmock.NewResult(1, 1))

		err := statusRepo.Delete(context.Background(), statusUID)

		require.NoError(t, err)
	})
}

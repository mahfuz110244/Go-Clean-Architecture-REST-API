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
		authorUID := uuid.New()
		title := "title"
		content := "content"

		rows := sqlmock.NewRows([]string{"author_id", "title", "content"}).AddRow(authorUID, title, content)

		status := &models.Status{
			AuthorID: authorUID,
			Title:    title,
			Content:  content,
		}

		mock.ExpectQuery(createStatus).WithArgs(status.AuthorID, status.Title, status.Content, status.Category).WillReturnRows(rows)

		createdStatus, err := statusRepo.Create(context.Background(), status)

		require.NoError(t, err)
		require.NotNil(t, createdStatus)
		require.Equal(t, status.Title, createdStatus.Title)
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
		title := "title"
		content := "content"

		rows := sqlmock.NewRows([]string{"status_id", "title", "content"}).AddRow(statusUID, title, content)

		status := &models.Status{
			StatusID: statusUID,
			Title:    title,
			Content:  content,
		}

		mock.ExpectQuery(updateStatus).WithArgs(status.Title,
			status.Content,
			status.ImageURL,
			status.Category,
			status.StatusID,
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

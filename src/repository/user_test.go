package repository

import (
	"context"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
	"gitlab.sima-land.ru/dev-dep/secretary-bot/models"
)

func TestCreateUserOfGitlab(t *testing.T) {
	mock, err := pgxmock.NewConn()

	t.Parallel()
	require.NoError(t, err)

	defer mock.Close(context.Background())

	user := models.User{
		GitLabUsername:   "test",
		GitlabID:         1,
		GitlabUpdatedAt:  time.Now(),
		GitlabStateID:    models.GitlabStateActive,
		GitlabWebsiteUrl: "gitlab.sima-land.ru",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.GitlabID, user.GitLabUsername, user.GitlabUpdatedAt, user.GitlabStateID, user.GitlabWebsiteUrl).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = CreateUserOfGitlab(mock, user)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

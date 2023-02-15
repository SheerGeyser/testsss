package repository

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.sima-land.ru/dev-dep/secretary-bot/models"
)

const createUserSQL = `
	INSERT INTO users (gitlab_id, gitlab_username, gitlab_updated_at, gitlab_state_id, gitlab_website_url)
	VALUES ($1, $2, $3, $4, $5)`

func CreateUserOfGitlab(db Db, user models.User) error {
	_, err := db.Exec(context.Background(), createUserSQL, user.GitlabID, user.GitLabUsername,
		user.GitlabUpdatedAt, user.GitlabStateID, user.GitlabWebsiteUrl)
	if err != nil {
		return errors.Wrap(err, "can not insert user")
	}

	return nil
}

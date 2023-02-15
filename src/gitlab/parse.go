package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"gitlab.sima-land.ru/dev-dep/secretary-bot/models"
)

func ReadGroupUsers() ([]byte, error) {
	url := "https://gitlab.sima-land.ru/api/v4/groups/" + os.Getenv("GITLAB_GROUP_ID") + "/members"
	token := "Bearer " + os.Getenv("GITLAB_TOKEN")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "can not create http request")
	}

	req.Header.Add("Authorization", token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "can not do http request")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "can not read response body")
	}
	defer resp.Body.Close()

	return body, nil
}

func ParseGroupUsers(parse []byte) ([]models.GitUser, error) {
	var members []models.GitUser

	err := json.Unmarshal(parse, &members)
	if err != nil {
		return []models.GitUser{}, fmt.Errorf("%w", err)
	}

	return append([]models.GitUser{}, members...), nil
}

func GroupUsers() ([]models.GitUser, error) {
	readUser, err := ReadGroupUsers()
	if err != nil {
		return nil, err
	}

	return ParseGroupUsers(readUser)
}

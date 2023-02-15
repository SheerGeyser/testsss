package gitlab

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.sima-land.ru/dev-dep/secretary-bot/models"
)

const responseFixture = `
[
  {
    "id": 15,
    "username": "tetyaev",
    "name": "Pavel Tetyaev",
    "state": "active",
    "web_url": "https://gitlab.sima-land.ru/tetyaev",
    "access_level": 50,
    "created_at": "2019-05-15T09:01:27.235Z",
    "expires_at": null,
    "membership_state": "active"
  },
  {
    "id": 26,
    "username": "Pastukhov_k",
    "name": "Kirill Pastukhov",
    "state": "active",
    "web_url": "https://gitlab.sima-land.ru/Pastukhov_k",
    "access_level": 50,
    "created_at": "2019-05-15T09:03:13.862Z",
    "expires_at": null,
    "membership_state": "active"
  }
]
`

func TestParseGroupUsers(t *testing.T) {
	require := require.New(t)
	t.Parallel()

	gitUsers, err := ParseGroupUsers([]byte(responseFixture))
	require.Equal(len(gitUsers), 2)
	require.NoError(err)

	require.Equal(models.GitUser{ID: 15, Username: "tetyaev"}, gitUsers[0])
	require.Equal(models.GitUser{ID: 26, Username: "Pastukhov_k"}, gitUsers[1])

	test, err := ParseGroupUsers([]byte("Ad"))
	require.Equal(test, []models.GitUser{})
	require.Error(err)
}

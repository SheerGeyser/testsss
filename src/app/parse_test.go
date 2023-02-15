package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTelegramUsername(t *testing.T) {
	require := require.New(t)
	t.Parallel()

	sima, err := parseTelegramUsername("https://sima-land.ru/user")
	require.Equal("", sima)
	require.Error(err)

	usr, err := parseTelegramUsername("user")
	require.Error(err)
	require.Equal("", usr)

	just, err := parseTelegramUsername("just_string")
	require.Error(err)
	require.Equal("", just)

	user, e := parseTelegramUsername("https://t.me/user")
	require.Equal(nil, e)
	require.Equal("user", user)

	ser, a := parseTelegramUsername("t.me/sergey")
	require.Equal(nil, a)
	require.Equal("sergey", ser)

	nik, err := parseTelegramUsername("t.me/Nik?xxxx")
	require.Equal(nil, err)
	require.Equal("Nik", nik)

	anna, q := parseTelegramUsername("http://t.me/anna")
	require.Equal(nil, q)
	require.Equal("anna", anna)
}

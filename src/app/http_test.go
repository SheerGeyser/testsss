package app

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		key   string
		value string
		want  int
	}{
		{
			name:  "Correct",
			key:   "X-Gitlab-Token",
			value: os.Getenv("GITLAB_WEBHOOK_TOKEN"),
			want:  200,
		},
		{
			name:  "Not correct value",
			key:   "X-Gitlab-Token",
			value: "asdasd",
			want:  403,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "localhost:8000", nil)

			req.Header.Add(tc.key, tc.value)

			handler(w, req)

			resp := w.Result()
			resp.Body.Close()
			assert.Equal(t, tc.want, resp.StatusCode)
		})
	}
}

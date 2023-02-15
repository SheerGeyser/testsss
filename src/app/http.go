package app

import (
	"log"
	http "net/http"
	"os"
)

func Server() {
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe("localhost:8000", nil); err != nil { //nolint: gosec
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	gitLabToken := r.Header.Get("X-Gitlab-Token")
	if gitLabToken == os.Getenv("GITLAB_WEBHOOK_TOKEN") {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

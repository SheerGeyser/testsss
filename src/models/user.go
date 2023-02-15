package models

import "time"

const (
	GitlabStateActive  = 1
	GitlabStateBlocked = 2
)

type User struct {
	ID               int
	GitlabID         int
	GitLabUsername   string
	GitlabStateID    int
	GitlabWebsiteUrl string
	GitlabUpdatedAt  time.Time
	TgChatID         int
	TgUsername       string
	IsSubscribed     bool
}

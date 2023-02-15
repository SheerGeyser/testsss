package main

import (
	"gitlab.sima-land.ru/dev-dep/secretary-bot/app"
)

func main() {
	go app.Server()
	app.TrackMessage()
}

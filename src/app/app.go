package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	Config Config
}

type Config struct {
	HttpAddress   string
	GitlabGroupID string
	GitlabToken   string
	DBConn        string
	TGBotToken    string
}

func NewApp(cfg Config) *App {
	cfg = Config{
		HttpAddress:   os.Getenv("GITLAB_TOKEN"),
		GitlabGroupID: os.Getenv("GITLAB_GROUP_ID"),
		GitlabToken:   os.Getenv("GITLAB_TOKEN"),
		DBConn:        os.Getenv("DB_CONN"),
		TGBotToken:    os.Getenv("TG_BOT_TOKEN"),
	}

	return &App{cfg}
}

func (a *App) Serve() {
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		Server()
	}()

	sig := <-cancelChan
	log.Printf("Stop %v", sig)
}

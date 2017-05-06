package main

import (
	"database/sql"
	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dsqlstate"
	"os"
	"time"
)

var (
	server *dsqlstate.Server
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	session, err := discordgo.New(os.Getenv("DG_TOKEN"))
	if err != nil {
		logrus.WithError(err).Fatal("Failed creating session")
	}
	// session.LogLevel = discordgo.LogInformational

	db, err := sql.Open("postgres", `dbname=dstate host=localhost user=postgres password=123 sslmode=disable`)
	if err != nil {
		logrus.WithError(err).Fatal("Failed opening db connection")
	}
	db.SetMaxOpenConns(25)

	server = dsqlstate.NewServer(db)
	server.LoadAllMembers = true
	server.RunWorkers(0)

	session.AddHandler(server.HandleEvent)
	session.Open()

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			printGuildCounts()
		}
	}
}

func printGuildCounts() {
	// n, err := state.JoinedGuildsCount()
	// logrus.Info("Guilds:", n, err)
}

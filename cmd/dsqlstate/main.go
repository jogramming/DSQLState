package main

import (
	"database/sql"
	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dsqlstate"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

var (
	server *dsqlstate.Server
)

func main() {
	logrus.Info("Starting...")
	logrus.SetLevel(logrus.DebugLevel)

	session, err := discordgo.New(os.Getenv("DG_TOKEN"))
	if err != nil {
		logrus.WithError(err).Fatal("Failed creating session")
	}
	session.State = discordgo.NewState()
	session.StateEnabled = false
	// session.LogLevel = discordgo.LogInformational

	db, err := sql.Open("postgres", `dbname=dstate host=localhost user=postgres password=123 sslmode=disable`)
	if err != nil {
		logrus.WithError(err).Fatal("Failed opening db connection")
		return
	}
	db.SetMaxOpenConns(25)

	server = dsqlstate.NewServer(db, 0)
	// server.Debug = true
	server.LoadAllMembers = true
	server.RunWorkers(0)

	session.AddHandler(server.HandleEvent)
	session.Open()

	ticker := time.NewTicker(time.Second)

	go http.ListenAndServe(":8080", nil)

	for {
		select {
		case <-ticker.C:
			printGuildCounts()
		}
	}
}

func printGuildCounts() {
	b, n := server.NumNotReady()
	logrus.Info("Shards ready: ", b, " guilds not ready: ", n, " GO: ", runtime.NumGoroutine())
	// n, err := state.JoinedGuildsCount()
	// logrus.Info("Guilds:", n, err)
}

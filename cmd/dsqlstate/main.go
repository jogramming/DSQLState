package main

import (
	"database/sql"
	"github.com/Sirupsen/logrus"
	"github.com/jonas747/discordgo"
	"github.com/jonas747/dsqlstate"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

var (
	server  *dsqlstate.Server
	doTrace = false
)

func main() {
	logrus.Info("Starting... v0.0.3")
	logrus.SetLevel(logrus.DebugLevel)

	if doTrace {
		traceOutput, err := os.Create("trace")
		if err != nil {
			logrus.WithError(err).Fatal("Failed creating trace file")
		}
		err = trace.Start(traceOutput)
		if err != nil {
			logrus.WithError(err).Fatal("Failed starting trace")
		}
		defer func() {
			trace.Stop()
			traceOutput.Close()
		}()
	}

	session, err := discordgo.New(os.Getenv("DG_TOKEN"))
	if err != nil {
		logrus.WithError(err).Fatal("Failed creating session")
	}
	session.State = discordgo.NewState()
	session.StateEnabled = false
	// session.LogLevel = discordgo.LogDebug

	db, err := sql.Open("postgres", `dbname=dstate host=localhost user=postgres password=123 sslmode=disable`)
	if err != nil {
		logrus.WithError(err).Fatal("Failed opening db connection")
		return
	}
	db.SetMaxOpenConns(10)

	server = dsqlstate.NewServer(db, 0)
	// server.Debug = true
	server.LoadAllMembers = true
	server.RunWorkers(0)

	session.AddHandler(server.HandleEvent)
	session.Open()

	ticker := time.NewTicker(time.Second)

	go http.ListenAndServe(":8080", nil)

	ticker2 := time.NewTicker(time.Millisecond * 100)
	started := time.Now()

	for {
		select {
		case <-ticker.C:
			printGuildCounts()
		case <-ticker2.C:
			if server.AllGuildsReady() {
				ticker2.Stop()
				logrus.Info("All ready! took: ", time.Since(started))
				// return
			}
		}
	}
}

func printGuildCounts() {
	b, n := server.NumNotReady()
	logrus.Info("Shards ready: ", b, " guilds not ready: ", n, " GO: ", runtime.NumGoroutine())
	// n, err := state.JoinedGuildsCount()
	// logrus.Info("Guilds:", n, err)
}

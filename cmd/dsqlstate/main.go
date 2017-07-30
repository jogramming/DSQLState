package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/jonas747/discordgo"
	"github.com/jonas747/dshardmanager"
	"github.com/jonas747/dsqlstate"
	_ "github.com/lib/pq"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	stdlog "log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

var (
	servers   []*dsqlstate.Server
	doTrace   = false
	db        *sql.DB
	numShards int

	FlagToken              string
	FlagDB                 string
	FlagHost               string
	FlagUser               string
	FlagPW                 string
	FlagSSLMode            string
	FlagMaxConnections     int
	FlagLogErrors          string
	FlagConnEventsChannel  string
	FlagShardStatusChannel string
	FlagBotName            string
)

func main() {
	stdlog.SetFlags(0)
	flag.StringVar(&FlagToken, "t", "", "The discord token to use, will also check the env variable DG_TOKEN")
	flag.StringVar(&FlagDB, "db", "dstate", "The database to use")
	flag.StringVar(&FlagHost, "host", "localhost", "The host to use when connecting to the db")
	flag.StringVar(&FlagUser, "user", "postgres", "The user to use when connecting to the db")
	flag.StringVar(&FlagPW, "pw", "123", "The password to use when connecting to the datbase")
	flag.StringVar(&FlagSSLMode, "sslmode", "disable", "The sslmode to use when connecting to the datbase")
	flag.IntVar(&FlagMaxConnections, "maxconn", 10, "Max number of connections to the database")
	flag.StringVar(&FlagLogErrors, "errors", "dsqlstate_errors.log", "Where to log errors, if empty, they will not be logger to disk")
	flag.StringVar(&FlagConnEventsChannel, "connevtchan", "", "Channel to log connections events to in discord, leave empty for none")
	flag.StringVar(&FlagShardStatusChannel, "shardstatuschan", "", "Channel to keep updated sharding status message in")
	flag.StringVar(&FlagBotName, "name", "dsqlstate", "Bot name to use")
	flag.Parse()

	logrus.Info("Starting... v" + dsqlstate.VersionString)
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

	if FlagLogErrors != "" {
		logrus.AddHook(lfshook.NewHook(lfshook.PathMap{
			// logrus.ErrorLevel: FlagLogErrors,
			logrus.InfoLevel: FlagLogErrors,
		}))
		logrus.Info("Added log hook")
	}
	stdlog.SetOutput(logrus.StandardLogger().Writer())

	if FlagToken == "" {
		FlagToken = os.Getenv("DG_TOKEN")
		if FlagToken == "" {
			logrus.Fatal("No discord token specified through -t or $DG_TOKEN")
			return
		}
	}

	sm, err := SetupShardManager()
	if err != nil {
		logrus.WithError(err).Fatal("Failed setting up shard manager")
		return
	}

	db, err = sql.Open("postgres", fmt.Sprintf(`dbname=%s host=%s user=%s password=%s sslmode=%s`, FlagDB, FlagHost, FlagUser, FlagPW, FlagSSLMode))
	if err != nil {
		logrus.WithError(err).Fatal("Failed opening db connection")
		return
	}

	// boil.DebugMode = true
	db.SetMaxOpenConns(10)
	logrus.Info("Connected to database")

	// Set up the db
	_, err = db.Exec(schema)
	if err != nil {
		logrus.WithError(err).Fatal("Failed setting up db tables")
		return
	}
	logrus.Info("Initilaized db schema")

	err = sm.Start()
	if err != nil {
		logrus.WithError(err).Fatal("Failed starting shard manager")
		return
	}

	ticker := time.NewTicker(time.Second * 5)

	go http.ListenAndServe(":8080", nil)

	// ticker2 := time.NewTicker(time.Millisecond * 100)
	// started := time.Now()

	for {
		select {
		case <-ticker.C:
			printGuildCounts()
			// case <-ticker2.C:
			// 	if server.AllGuildsReady() {
			// 		ticker2.Stop()
			// 		logrus.Info("All ready! took: ", time.Since(started))
			// 		// return
			// 	}
		}
	}
}

func printGuildCounts() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	notReady := 0
	numShard := len(servers)

	events := 0
	for _, v := range servers {
		notReady += v.NumNotReady()
		events += v.FlushEventCount()
	}
	logrus.Info("-----------------")
	logrus.Info("Shards: ", numShard, " guilds not ready: ", notReady, " GO: ", runtime.NumGoroutine(), ", Events: ", events, ", alloc(M): ", m.Alloc/1000000)
	logrus.Info("-----------------")
}

func SetupShardManager() (*dshardmanager.Manager, error) {
	sm := dshardmanager.New(FlagToken)

	sm.SessionFunc = smSessionFunc

	sm.LogChannel = FlagConnEventsChannel
	sm.StatusMessageChannel = FlagShardStatusChannel
	sm.Name = FlagBotName

	var err error
	numShards, err = sm.GetRecommendedCount()
	if err != nil {
		return sm, err
	}

	sm.GuildCountsFunc = func() []int {
		counts, err := dsqlstate.NumGuildsPerShard(db, numShards)
		if err != nil {
			logrus.WithError(err).Error("Failed counting guild shards")
		}

		return counts
	}

	sm.OnEvent = func(e *dshardmanager.Event) {
		if e.Type != dshardmanager.EventError {
			sm.LogConnectionEventStd(e)
			return
		}

		logrus.WithError(errors.New(e.Msg)).Error("Shard manager reported an error")
	}

	return sm, nil

}

type Evt struct {
	S   *discordgo.Session
	Evt interface{}
}

var (
	userCheckCache = dsqlstate.NewUserCheckCache()
)

func smSessionFunc(token string) (*discordgo.Session, error) {
	session, err := discordgo.New(token)
	if err != nil {
		return nil, err
	}
	session.SyncEvents = true
	session.StateEnabled = false
	server, err := dsqlstate.NewServer(session, db, userCheckCache)
	if err != nil {
		logrus.WithError(err).Fatal("Failed creating dsqlstate")
		return nil, err
	}

	server.LoadAllMembers = true
	go server.RunWorkers()

	evtChan := make(chan Evt, 100)

	session.AddHandler(func(s *discordgo.Session, evt interface{}) {
		if _, ok := evt.(*discordgo.Event); ok {
			// Do this check beforehand
			return
		}
		evtChan <- Evt{s, evt}
	})

	go func() {
		for {
			evt := <-evtChan

			err := server.HandleEvent(evt.S, evt.Evt)
			if err != nil {
				logrus.WithError(err).Error("DSQLState encountered an error")
			}
		}
	}()

	servers = append(servers, server)

	return session, nil
}

package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dshardmanager"
	"github.com/jonas747/dsqlstate"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
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

	FlagToken          string
	FlagDB             string
	FlagHost           string
	FlagUser           string
	FlagPW             string
	FlagSSLMode        string
	FlagMaxConnections int
	FlagLogErrors      string
)

func main() {
	flag.StringVar(&FlagToken, "t", "", "The discord token to use, will also check the env variable DG_TOKEN")
	flag.StringVar(&FlagDB, "db", "dstate", "The database to use")
	flag.StringVar(&FlagHost, "host", "localhost", "The host to use when connecting to the db")
	flag.StringVar(&FlagUser, "user", "postgres", "The user to use when connecting to the db")
	flag.StringVar(&FlagPW, "pw", "123", "The password to use when connecting to the datbase")
	flag.StringVar(&FlagSSLMode, "sslmode", "disable", "The sslmode to use when connecting to the datbase")
	flag.IntVar(&FlagMaxConnections, "maxconn", 10, "Max number of connections to the database")
	flag.StringVar(&FlagLogErrors, "errors", "dsqlstate_errors.log", "Where to log errors, if empty, they will not be logger to disk")
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
			logrus.ErrorLevel: FlagLogErrors,
		}))
		logrus.Info("Added log hook")
	}

	if FlagToken == "" {
		FlagToken = os.Getenv("DG_TOKEN")
		if FlagToken == "" {
			logrus.Fatal("No discord token specified through -t or $DG_TOKEN")
			return
		}
	}

	sm := dshardmanager.New(FlagToken, dshardmanager.OptSessionFunc(func(token string) (*discordgo.Session, error) {
		session, err := discordgo.New(token)
		if err != nil {
			return nil, err
		}
		session.StateEnabled = false
		return session, nil
	}))

	sm.OnEvent = func(e *dshardmanager.Event) {
		if e.Type != dshardmanager.EventError {
			return
		}

		logrus.WithError(errors.New(e.Msg)).Error("Shard manager reported an error")
	}

	db, err := sql.Open("postgres", fmt.Sprintf(`dbname=%s host=%s user=%s password=%s sslmode=%s`, FlagDB, FlagHost, FlagUser, FlagPW, FlagSSLMode))
	if err != nil {
		logrus.WithError(err).Fatal("Failed opening db connection")
		return
	}

	db.SetMaxOpenConns(10)
	logrus.Info("Connected to database")

	// Set up the db
	_, err = db.Exec(schema)
	if err != nil {
		logrus.WithError(err).Fatal("Failed setting up db tables")
		return
	}
	logrus.Info("Initilaized db schema")

	server, err = dsqlstate.NewServer(db, 0)
	if err != nil {
		logrus.WithError(err).Fatal("Failed creating dsqlstate")
		return
	}

	server.LoadAllMembers = true
	server.RunWorkers(0)

	sm.AddHandler(handleEvent)
	err = sm.Start()
	if err != nil {
		logrus.WithError(err).Fatal("Failed starting shard manager")
		return
	}

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
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	logrus.Info("Shards ready: ", b, " guilds not ready: ", n, " GO: ", runtime.NumGoroutine(), ", alloc: ", m.Alloc/1000000)
}

func handleEvent(session *discordgo.Session, evt interface{}) {
	if _, ok := evt.(*discordgo.Event); ok {
		// Do this check beforehand
		return
	}

	err := server.HandleEvent(session, evt)
	if err != nil {
		logrus.WithError(err).Error("DSQLState encounteredn an error")
	}
}

const schema = `
CREATE TABLE IF NOT EXISTS discord_users (
	id            bigint PRIMARY KEY,
	created_at    TIMESTAMP WITH TIME ZONE NOT NULL,
	
	username      varchar(32) NOT NULL,
	discriminator varchar(4) NOT NULL,
	bot           bool NOT NULL,
	avatar        text NOT NULL,

	status text NOT NULL,
	game_name text,
	game_type int,
	game_url text
);

CREATE INDEX ON discord_users(lower(username));

CREATE TABLE IF NOT EXISTS discord_guilds (
	id bigint PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	left_at TIMESTAMP WITH TIME ZONE,

	name               text NOT NULL,
	icon               text NOT NULL,
	region             text NOT NULL,
	afk_channel_id     bigint NOT NULL,
	embed_channel_id   bigint NOT NULL,
	owner_id           bigint NOT NULL,
	splash             text NOT NULL,
	afk_timeout        int NOT NULL,
	member_count       int NOT NULL,
	verification_level smallint NOT NULL,
	embed_enabled      bool NOT NULL,
	large               bool NOT NULL,
	default_message_notifications smallint NOT NULL
);

CREATE TABLE IF NOT EXISTS discord_guild_roles (
	id bigint PRIMARY KEY,
	guild_id bigint NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE,

	name text NOT NULL,
	managed bool NOT NULL,
	mentionable bool NOT NULL,
	hoist bool NOT NULL,
	color int NOT NULL,
	position int NOT NULL,
	permissions int NOT NULL
);

CREATE TABLE IF NOT EXISTS discord_channels (
	id bigint PRIMARY KEY,
	guild_id bigint,
	recipient_id bigint,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE,

	name text NOT NULL,
	topic text NOT NULL,
	type text NOT NULL,
	last_message_id bigint NOT NULL,
	position int NOT NULL,
	bitrate int NOT NULL
);

CREATE INDEX ON discord_channels(guild_id);
CREATE INDEX ON discord_channels(recipient_id);

CREATE TABLE IF NOT EXISTS discord_channel_overwrites (
	id bigint NOT NULL,
	channel_id bigint references discord_channels(id) NOT NULL,

	type varchar(10) NOT NULL,
	allow int NOT NULL,
	deny int NOT NULL,

	PRIMARY KEY(channel_id, id)
);

CREATE INDEX ON discord_channel_overwrites(channel_id);
CREATE INDEX ON discord_channel_overwrites(id);


CREATE TABLE IF NOT EXISTS discord_members (
	user_id bigint references discord_users(id) NOT NULL,
	guild_id bigint NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,

	left_at TIMESTAMP WITH TIME ZONE,

	joined_at TIMESTAMP WITH TIME ZONE NOT NULL,
	nick varchar(32) NOT NULL,
	deaf bool NOT NULL,
	mute bool NOT NULL,
	roles bigint[] NOT NULL,

	PRIMARY KEY(user_id, guild_id)
);

CREATE INDEX ON discord_members(user_id);
CREATE INDEX ON discord_members(guild_id);

CREATE TABLE IF NOT EXISTS discord_voice_states (
	user_id bigint NOT NULL,
	guild_id bigint,
	channel_id bigint references discord_channels(id) NOT NULL,
	session_id text NOT NULL,

	surpress bool NOT NULL,
	self_mute bool NOT NULL,
	self_deaf bool NOT NULL,
	mute bool NOT NULL,
	deaf bool NOT NULL,

	PRIMARY KEY(guild_id, user_id)
);

CREATE INDEX ON discord_voice_states(guild_id);
CREATE INDEX ON discord_voice_states(channel_id);

CREATE TABLE IF NOT EXISTS discord_messages (
	id bigint PRIMARY KEY,
	channel_id bigint NOT NULL,

	timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
	edited_timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE,

	-- sqlboiler has a hard time with nullable arrays
	mention_roles bigint[] NOT NULL,
	mentions bigint[] NOT NULL,
	mention_everyone bool NOT NULL,
	
	author_id bigint NOT NULL,
	author_username varchar(32) NOT NULL,
	author_discrim int NOT NULL,
	author_avatar text NOT NULL,
	author_bot bool NOT NULL,

	content text NOT NULL,
	embeds bigint[] NOT NULL
);

CREATE INDEX ON discord_messages(channel_id);

CREATE TABLE IF NOT EXISTS discord_message_revisions (
	revision_num int,
	message_id bigint references discord_messages(id) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,

	content text NOT NULL,
	embeds bigint[] NOT NULL,

	mentions bigint[] NOT NULL,
	mention_roles bigint[] NOT NULL,

	PRIMARY KEY(message_id, revision_num)
);

CREATE INDEX ON discord_message_revisions(message_id);

CREATE TABLE IF NOT EXISTS discord_message_embeds (
	id bigserial PRIMARY KEY,
	message_id bigint references discord_messages(id) NOT NULL,
	revision_num int NOT NULL,

	url text NOT NULL,
	type text NOT NULL,
	title text NOT NULL,
	description text NOT NULL,
	timestamp text NOT NULL,
	color int NOT NULL,

	field_names text[] NOT NULL,
	field_values text[] NOT NULL,
	field_inlines bool[] NOT NULL,

	footer_text text,
	footer_icon_url text,
	footer_proxy_icon_url text,

	image_url text,
	image_proxy_url text,
	image_width int,
	image_height int,

	thumbnail_url text,
	thumbnail_proxy_url text,
	thumbnail_width int,
	thumbnail_height int,

	video_url text,
	video_proxy_url text,
	video_width int,
	video_height int,

	provider_url text,
	provider_name text,

	author_url text,
	author_name text,
	author_icon_url text,
	author_proxy_icon_url text
);

CREATE INDEX ON discord_message_embeds(message_id);
`

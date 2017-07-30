package dsqlstate

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jonas747/discordgo"
	"github.com/jonas747/dsqlstate/models"
	// "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries/qm"
	"gopkg.in/nullbio/null.v6"
	"log"
	"strconv"
	"sync"
	"time"
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

type memCache struct {
	SelfUser  *discordgo.User
	SessionID string
	sync.Mutex
}

// TrackChangesSettings contains a bunch of toggles for what to send change events on
// change events are sent to a table you can query on a interval to then process.
// Toggles that are commented out are not yet implemented but planned
type TrackChangesSettings struct {
	Username       bool
	PresenceGame   bool
	PresenceURL    bool
	PresenceStatus bool

	ChannelName       bool
	ChannelTopic      bool
	ChannelPermission bool

	MemberNickname bool
	MemberAdded    bool
	MemberRoles    bool
	UserAdded      bool

	VoiceStateChangeChannel bool
	VoiceStateMuteDeaf      bool
}

// Server keeps the database up to date
type Server struct {
	Session          *discordgo.Session
	self             *discordgo.User
	db               *sql.DB
	Debug            bool
	OnError          func(err error)
	OnLog            func(msg string)
	LoadAllMembers   bool
	UpdateGameStatus bool

	LogChanges TrackChangesSettings

	guildsLock          sync.RWMutex
	guildsToBeProcessed map[int64]bool

	membersPullQueue       []string
	membersPullLock        sync.Mutex
	stopMembersWorker      chan bool
	membersChunkEvtHandled bool

	cache memCache
}

type QueuedEvent struct {
	Evt     interface{}
	Session *discordgo.Session
}

// New returns a default server using the database
func NewServer(s *discordgo.Session, db *sql.DB) (*Server, error) {

	// queue, err := NewEventQueue()
	// if err != nil {
	// 	return nil, err
	// }

	return &Server{
		Session: s,
		db:      db,
		// queue:               queue,
		LoadAllMembers:      true,
		membersPullQueue:    make([]string, 0),
		guildsToBeProcessed: make(map[int64]bool),
	}, nil
}

// RunWorkers starts the shard workers, this is required if you want all members loaded into the db
func (s *Server) RunWorkers() {
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-s.stopMembersWorker:
			return
		case <-t.C:
			s.NextMembersPull()
		}
	}
}

func (s *Server) NextMembersPull() {
	s.membersPullLock.Lock()
	if len(s.membersPullQueue) < 1 || !s.membersChunkEvtHandled {
		s.membersPullLock.Unlock()
		return
	}

	next := s.membersPullQueue[0]
	queueFiltered := make([]string, 0, len(s.membersPullQueue))
	for _, elem := range s.membersPullQueue {
		if elem != next {
			queueFiltered = append(queueFiltered, elem)
		}
	}
	s.membersPullQueue = queueFiltered
	s.membersPullLock.Unlock()

	s.Session.RequestGuildMembers(next, "", 0)
}

// StopWorkers stops the member fetcher workers
func (s *Server) StopWorkers() {
	close(s.stopMembersWorker)
}

func (s *Server) handleError(err error, message string) bool {
	if err == nil {
		return false
	}

	if s.OnError != nil {
		s.OnError(errors.Wrap(err, message))
	} else {
		log.Println("[DSQLSTATE] error: " + message + ": " + err.Error())
	}

	return true
}

func (s *Server) AllGuildsReadyNL() bool {
	if len(s.guildsToBeProcessed) > 0 {
		return false
	}

	return true
}

func (s *Server) AllGuildsReady() bool {
	s.guildsLock.RLock()
	defer s.guildsLock.RUnlock()
	return s.AllGuildsReadyNL()
}

func (s *Server) NumNotReady() int {
	s.guildsLock.RLock()
	defer s.guildsLock.RUnlock()
	return len(s.guildsToBeProcessed)
}

func (srv *Server) HandleEvent(s *discordgo.Session, evt interface{}) error {

	switch t := evt.(type) {
	case *discordgo.Ready:
		return srv.ready(s, t)
	}

	switch t := evt.(type) {
	// Guilds
	case *discordgo.GuildCreate:
		return srv.guildCreate(s, t.Guild)
	case *discordgo.GuildDelete:
		return srv.guildRemove(t.Guild)
	}

	return srv.handleNoSessionEvent(evt, true)
}

// BotID is a helper for retrieving the bot's id
func (srv *Server) BotID() string {
	srv.cache.Lock()
	defer srv.cache.Unlock()
	if srv.cache.SelfUser == nil {
		return ""
	}

	return srv.cache.SelfUser.ID
}

func (srv *Server) handleNoSessionEvent(evt interface{}, retry bool) error {
	var err error
	switch t := evt.(type) {
	case *discordgo.GuildUpdate:
		err = srv.guildUpdate(t.Guild)

	// Members
	case *discordgo.GuildMemberAdd:
		err = srv.updateMember(srv.db, t.Member, true)
	case *discordgo.GuildMemberUpdate:
		err = srv.updateMember(srv.db, t.Member, true)
	case *discordgo.GuildMemberRemove:
		err = models.DMembers(srv.db, qm.Where("user_id = ?", t.User.ID), qm.Where("guild_id = ?", t.GuildID)).UpdateAll(models.M{"left_at": time.Now()})
	// Roles
	case *discordgo.GuildRoleCreate:
		err = srv.updateRole(t.GuildID, t.Role)
	case *discordgo.GuildRoleUpdate:
		err = srv.updateRole(t.GuildID, t.Role)
	case *discordgo.GuildRoleDelete:
		err = srv.removeRole(t.RoleID)

	// Channels
	case *discordgo.ChannelCreate:
		if t.Channel.GuildID != "" {
			err = srv.updateGuildChannel(t.Channel)
		} else if t.Channel.Type == discordgo.ChannelTypeDM || t.Channel.Type == discordgo.ChannelTypeGroupDM {
			err = srv.updatePrivateChannel(t.Channel)
		}
	case *discordgo.ChannelUpdate:
		if t.Channel.GuildID != "" {
			err = srv.updateGuildChannel(t.Channel)
		} else if t.Channel.Type == discordgo.ChannelTypeDM || t.Channel.Type == discordgo.ChannelTypeGroupDM {
			err = srv.updatePrivateChannel(t.Channel)
		}
	case *discordgo.ChannelDelete:
		if t.Channel.GuildID != "" {
			models.DChannels(srv.db, qm.Where("id = ?", t.Channel.ID)).UpdateAll(models.M{"deleted_at": time.Now()})
		}
	// Messages
	case *discordgo.MessageCreate:
		err = srv.messageCreate(t.Message)
	case *discordgo.MessageUpdate:
		err = srv.messageUpdate(nil, nil, t.Message, retry)
	case *discordgo.MessageDelete:
		err = srv.messageDelete(t.Message)

	// Other
	case *discordgo.VoiceStateUpdate:
		err = srv.updateVoiecState(t.VoiceState)
	case *discordgo.UserUpdate:
		err = srv.updateUser(nil, t.User)
	case *discordgo.PresenceUpdate:
		err = srv.presenceUpdate(srv.db, &t.Presence)
	case *discordgo.GuildMembersChunk:
		err = srv.guildMembersChunk(t)
	}

	return err
}

func shardClause(guildColumn string, numShards, current int) string {
	// (guild_id >> 22) % num_shards == shard_id
	q := fmt.Sprintf("(%s >> 22) %% %d = %d", guildColumn, numShards, current)
	return q
}

func shardClauseAnd(guildColumn string, numShards, current int) string {
	if numShards < 2 {
		return ""
	}
	return " AND " + shardClause(guildColumn, numShards, current)
}

func (srv *Server) ready(s *discordgo.Session, r *discordgo.Ready) error {
	var now = time.Now()

	// Mark all guilds on this shard as deleted
	_, err := srv.db.Exec("UPDATE d_guilds SET left_at = $1 WHERE left_at IS NULL"+shardClauseAnd("id", s.ShardCount, s.ShardID), now)
	srv.handleError(err, "Failed marking shard guilds as left")

	sc := shardClauseAnd("guild_id", s.ShardCount, s.ShardID)

	// Mark all guild roles as deleted
	_, err = srv.db.Exec("UPDATE d_guild_roles SET deleted_at = $1 WHERE deleted_at IS NULL"+sc, now)
	srv.handleError(err, "Failed marking shard guild roles as deleted")

	// Mark all guild channels as deleted
	_, err = srv.db.Exec("UPDATE d_channels SET deleted_at = $1 WHERE deleted_at IS NULL"+sc, now)
	srv.handleError(err, "Failed marking shard guild channels as deleted")

	vcClause := shardClause("guild_id", s.ShardCount, s.ShardID)
	if vcClause != "" {
		vcClause = " WHERE " + vcClause
	}
	// Clear the voice srvates, as we get a new fresh set in the guild creates
	_, err = srv.db.Exec("DELETE FROM d_voice_states" + vcClause)
	srv.handleError(err, "Failed marking shard guild voice_states as deleted")

	// Clear members, as people can have left in the meantime, it is now unclear who is srvill on the server
	_, err = srv.db.Exec("UPDATE d_members SET left_at = $1 WHERE left_at IS NULL"+sc, now)
	srv.handleError(err, "Failed marking shard guild members as left")

	srv.guildsLock.Lock()

	doneChan := make(chan error)
	count := 0
	for _, v := range r.Guilds {
		parsedGID, err := strconv.ParseInt(v.ID, 10, 64)
		if err != nil {
			srv.guildsLock.Unlock()
			return errors.WithMessage(err, "ready, ParseGuildID, id: '"+v.ID+"'")
		}

		if v.Unavailable {
			srv.guildsToBeProcessed[parsedGID] = false
			srv.db.Exec("UPDATE d_guilds SET left_at = NULL WHERE id = $1", v.ID)
		} else {
			count++
			go func(g *discordgo.Guild) {
				tErr := srv.guildCreate(s, g)
				doneChan <- tErr
			}(v)
		}
	}
	srv.guildsLock.Unlock()

	for i := 0; i < count; i++ {
		tErr := <-doneChan
		if tErr != nil {
			err = tErr
		}
	}

	return errors.WithMessage(err, "ready, guildCreate")
}

func (srv *Server) guildMembersChunk(chunk *discordgo.GuildMembersChunk) error {
	for _, v := range chunk.Members {
		v.GuildID = chunk.GuildID
		err := srv.updateMember(srv.db, v, true)
		if err != nil {
			return errors.WithMessage(err, "guildMembersChunk")
		}
	}

	srv.membersPullLock.Lock()
	srv.membersChunkEvtHandled = true
	srv.membersPullLock.Unlock()

	return nil
}

func (srv *Server) guildCreate(session *discordgo.Session, g *discordgo.Guild) error {
	parsedGID, _ := strconv.ParseInt(g.ID, 10, 64)
	defer func() {
		srv.guildsLock.Lock()
		delete(srv.guildsToBeProcessed, parsedGID)
		srv.guildsLock.Unlock()

		// if !srv.removedFK {
		// 	return
		// }

		// if srv.AllGuildsReadyNL() {
		// 	// Re-instantiate the foreign key
		// 	_, err := srv.db.Exec("ALTER TABLE d_members ADD FOREIGN KEY(user_id) REFERENCES d_users(id)")
		// 	if !srv.handleError(err, "Failed adding back foreign key") {
		// 		srv.removedFK = false
		// 	}
		// }
	}()

	if srv.LoadAllMembers && g.Large {
		go func(gid string) {
			srv.membersPullLock.Lock()
			srv.membersPullQueue = append(srv.membersPullQueue, gid)
			srv.membersPullLock.Unlock()
		}(g.ID)
	}

	log.Println("GC! ", g.Name, " ID: ", g.ID, "PID: ", parsedGID)
	err := srv.guildUpdate(g)
	if err != nil {
		return errors.WithMessage(err, "GuildCreate")
	}

	// Update all roles
	for _, v := range g.Roles {
		err = srv.updateRole(g.ID, v)
		if err != nil {
			return errors.WithMessage(err, "GuildCreate")
		}
	}

	// Update all channels
	for _, v := range g.Channels {
		err = srv.updateGuildChannel(v)
		if err != nil {
			return errors.WithMessage(err, "GuildCreate")
		}
	}

	// Create the update list, and dont load users in twice on the initial startup
	// Thsi has to be carefully done to avoid deadlocks
	// toUpdatePresences := make(map[string]*discordgo.Presence)
	// srv.readyLock.Lock()

	// if srv.removedFK {
	// 	for _, v := range g.Presences {
	// 		if _, ok := srv.loadedUsers[v.User.ID]; !ok {
	// 			toUpdatePresences[v.User.ID] = v
	// 			srv.loadedUsers[v.User.ID] = true
	// 		}
	// 	}
	// 	for _, v := range g.Members {
	// 		if p, ok := toUpdatePresences[v.User.ID]; ok {
	// 			p.User = v.User
	// 		} else {
	// 			if _, ok := srv.loadedUsers[v.User.ID]; !ok {
	// 				toUpdatePresences[v.User.ID] = &discordgo.Presence{User: v.User}
	// 				srv.loadedUsers[v.User.ID] = true
	// 			}
	// 		}
	// 	}
	// } else {
	// 	for _, v := range g.Presences {
	// 		toUpdatePresences[v.User.ID] = v
	// 	}

	// 	for _, v := range g.Members {
	// 		if p, ok := toUpdatePresences[v.User.ID]; ok {
	// 			p.User = v.User
	// 		} else {

	// 			toUpdatePresences[v.User.ID] = &discordgo.Presence{User: v.User}
	// 		}
	// 	}
	// }
	// srv.readyLock.Unlock()

	// tx, err := srv.db.Begin()
	// if err != nil {
	// 	return errors.WithMessage(err, "guildCreate, tx begin")
	// }

	// for _, v := range toUpdatePresences {
	// 	// These savepoints made postgres run out of shared memory
	// 	// tx.Exec("SAVEPOINT sp")
	// 	err = srv.presenceUpdate(tx, v)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return errors.WithMessage(err, "guildCreate")
	// 	}
	// }

	// Update all the members and users
	for _, v := range g.Members {
		// tx.Exec("SAVEPOINT sp")
		err = srv.updateMember(srv.db, v, true)
		if err != nil {
			// tx.Rollback()
			return errors.WithMessage(err, "guildCreate")
		}
	}

	for _, v := range g.VoiceStates {
		err = srv.updateVoiecState(v)
		if err != nil {
			return errors.WithMessage(err, "guildCreate")
		}
	}

	return nil
}

func (srv *Server) guildRemove(g *discordgo.Guild) error {
	parsedID, _ := strconv.ParseInt(g.ID, 10, 64)

	srv.guildsLock.Lock()
	delete(srv.guildsToBeProcessed, parsedID)
	srv.guildsLock.Unlock()

	srv.db.Exec("UPDATE d_guilds SET left_at = $1 WHERE id = $2", time.Now(), g.ID)
	models.DVoiceStates(srv.db, qm.Where("guild_id = ?", g.ID)).DeleteAll()
	models.DMembers(srv.db, qm.Where("id = ?", g.ID)).DeleteAll()
	models.DGuildRoles(srv.db, qm.Where("guild_id = ?", g.ID)).DeleteAll()
	models.DChannels(srv.db, qm.Where("guild_id = ?", g.ID)).DeleteAll()

	return nil
}

func (srv *Server) guildUpdate(g *discordgo.Guild) error {
	parsedId, err := strconv.ParseInt(g.ID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "GuildUpdate, ParseID")
	}

	// Servers can have no owner in edge cases...
	ownerID, _ := strconv.ParseInt(g.OwnerID, 10, 64)

	var parsedAFK int64
	var embedChannel int64

	if g.AfkChannelID != "" {
		parsedAFK, err = strconv.ParseInt(g.AfkChannelID, 10, 64)
		if err != nil {
			return errors.WithMessage(err, "GuildUpdate, ParseAFK")
		}
	}

	if g.EmbedChannelID != "" {
		embedChannel, err = strconv.ParseInt(g.EmbedChannelID, 10, 64)
		if err != nil {
			return errors.WithMessage(err, "GuildUpdate, ParseEmbedChannel")
		}
	}

	model := &models.DGuild{
		ID: parsedId,

		OwnerID: ownerID,

		Name:   g.Name,
		Icon:   g.Icon,
		Region: g.Region,

		EmbedEnabled:   g.EmbedEnabled,
		EmbedChannelID: embedChannel,
		AfkChannelID:   parsedAFK,
		AfkTimeout:     g.AfkTimeout,

		Splash:            g.Splash,
		MemberCount:       g.MemberCount,
		VerificationLevel: int16(g.VerificationLevel),
		Large:             g.Large,
		DefaultMessageNotifications: int16(g.DefaultMessageNotifications),
	}

	err = model.Upsert(srv.db, true, []string{"id"}, []string{"name", "icon", "region", "afk_channel_id", "embed_channel_id", "owner_id", "splash", "afk_timeout", "member_count", "verification_level", "embed_enabled", "large", "default_message_notifications", "left_at"})
	if err != nil {
		return errors.WithMessage(err, "GuildUpdate")
	}
	return nil
}

func (s *Server) updateUser(exec boil.Executor, user *discordgo.User) error {
	if exec == nil {
		exec = s.db
	}

	parsedId, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "updateUser, ParseID")
	}

	model := &models.DUser{
		ID:            parsedId,
		Username:      user.Username,
		Discriminator: user.Discriminator,
		Bot:           user.Bot,
		Avatar:        user.Avatar,
	}

	toUpdate := []string{}
	if user.Username != "" {
		toUpdate = []string{"username", "discriminator", "avatar"}
	}

	err = model.Upsert(exec, true, []string{"id"}, toUpdate)
	return errors.WithMessage(err, "updateUser")
}

func (s *Server) presenceUpdate(exec boil.Executor, p *discordgo.Presence) error {

	parsedId, err := strconv.ParseInt(p.User.ID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "presenceUpdate, ParseID")
	}

	model := &models.DUser{
		ID:     parsedId,
		Status: string(p.Status),
	}

	columns := make([]string, 0, 12)
	// These columns, 90% sure they are available at all times
	columns = append(columns, "game_name", "game_type", "game_url")

	if p.User.Username != "" {
		// Update the user columns if they are here
		model.Username = p.User.Username
		model.Discriminator = p.User.Discriminator
		model.Avatar = p.User.Avatar
		model.Bot = p.User.Bot

		columns = append(columns, "username", "discriminator", "bot", "avatar")
	}

	if p.Status != "" {
		columns = append(columns, "status")
	}

	if p.Game != nil {
		model.GameName = null.StringFrom(p.Game.Name)
		model.GameType = null.IntFrom(p.Game.Type)
		model.GameURL = null.StringFrom(p.Game.URL)
	}

	err = model.Upsert(exec, true, []string{"id"}, columns)
	return errors.WithMessage(err, "presenceUpdate")
}

func (s *Server) updateMember(exec boil.Executor, member *discordgo.Member, updtUser bool) error {
	// Update roles
	if updtUser {
		err := s.updateUser(exec, member.User)
		if err != nil {
			return errors.WithMessage(err, "updateMember")
		}
	}

	parsedMID, err := strconv.ParseInt(member.User.ID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "updateMember, ParseID")
	}

	parsedGID, err := strconv.ParseInt(member.GuildID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "updateMember, ParseGuildID")
	}

	joinedParsed, _ := discordgo.Timestamp(member.JoinedAt).Parse()
	model := &models.DMember{
		UserID:  parsedMID,
		GuildID: parsedGID,

		JoinedAt: joinedParsed,
		Nick:     member.Nick,
		Deaf:     member.Deaf,
		Mute:     member.Mute,
		Roles:    make([]int64, 0, len(member.Roles)),
	}

	for _, r := range member.Roles {
		parsed, _ := strconv.ParseInt(r, 10, 64)
		model.Roles = append(model.Roles, parsed)
	}

	err = model.Upsert(exec, true, []string{"user_id", "guild_id"}, []string{"left_at", "joined_at", "nick", "deaf", "mute", "roles"})
	return errors.WithMessage(err, "updateMember")
}

func (s *Server) updateRole(guildID string, role *discordgo.Role) error {
	parsedGuildID, err := strconv.ParseInt(guildID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "updateRole, ParseGuildID")
	}

	roleIdParsed, err := strconv.ParseInt(role.ID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "updateRole, ParseID")
	}

	model := &models.DGuildRole{
		ID:      roleIdParsed,
		GuildID: parsedGuildID,

		Name:        role.Name,
		Managed:     role.Managed,
		Mentionable: role.Mentionable,
		Hoist:       role.Hoist,
		Color:       role.Color,
		Position:    role.Position,
		Permissions: role.Permissions,
	}

	err = model.Upsert(s.db, true, []string{"id"}, []string{"name", "mentionable", "hoist", "color", "position", "permissions"})
	return errors.WithMessage(err, "updateRole")
}

func (s *Server) removeRole(roleID string) error {
	_, err := s.db.Exec("UPDATE d_guild_roles SET deleted_at = $2 WHERE id = $1", roleID, time.Now())
	return errors.WithMessage(err, "removeRole")
}

func (s *Server) updateGuildChannel(channel *discordgo.Channel) error {
	parsedChannelId, err := strconv.ParseInt(channel.ID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "updateGuildChannel, ParseChannelId")
	}

	parsedGuildID, err := strconv.ParseInt(channel.GuildID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "updateGuildChannel, ParseGuildId")
	}

	var lastMessageID int64
	if channel.LastMessageID != "" {
		lastMessageID, _ = strconv.ParseInt(channel.LastMessageID, 10, 64)
	}

	typ := "text"
	if channel.Type == discordgo.ChannelTypeGuildVoice {
		typ = "voice"
	} else if channel.Type == discordgo.ChannelTypeGuildCategory {
		typ = "guildcategory"
	} else if channel.Type == discordgo.ChannelTypeDM {
		typ = "dm"
	} else if channel.Type == discordgo.ChannelTypeGroupDM {
		typ = "gdm"
	}

	model := &models.DChannel{
		ID:      parsedChannelId,
		GuildID: null.Int64From(parsedGuildID),

		Name:          channel.Name,
		Topic:         channel.Topic,
		Type:          typ,
		LastMessageID: lastMessageID,
		Position:      channel.Position,
		Bitrate:       channel.Bitrate,
	}

	// Update permission overwrites
	transaction, err := s.db.Begin()
	if err != nil {
		return errors.WithMessage(err, "updateGuildChannel, Begin")
	}

	err = model.Upsert(transaction, true, []string{"id"}, []string{"name", "topic", "last_message_id", "position", "bitrate"})
	if err != nil {
		transaction.Rollback()
		return errors.WithMessage(err, "updateGuildChannel")
	}

	args := []interface{}{parsedChannelId}
	query := "DELETE FROM d_channel_overwrites WHERE channel_id = $1"
	if len(channel.PermissionOverwrites) > 0 {
		query += " AND id NOT IN ("
		for i, v := range channel.PermissionOverwrites {
			if i != 0 {
				query += ","
			}
			query += "$" + strconv.Itoa(i+2)
			args = append(args, v.ID)
		}
		query += ")"
	}

	_, err = transaction.Exec(query, args...)
	if err != nil {
		transaction.Rollback()
		return errors.WithMessage(err, "updateGuildChannel, clear PermOverwrites")
	}

	for _, v := range channel.PermissionOverwrites {
		parsedID, err := strconv.ParseInt(v.ID, 10, 64)
		if s.handleError(err, "Failed parsing channel persmission overwrite id") {
			continue
		}

		model := models.DChannelOverwrite{
			ID:        parsedID,
			ChannelID: parsedChannelId,
			Type:      v.Type,
			Allow:     v.Allow,
			Deny:      v.Deny,
		}

		err = model.Upsert(transaction, true, []string{"channel_id", "id"}, []string{"allow", "deny"})
		if err != nil {
			transaction.Rollback()
			return errors.WithMessage(err, "updateGuildChannel, permOverwrite upsert")
		}
	}

	err = transaction.Commit()
	return errors.WithMessage(err, "updateGuildChannel")
}

func (s *Server) updatePrivateChannel(channel *discordgo.Channel) error {
	var recipient *discordgo.User
	for _, v := range channel.Recipients {
		if v.ID == s.self.ID {
			continue
		}

		recipient = v
		break
	}

	if recipient == nil {
		return errors.New("updatePrivateChannel: No recipient")
	}

	for _, v := range channel.Recipients {
		err := s.updateUser(nil, v)
		if err != nil {
			return errors.WithMessage(err, "updatePrivateChannel")
		}
	}

	parsedChannelId, err := strconv.ParseInt(channel.ID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "updatePrivateChannel, ParseChannelID")
	}

	parsedRecipient, err := strconv.ParseInt(recipient.ID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "updatePrivateChannel, ParseRecipientID")
	}

	var lastMessageID int64
	if channel.LastMessageID != "" {
		lastMessageID, _ = strconv.ParseInt(channel.LastMessageID, 10, 64)
	}

	model := &models.DChannel{
		ID:          parsedChannelId,
		RecipientID: null.Int64From(parsedRecipient),

		Name:          channel.Name,
		Topic:         channel.Topic,
		LastMessageID: lastMessageID,
	}

	err = model.Upsert(s.db, true, []string{"id"}, []string{"name", "topic", "last_message_id"})
	return errors.WithMessage(err, "updatePrivateChannel")
}

func (s *Server) updateVoiecState(vc *discordgo.VoiceState) error {
	parsedUser, err := strconv.ParseInt(vc.UserID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "updateVoiecState, ParseUserID")
	}

	// Groups
	parsedGuildID, _ := strconv.ParseInt(vc.GuildID, 10, 64)

	// Check if they left voice
	if vc.ChannelID == "" {
		query := "DELETE FROM d_voice_States WHERE user_id = $1"
		args := []interface{}{parsedUser}
		if parsedGuildID != 0 {
			query += " AND guild_id = $2"
			args = append(args, parsedGuildID)
		}

		_, err := s.db.Exec(query, args...)
		if err != nil {
			return errors.WithMessage(err, "updateVoiceState, delete")
		}
		return nil
	}

	parsedChannelID, err := strconv.ParseInt(vc.ChannelID, 10, 64)
	if err != nil {
		return errors.New("updateVoiceState, ParseChannleID")
	}

	model := &models.DVoiceState{
		UserID:    parsedUser,
		ChannelID: parsedChannelID,
		GuildID:   parsedGuildID,

		Surpress: vc.Suppress,
		SelfMute: vc.SelfMute,
		SelfDeaf: vc.SelfDeaf,
		Mute:     vc.Mute,
		Deaf:     vc.Deaf,
	}

	err = model.Upsert(s.db, true, []string{"guild_id", "user_id"}, []string{"surpress", "self_mute", "self_deaf", "mute", "deaf"})
	return errors.WithMessage(err, "updateVoiceState")
}

func (s *Server) messageCreate(m *discordgo.Message) error {
	parsedMID, err := strconv.ParseInt(m.ID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "messageCreate, ParseID")
	}

	parsedCID, err := strconv.ParseInt(m.ChannelID, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "messageCreate, ParseChannelID")
	}

	parsedTimeStamp, _ := m.Timestamp.Parse()

	transaction, err := s.db.Begin()
	if err != nil {
		return errors.WithMessage(err, "messageCreate, tx begin")
	}

	parsedAuthorID, _ := strconv.ParseInt(m.Author.ID, 10, 64)
	parsedAuthorDiscrim, _ := strconv.ParseInt(m.Author.Discriminator, 10, 32)

	model := &models.DMessage{
		ID:        parsedMID,
		ChannelID: parsedCID,
		Timestamp: parsedTimeStamp,

		AuthorID:       parsedAuthorID,
		AuthorUsername: m.Author.Username,
		AuthorDiscrim:  int(parsedAuthorDiscrim),
		AuthorAvatar:   m.Author.Avatar,
		AuthorBot:      m.Author.Bot,

		Mentions:        []int64{},
		MentionRoles:    []int64{},
		MentionEveryone: m.MentionEveryone,

		Content: m.Content,
		Embeds:  []int64{},
	}

	err = model.Insert(transaction)
	if err != nil {
		transaction.Rollback()
		return errors.WithMessage(err, "mesageCreate, insert")
	}

	return errors.WithMessage(s.messageUpdate(transaction, model, m, false), "messageCreate")
}

// Somewhat complicated update procedure:
// 1. Retry in 2 seconds if we got the update before the create
// 2. Lock the message for update
// 3. Create the revision model
// 4. create the embeds models
// 5. update the revision model
// 6. update the message model
// 7. Commit if all went well
//
// I need to do this in a move efficient way
func (s *Server) messageUpdate(transaction *sql.Tx, messageModel *models.DMessage, m *discordgo.Message, retry bool) error {
	parsedMID, _ := strconv.ParseInt(m.ID, 10, 64)

	if transaction == nil {
		var err error
		transaction, err = s.db.Begin()
		if err != nil {
			return errors.WithMessage(err, "messageUpdate, tx begin")
		}

		messageModel, err = models.DMessages(transaction, qm.Where("id = ?", parsedMID), qm.For("UPDATE")).One()
		if err == sql.ErrNoRows && retry {
			// Try again in a second in case of a fast embed update, or something like that
			transaction.Rollback()
			time.Sleep(time.Second)
			return s.messageUpdate(nil, nil, m, false)
		}

		if err != nil {
			transaction.Rollback()
			return errors.WithMessage(err, "messageUpdate, fetch message")
		}
	}

	num, err := models.DMessageRevisions(transaction, qm.Where("message_id = ?", parsedMID)).Count()
	if err != nil {
		transaction.Rollback()
		return errors.WithMessage(err, "messageUpdate, count revisions")
	}

	revisionModel := &models.DMessageRevision{
		MessageID:    parsedMID,
		RevisionNum:  int(num),
		Content:      m.Content,
		Embeds:       []int64{},
		Mentions:     []int64{},
		MentionRoles: []int64{},
	}

	for _, u := range m.Mentions {
		parsed, err := strconv.ParseInt(u.ID, 10, 64)
		if err == nil {
			revisionModel.Mentions = append(revisionModel.Mentions, parsed)
		}
	}
	for _, r := range m.MentionRoles {
		parsed, err := strconv.ParseInt(r, 10, 64)
		if err == nil {
			revisionModel.MentionRoles = append(revisionModel.MentionRoles, parsed)
		}
	}
	messageModel.Mentions = revisionModel.Mentions
	messageModel.MentionRoles = revisionModel.MentionRoles

	err = revisionModel.Insert(transaction)
	if err != nil {
		transaction.Rollback()
		return errors.WithMessage(err, "updateMessage, insert revision")
	}

	embedIds := make([]int64, 0)

	messageModel.Content = m.Content
	for _, v := range m.Embeds {
		embedmodel := createEmbedModel(v)
		embedmodel.MessageID = parsedMID
		embedmodel.RevisionNum = int(num)
		err = embedmodel.Insert(transaction)
		if err != nil {
			transaction.Rollback()
			return errors.WithMessage(err, "updateMessage, embedModel")
		}
		embedIds = append(embedIds, embedmodel.ID)
	}

	revisionModel.Embeds = embedIds
	err = revisionModel.Update(transaction, "embeds")
	if err != nil {
		transaction.Rollback()
		return errors.WithMessage(err, "updateMessage, revision update")
	}

	messageModel.Embeds = embedIds
	parsedEdited, _ := m.EditedTimestamp.Parse()
	messageModel.EditedTimestamp = parsedEdited

	err = messageModel.Update(transaction, "content", "embeds", "edited_timestamp")
	if err != nil {
		transaction.Rollback()
		return errors.WithMessage(err, "updateMessage, model update")
	}

	return errors.WithMessage(transaction.Commit(), "updateMessage, commit")
}

func createEmbedModel(embed *discordgo.MessageEmbed) *models.DMessageEmbed {
	// And here the long ass journy of creating an embed starts
	model := &models.DMessageEmbed{
		URL:         embed.URL,
		Type:        embed.Type,
		Title:       embed.Title,
		Description: embed.Description,
		Timestamp:   embed.Timestamp,
		Color:       embed.Color,
	}

	model.FieldNames = make([]string, len(embed.Fields))
	model.FieldValues = make([]string, len(embed.Fields))
	model.FieldInlines = make([]bool, len(embed.Fields))
	if len(embed.Fields) > 0 {
		for i := 0; i < len(embed.Fields); i++ {
			f := embed.Fields[i]
			model.FieldNames[i] = f.Name
			model.FieldValues[i] = f.Value
			model.FieldInlines[i] = f.Inline
		}
	}

	if embed.Footer != nil {
		model.FooterText = null.StringFrom(embed.Footer.Text)
		model.FooterIconURL = null.StringFrom(embed.Footer.IconURL)
		model.FooterProxyIconURL = null.StringFrom(embed.Footer.ProxyIconURL)
	}

	if embed.Thumbnail != nil {
		model.ThumbnailURL = null.StringFrom(embed.Thumbnail.URL)
		model.ThumbnailProxyURL = null.StringFrom(embed.Thumbnail.ProxyURL)
		model.ThumbnailWidth = null.IntFrom(embed.Thumbnail.Width)
		model.ThumbnailHeight = null.IntFrom(embed.Thumbnail.Height)
	}

	if embed.Image != nil {
		model.ImageURL = null.StringFrom(embed.Image.URL)
		model.ImageProxyURL = null.StringFrom(embed.Image.ProxyURL)
		model.ImageHeight = null.IntFrom(embed.Image.Height)
		model.ImageWidth = null.IntFrom(embed.Image.Width)
	}

	if embed.Video != nil {
		model.VideoURL = null.StringFrom(embed.Video.URL)
		model.VideoProxyURL = null.StringFrom(embed.Video.ProxyURL)
		model.VideoWidth = null.IntFrom(embed.Video.Width)
		model.VideoHeight = null.IntFrom(embed.Video.Height)
	}

	if embed.Provider != nil {
		model.ProviderName = null.StringFrom(embed.Provider.Name)
		model.ProviderURL = null.StringFrom(embed.Provider.URL)
	}

	if embed.Author != nil {
		model.AuthorURL = null.StringFrom(embed.Author.URL)
		model.AuthorName = null.StringFrom(embed.Author.Name)
		model.AuthorIconURL = null.StringFrom(embed.Author.IconURL)
		model.AuthorProxyIconURL = null.StringFrom(embed.Author.ProxyIconURL)
	}

	return model
}

func (s *Server) messageDelete(m *discordgo.Message) error {
	err := models.DMessages(s.db, qm.Where("id = ?", m.ID)).UpdateAll(models.M{"deleted_at": time.Now()})
	return errors.WithMessage(err, "messageDelete")
}

func (s *Server) SetMeta(name string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return errors.WithMessage(err, "SetMeta")
	}
	model := &models.DMetum{
		Key:   name,
		Value: b,
	}

	err = model.Upsert(s.db, true, []string{"key"}, []string{"value"})
	return errors.WithMessage(err, "SetMeta")
}

func (s *Server) QueueLength() uint64 {
	return 0
	// return s.queue.queue.Length()
}

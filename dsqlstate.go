package dsqlstate

import (
	"strconv"
)

//go:generate sqlboiler --no-hooks -w "d_users,d_guilds,d_guild_roles,d_channels,d_members,d_channel_overwrites,d_voice_states,d_messages,d_message_revisions,d_message_embeds,d_change_logs,d_meta" postgres

const (
	VersionMajor = 0
	VersionMinor = 3
	VersionPatch = 1
)

var (
	VersionString = strconv.Itoa(VersionMajor) + "." + strconv.Itoa(VersionMinor) + "." + strconv.Itoa(VersionPatch)
)

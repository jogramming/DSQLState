package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessages)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbeds)
	t.Run("DiscordChangeLogs", testDiscordChangeLogs)
	t.Run("DiscordGuilds", testDiscordGuilds)
	t.Run("DiscordGuildRoles", testDiscordGuildRoles)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwrites)
	t.Run("DiscordUsers", testDiscordUsers)
	t.Run("DiscordMembers", testDiscordMembers)
	t.Run("DiscordChannels", testDiscordChannels)
	t.Run("DiscordVoiceStates", testDiscordVoiceStates)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisions)
}

func TestDelete(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesDelete)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsDelete)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsDelete)
	t.Run("DiscordGuilds", testDiscordGuildsDelete)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesDelete)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesDelete)
	t.Run("DiscordUsers", testDiscordUsersDelete)
	t.Run("DiscordMembers", testDiscordMembersDelete)
	t.Run("DiscordChannels", testDiscordChannelsDelete)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesDelete)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesQueryDeleteAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsQueryDeleteAll)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsQueryDeleteAll)
	t.Run("DiscordGuilds", testDiscordGuildsQueryDeleteAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesQueryDeleteAll)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesQueryDeleteAll)
	t.Run("DiscordUsers", testDiscordUsersQueryDeleteAll)
	t.Run("DiscordMembers", testDiscordMembersQueryDeleteAll)
	t.Run("DiscordChannels", testDiscordChannelsQueryDeleteAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesQueryDeleteAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesSliceDeleteAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsSliceDeleteAll)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsSliceDeleteAll)
	t.Run("DiscordGuilds", testDiscordGuildsSliceDeleteAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesSliceDeleteAll)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesSliceDeleteAll)
	t.Run("DiscordUsers", testDiscordUsersSliceDeleteAll)
	t.Run("DiscordMembers", testDiscordMembersSliceDeleteAll)
	t.Run("DiscordChannels", testDiscordChannelsSliceDeleteAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesSliceDeleteAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesExists)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsExists)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsExists)
	t.Run("DiscordGuilds", testDiscordGuildsExists)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesExists)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesExists)
	t.Run("DiscordUsers", testDiscordUsersExists)
	t.Run("DiscordMembers", testDiscordMembersExists)
	t.Run("DiscordChannels", testDiscordChannelsExists)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesExists)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsExists)
}

func TestFind(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesFind)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsFind)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsFind)
	t.Run("DiscordGuilds", testDiscordGuildsFind)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesFind)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesFind)
	t.Run("DiscordUsers", testDiscordUsersFind)
	t.Run("DiscordMembers", testDiscordMembersFind)
	t.Run("DiscordChannels", testDiscordChannelsFind)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesFind)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsFind)
}

func TestBind(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesBind)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsBind)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsBind)
	t.Run("DiscordGuilds", testDiscordGuildsBind)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesBind)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesBind)
	t.Run("DiscordUsers", testDiscordUsersBind)
	t.Run("DiscordMembers", testDiscordMembersBind)
	t.Run("DiscordChannels", testDiscordChannelsBind)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesBind)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsBind)
}

func TestOne(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesOne)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsOne)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsOne)
	t.Run("DiscordGuilds", testDiscordGuildsOne)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesOne)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesOne)
	t.Run("DiscordUsers", testDiscordUsersOne)
	t.Run("DiscordMembers", testDiscordMembersOne)
	t.Run("DiscordChannels", testDiscordChannelsOne)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesOne)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsOne)
}

func TestAll(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsAll)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsAll)
	t.Run("DiscordGuilds", testDiscordGuildsAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesAll)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesAll)
	t.Run("DiscordUsers", testDiscordUsersAll)
	t.Run("DiscordMembers", testDiscordMembersAll)
	t.Run("DiscordChannels", testDiscordChannelsAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsAll)
}

func TestCount(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesCount)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsCount)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsCount)
	t.Run("DiscordGuilds", testDiscordGuildsCount)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesCount)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesCount)
	t.Run("DiscordUsers", testDiscordUsersCount)
	t.Run("DiscordMembers", testDiscordMembersCount)
	t.Run("DiscordChannels", testDiscordChannelsCount)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesCount)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsCount)
}

func TestInsert(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesInsert)
	t.Run("DiscordMessages", testDiscordMessagesInsertWhitelist)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsInsert)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsInsertWhitelist)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsInsert)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsInsertWhitelist)
	t.Run("DiscordGuilds", testDiscordGuildsInsert)
	t.Run("DiscordGuilds", testDiscordGuildsInsertWhitelist)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesInsert)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesInsertWhitelist)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesInsert)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesInsertWhitelist)
	t.Run("DiscordUsers", testDiscordUsersInsert)
	t.Run("DiscordUsers", testDiscordUsersInsertWhitelist)
	t.Run("DiscordMembers", testDiscordMembersInsert)
	t.Run("DiscordMembers", testDiscordMembersInsertWhitelist)
	t.Run("DiscordChannels", testDiscordChannelsInsert)
	t.Run("DiscordChannels", testDiscordChannelsInsertWhitelist)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesInsert)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesInsertWhitelist)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsInsert)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("DiscordMessageEmbedToDiscordMessageUsingMessage", testDiscordMessageEmbedToOneDiscordMessageUsingMessage)
	t.Run("DiscordChannelOverwriteToDiscordChannelUsingChannel", testDiscordChannelOverwriteToOneDiscordChannelUsingChannel)
	t.Run("DiscordMemberToDiscordUserUsingUser", testDiscordMemberToOneDiscordUserUsingUser)
	t.Run("DiscordVoiceStateToDiscordChannelUsingChannel", testDiscordVoiceStateToOneDiscordChannelUsingChannel)
	t.Run("DiscordMessageRevisionToDiscordMessageUsingMessage", testDiscordMessageRevisionToOneDiscordMessageUsingMessage)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("DiscordMessageToMessageDiscordMessageEmbeds", testDiscordMessageToManyMessageDiscordMessageEmbeds)
	t.Run("DiscordMessageToMessageDiscordMessageRevisions", testDiscordMessageToManyMessageDiscordMessageRevisions)
	t.Run("DiscordUserToUserDiscordMembers", testDiscordUserToManyUserDiscordMembers)
	t.Run("DiscordChannelToChannelDiscordChannelOverwrites", testDiscordChannelToManyChannelDiscordChannelOverwrites)
	t.Run("DiscordChannelToChannelDiscordVoiceStates", testDiscordChannelToManyChannelDiscordVoiceStates)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("DiscordMessageEmbedToDiscordMessageUsingMessage", testDiscordMessageEmbedToOneSetOpDiscordMessageUsingMessage)
	t.Run("DiscordChannelOverwriteToDiscordChannelUsingChannel", testDiscordChannelOverwriteToOneSetOpDiscordChannelUsingChannel)
	t.Run("DiscordMemberToDiscordUserUsingUser", testDiscordMemberToOneSetOpDiscordUserUsingUser)
	t.Run("DiscordVoiceStateToDiscordChannelUsingChannel", testDiscordVoiceStateToOneSetOpDiscordChannelUsingChannel)
	t.Run("DiscordMessageRevisionToDiscordMessageUsingMessage", testDiscordMessageRevisionToOneSetOpDiscordMessageUsingMessage)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("DiscordMessageToMessageDiscordMessageEmbeds", testDiscordMessageToManyAddOpMessageDiscordMessageEmbeds)
	t.Run("DiscordMessageToMessageDiscordMessageRevisions", testDiscordMessageToManyAddOpMessageDiscordMessageRevisions)
	t.Run("DiscordUserToUserDiscordMembers", testDiscordUserToManyAddOpUserDiscordMembers)
	t.Run("DiscordChannelToChannelDiscordChannelOverwrites", testDiscordChannelToManyAddOpChannelDiscordChannelOverwrites)
	t.Run("DiscordChannelToChannelDiscordVoiceStates", testDiscordChannelToManyAddOpChannelDiscordVoiceStates)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesReload)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsReload)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsReload)
	t.Run("DiscordGuilds", testDiscordGuildsReload)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesReload)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesReload)
	t.Run("DiscordUsers", testDiscordUsersReload)
	t.Run("DiscordMembers", testDiscordMembersReload)
	t.Run("DiscordChannels", testDiscordChannelsReload)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesReload)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesReloadAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsReloadAll)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsReloadAll)
	t.Run("DiscordGuilds", testDiscordGuildsReloadAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesReloadAll)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesReloadAll)
	t.Run("DiscordUsers", testDiscordUsersReloadAll)
	t.Run("DiscordMembers", testDiscordMembersReloadAll)
	t.Run("DiscordChannels", testDiscordChannelsReloadAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesReloadAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesSelect)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsSelect)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsSelect)
	t.Run("DiscordGuilds", testDiscordGuildsSelect)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesSelect)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesSelect)
	t.Run("DiscordUsers", testDiscordUsersSelect)
	t.Run("DiscordMembers", testDiscordMembersSelect)
	t.Run("DiscordChannels", testDiscordChannelsSelect)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesSelect)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesUpdate)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsUpdate)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsUpdate)
	t.Run("DiscordGuilds", testDiscordGuildsUpdate)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesUpdate)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesUpdate)
	t.Run("DiscordUsers", testDiscordUsersUpdate)
	t.Run("DiscordMembers", testDiscordMembersUpdate)
	t.Run("DiscordChannels", testDiscordChannelsUpdate)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesUpdate)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesSliceUpdateAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsSliceUpdateAll)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsSliceUpdateAll)
	t.Run("DiscordGuilds", testDiscordGuildsSliceUpdateAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesSliceUpdateAll)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesSliceUpdateAll)
	t.Run("DiscordUsers", testDiscordUsersSliceUpdateAll)
	t.Run("DiscordMembers", testDiscordMembersSliceUpdateAll)
	t.Run("DiscordChannels", testDiscordChannelsSliceUpdateAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesSliceUpdateAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("DiscordMessages", testDiscordMessagesUpsert)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsUpsert)
	t.Run("DiscordChangeLogs", testDiscordChangeLogsUpsert)
	t.Run("DiscordGuilds", testDiscordGuildsUpsert)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesUpsert)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesUpsert)
	t.Run("DiscordUsers", testDiscordUsersUpsert)
	t.Run("DiscordMembers", testDiscordMembersUpsert)
	t.Run("DiscordChannels", testDiscordChannelsUpsert)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesUpsert)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsUpsert)
}

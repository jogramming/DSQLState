package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwrites)
	t.Run("DiscordVoiceStates", testDiscordVoiceStates)
	t.Run("DiscordGuildChannels", testDiscordGuildChannels)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannels)
	t.Run("DiscordMembers", testDiscordMembers)
	t.Run("DiscordGuilds", testDiscordGuilds)
	t.Run("DiscordUsers", testDiscordUsers)
	t.Run("DiscordGuildRoles", testDiscordGuildRoles)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbeds)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisions)
	t.Run("DiscordMessages", testDiscordMessages)
}

func TestDelete(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesDelete)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesDelete)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsDelete)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsDelete)
	t.Run("DiscordMembers", testDiscordMembersDelete)
	t.Run("DiscordGuilds", testDiscordGuildsDelete)
	t.Run("DiscordUsers", testDiscordUsersDelete)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesDelete)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsDelete)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsDelete)
	t.Run("DiscordMessages", testDiscordMessagesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesQueryDeleteAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesQueryDeleteAll)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsQueryDeleteAll)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsQueryDeleteAll)
	t.Run("DiscordMembers", testDiscordMembersQueryDeleteAll)
	t.Run("DiscordGuilds", testDiscordGuildsQueryDeleteAll)
	t.Run("DiscordUsers", testDiscordUsersQueryDeleteAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesQueryDeleteAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsQueryDeleteAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsQueryDeleteAll)
	t.Run("DiscordMessages", testDiscordMessagesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesSliceDeleteAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesSliceDeleteAll)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsSliceDeleteAll)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsSliceDeleteAll)
	t.Run("DiscordMembers", testDiscordMembersSliceDeleteAll)
	t.Run("DiscordGuilds", testDiscordGuildsSliceDeleteAll)
	t.Run("DiscordUsers", testDiscordUsersSliceDeleteAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesSliceDeleteAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsSliceDeleteAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsSliceDeleteAll)
	t.Run("DiscordMessages", testDiscordMessagesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesExists)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesExists)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsExists)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsExists)
	t.Run("DiscordMembers", testDiscordMembersExists)
	t.Run("DiscordGuilds", testDiscordGuildsExists)
	t.Run("DiscordUsers", testDiscordUsersExists)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesExists)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsExists)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsExists)
	t.Run("DiscordMessages", testDiscordMessagesExists)
}

func TestFind(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesFind)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesFind)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsFind)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsFind)
	t.Run("DiscordMembers", testDiscordMembersFind)
	t.Run("DiscordGuilds", testDiscordGuildsFind)
	t.Run("DiscordUsers", testDiscordUsersFind)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesFind)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsFind)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsFind)
	t.Run("DiscordMessages", testDiscordMessagesFind)
}

func TestBind(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesBind)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesBind)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsBind)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsBind)
	t.Run("DiscordMembers", testDiscordMembersBind)
	t.Run("DiscordGuilds", testDiscordGuildsBind)
	t.Run("DiscordUsers", testDiscordUsersBind)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesBind)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsBind)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsBind)
	t.Run("DiscordMessages", testDiscordMessagesBind)
}

func TestOne(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesOne)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesOne)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsOne)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsOne)
	t.Run("DiscordMembers", testDiscordMembersOne)
	t.Run("DiscordGuilds", testDiscordGuildsOne)
	t.Run("DiscordUsers", testDiscordUsersOne)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesOne)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsOne)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsOne)
	t.Run("DiscordMessages", testDiscordMessagesOne)
}

func TestAll(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesAll)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsAll)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsAll)
	t.Run("DiscordMembers", testDiscordMembersAll)
	t.Run("DiscordGuilds", testDiscordGuildsAll)
	t.Run("DiscordUsers", testDiscordUsersAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsAll)
	t.Run("DiscordMessages", testDiscordMessagesAll)
}

func TestCount(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesCount)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesCount)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsCount)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsCount)
	t.Run("DiscordMembers", testDiscordMembersCount)
	t.Run("DiscordGuilds", testDiscordGuildsCount)
	t.Run("DiscordUsers", testDiscordUsersCount)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesCount)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsCount)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsCount)
	t.Run("DiscordMessages", testDiscordMessagesCount)
}

func TestInsert(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesInsert)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesInsertWhitelist)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesInsert)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesInsertWhitelist)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsInsert)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsInsertWhitelist)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsInsert)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsInsertWhitelist)
	t.Run("DiscordMembers", testDiscordMembersInsert)
	t.Run("DiscordMembers", testDiscordMembersInsertWhitelist)
	t.Run("DiscordGuilds", testDiscordGuildsInsert)
	t.Run("DiscordGuilds", testDiscordGuildsInsertWhitelist)
	t.Run("DiscordUsers", testDiscordUsersInsert)
	t.Run("DiscordUsers", testDiscordUsersInsertWhitelist)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesInsert)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesInsertWhitelist)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsInsert)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsInsertWhitelist)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsInsert)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsInsertWhitelist)
	t.Run("DiscordMessages", testDiscordMessagesInsert)
	t.Run("DiscordMessages", testDiscordMessagesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("DiscordGuildChannelToDiscordGuildUsingGuild", testDiscordGuildChannelToOneDiscordGuildUsingGuild)
	t.Run("DiscordPrivateChannelToDiscordUserUsingRecipient", testDiscordPrivateChannelToOneDiscordUserUsingRecipient)
	t.Run("DiscordMemberToDiscordUserUsingUser", testDiscordMemberToOneDiscordUserUsingUser)
	t.Run("DiscordMemberToDiscordGuildUsingGuild", testDiscordMemberToOneDiscordGuildUsingGuild)
	t.Run("DiscordGuildRoleToDiscordGuildUsingGuild", testDiscordGuildRoleToOneDiscordGuildUsingGuild)
	t.Run("DiscordMessageEmbedToDiscordMessageUsingMessage", testDiscordMessageEmbedToOneDiscordMessageUsingMessage)
	t.Run("DiscordMessageRevisionToDiscordMessageUsingMessage", testDiscordMessageRevisionToOneDiscordMessageUsingMessage)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("DiscordGuildToGuildDiscordGuildChannels", testDiscordGuildToManyGuildDiscordGuildChannels)
	t.Run("DiscordGuildToGuildDiscordMembers", testDiscordGuildToManyGuildDiscordMembers)
	t.Run("DiscordGuildToGuildDiscordGuildRoles", testDiscordGuildToManyGuildDiscordGuildRoles)
	t.Run("DiscordUserToRecipientDiscordPrivateChannels", testDiscordUserToManyRecipientDiscordPrivateChannels)
	t.Run("DiscordUserToUserDiscordMembers", testDiscordUserToManyUserDiscordMembers)
	t.Run("DiscordMessageToMessageDiscordMessageEmbeds", testDiscordMessageToManyMessageDiscordMessageEmbeds)
	t.Run("DiscordMessageToMessageDiscordMessageRevisions", testDiscordMessageToManyMessageDiscordMessageRevisions)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("DiscordGuildChannelToDiscordGuildUsingGuild", testDiscordGuildChannelToOneSetOpDiscordGuildUsingGuild)
	t.Run("DiscordPrivateChannelToDiscordUserUsingRecipient", testDiscordPrivateChannelToOneSetOpDiscordUserUsingRecipient)
	t.Run("DiscordMemberToDiscordUserUsingUser", testDiscordMemberToOneSetOpDiscordUserUsingUser)
	t.Run("DiscordMemberToDiscordGuildUsingGuild", testDiscordMemberToOneSetOpDiscordGuildUsingGuild)
	t.Run("DiscordGuildRoleToDiscordGuildUsingGuild", testDiscordGuildRoleToOneSetOpDiscordGuildUsingGuild)
	t.Run("DiscordMessageEmbedToDiscordMessageUsingMessage", testDiscordMessageEmbedToOneSetOpDiscordMessageUsingMessage)
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
	t.Run("DiscordGuildToGuildDiscordGuildChannels", testDiscordGuildToManyAddOpGuildDiscordGuildChannels)
	t.Run("DiscordGuildToGuildDiscordMembers", testDiscordGuildToManyAddOpGuildDiscordMembers)
	t.Run("DiscordGuildToGuildDiscordGuildRoles", testDiscordGuildToManyAddOpGuildDiscordGuildRoles)
	t.Run("DiscordUserToRecipientDiscordPrivateChannels", testDiscordUserToManyAddOpRecipientDiscordPrivateChannels)
	t.Run("DiscordUserToUserDiscordMembers", testDiscordUserToManyAddOpUserDiscordMembers)
	t.Run("DiscordMessageToMessageDiscordMessageEmbeds", testDiscordMessageToManyAddOpMessageDiscordMessageEmbeds)
	t.Run("DiscordMessageToMessageDiscordMessageRevisions", testDiscordMessageToManyAddOpMessageDiscordMessageRevisions)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesReload)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesReload)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsReload)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsReload)
	t.Run("DiscordMembers", testDiscordMembersReload)
	t.Run("DiscordGuilds", testDiscordGuildsReload)
	t.Run("DiscordUsers", testDiscordUsersReload)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesReload)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsReload)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsReload)
	t.Run("DiscordMessages", testDiscordMessagesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesReloadAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesReloadAll)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsReloadAll)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsReloadAll)
	t.Run("DiscordMembers", testDiscordMembersReloadAll)
	t.Run("DiscordGuilds", testDiscordGuildsReloadAll)
	t.Run("DiscordUsers", testDiscordUsersReloadAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesReloadAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsReloadAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsReloadAll)
	t.Run("DiscordMessages", testDiscordMessagesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesSelect)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesSelect)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsSelect)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsSelect)
	t.Run("DiscordMembers", testDiscordMembersSelect)
	t.Run("DiscordGuilds", testDiscordGuildsSelect)
	t.Run("DiscordUsers", testDiscordUsersSelect)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesSelect)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsSelect)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsSelect)
	t.Run("DiscordMessages", testDiscordMessagesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesUpdate)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesUpdate)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsUpdate)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsUpdate)
	t.Run("DiscordMembers", testDiscordMembersUpdate)
	t.Run("DiscordGuilds", testDiscordGuildsUpdate)
	t.Run("DiscordUsers", testDiscordUsersUpdate)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesUpdate)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsUpdate)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsUpdate)
	t.Run("DiscordMessages", testDiscordMessagesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesSliceUpdateAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesSliceUpdateAll)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsSliceUpdateAll)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsSliceUpdateAll)
	t.Run("DiscordMembers", testDiscordMembersSliceUpdateAll)
	t.Run("DiscordGuilds", testDiscordGuildsSliceUpdateAll)
	t.Run("DiscordUsers", testDiscordUsersSliceUpdateAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesSliceUpdateAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsSliceUpdateAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsSliceUpdateAll)
	t.Run("DiscordMessages", testDiscordMessagesSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesUpsert)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesUpsert)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsUpsert)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsUpsert)
	t.Run("DiscordMembers", testDiscordMembersUpsert)
	t.Run("DiscordGuilds", testDiscordGuildsUpsert)
	t.Run("DiscordUsers", testDiscordUsersUpsert)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesUpsert)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsUpsert)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsUpsert)
	t.Run("DiscordMessages", testDiscordMessagesUpsert)
}

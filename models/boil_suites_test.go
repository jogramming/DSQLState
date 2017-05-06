package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannels)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwrites)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannels)
	t.Run("DiscordMembers", testDiscordMembers)
	t.Run("DiscordUsers", testDiscordUsers)
	t.Run("DiscordGuilds", testDiscordGuilds)
	t.Run("DiscordGuildRoles", testDiscordGuildRoles)
	t.Run("DiscordMemberRoles", testDiscordMemberRoles)
	t.Run("DiscordVoiceStates", testDiscordVoiceStates)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisions)
	t.Run("DiscordMessages", testDiscordMessages)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbeds)
}

func TestDelete(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsDelete)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesDelete)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsDelete)
	t.Run("DiscordMembers", testDiscordMembersDelete)
	t.Run("DiscordUsers", testDiscordUsersDelete)
	t.Run("DiscordGuilds", testDiscordGuildsDelete)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesDelete)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesDelete)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesDelete)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsDelete)
	t.Run("DiscordMessages", testDiscordMessagesDelete)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsQueryDeleteAll)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesQueryDeleteAll)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsQueryDeleteAll)
	t.Run("DiscordMembers", testDiscordMembersQueryDeleteAll)
	t.Run("DiscordUsers", testDiscordUsersQueryDeleteAll)
	t.Run("DiscordGuilds", testDiscordGuildsQueryDeleteAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesQueryDeleteAll)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesQueryDeleteAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesQueryDeleteAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsQueryDeleteAll)
	t.Run("DiscordMessages", testDiscordMessagesQueryDeleteAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsSliceDeleteAll)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesSliceDeleteAll)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsSliceDeleteAll)
	t.Run("DiscordMembers", testDiscordMembersSliceDeleteAll)
	t.Run("DiscordUsers", testDiscordUsersSliceDeleteAll)
	t.Run("DiscordGuilds", testDiscordGuildsSliceDeleteAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesSliceDeleteAll)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesSliceDeleteAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesSliceDeleteAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsSliceDeleteAll)
	t.Run("DiscordMessages", testDiscordMessagesSliceDeleteAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsExists)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesExists)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsExists)
	t.Run("DiscordMembers", testDiscordMembersExists)
	t.Run("DiscordUsers", testDiscordUsersExists)
	t.Run("DiscordGuilds", testDiscordGuildsExists)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesExists)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesExists)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesExists)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsExists)
	t.Run("DiscordMessages", testDiscordMessagesExists)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsExists)
}

func TestFind(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsFind)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesFind)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsFind)
	t.Run("DiscordMembers", testDiscordMembersFind)
	t.Run("DiscordUsers", testDiscordUsersFind)
	t.Run("DiscordGuilds", testDiscordGuildsFind)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesFind)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesFind)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesFind)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsFind)
	t.Run("DiscordMessages", testDiscordMessagesFind)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsFind)
}

func TestBind(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsBind)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesBind)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsBind)
	t.Run("DiscordMembers", testDiscordMembersBind)
	t.Run("DiscordUsers", testDiscordUsersBind)
	t.Run("DiscordGuilds", testDiscordGuildsBind)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesBind)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesBind)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesBind)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsBind)
	t.Run("DiscordMessages", testDiscordMessagesBind)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsBind)
}

func TestOne(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsOne)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesOne)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsOne)
	t.Run("DiscordMembers", testDiscordMembersOne)
	t.Run("DiscordUsers", testDiscordUsersOne)
	t.Run("DiscordGuilds", testDiscordGuildsOne)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesOne)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesOne)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesOne)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsOne)
	t.Run("DiscordMessages", testDiscordMessagesOne)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsOne)
}

func TestAll(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsAll)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesAll)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsAll)
	t.Run("DiscordMembers", testDiscordMembersAll)
	t.Run("DiscordUsers", testDiscordUsersAll)
	t.Run("DiscordGuilds", testDiscordGuildsAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesAll)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsAll)
	t.Run("DiscordMessages", testDiscordMessagesAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsAll)
}

func TestCount(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsCount)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesCount)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsCount)
	t.Run("DiscordMembers", testDiscordMembersCount)
	t.Run("DiscordUsers", testDiscordUsersCount)
	t.Run("DiscordGuilds", testDiscordGuildsCount)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesCount)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesCount)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesCount)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsCount)
	t.Run("DiscordMessages", testDiscordMessagesCount)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsCount)
}

func TestInsert(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsInsert)
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsInsertWhitelist)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesInsert)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesInsertWhitelist)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsInsert)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsInsertWhitelist)
	t.Run("DiscordMembers", testDiscordMembersInsert)
	t.Run("DiscordMembers", testDiscordMembersInsertWhitelist)
	t.Run("DiscordUsers", testDiscordUsersInsert)
	t.Run("DiscordUsers", testDiscordUsersInsertWhitelist)
	t.Run("DiscordGuilds", testDiscordGuildsInsert)
	t.Run("DiscordGuilds", testDiscordGuildsInsertWhitelist)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesInsert)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesInsertWhitelist)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesInsert)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesInsertWhitelist)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesInsert)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesInsertWhitelist)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsInsert)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsInsertWhitelist)
	t.Run("DiscordMessages", testDiscordMessagesInsert)
	t.Run("DiscordMessages", testDiscordMessagesInsertWhitelist)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsInsert)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("DiscordGuildChannelToDiscordGuildUsingGuild", testDiscordGuildChannelToOneDiscordGuildUsingGuild)
	t.Run("DiscordPrivateChannelToDiscordUserUsingRecipient", testDiscordPrivateChannelToOneDiscordUserUsingRecipient)
	t.Run("DiscordMemberToDiscordUserUsingUser", testDiscordMemberToOneDiscordUserUsingUser)
	t.Run("DiscordMemberToDiscordGuildUsingGuild", testDiscordMemberToOneDiscordGuildUsingGuild)
	t.Run("DiscordGuildRoleToDiscordGuildUsingGuild", testDiscordGuildRoleToOneDiscordGuildUsingGuild)
	t.Run("DiscordMemberRoleToDiscordUserUsingUser", testDiscordMemberRoleToOneDiscordUserUsingUser)
	t.Run("DiscordMemberRoleToDiscordGuildUsingGuild", testDiscordMemberRoleToOneDiscordGuildUsingGuild)
	t.Run("DiscordMemberRoleToDiscordGuildRoleUsingRole", testDiscordMemberRoleToOneDiscordGuildRoleUsingRole)
	t.Run("DiscordMessageRevisionToDiscordMessageUsingMessage", testDiscordMessageRevisionToOneDiscordMessageUsingMessage)
	t.Run("DiscordMessageEmbedToDiscordMessageUsingMessage", testDiscordMessageEmbedToOneDiscordMessageUsingMessage)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("DiscordUserToRecipientDiscordPrivateChannels", testDiscordUserToManyRecipientDiscordPrivateChannels)
	t.Run("DiscordUserToUserDiscordMembers", testDiscordUserToManyUserDiscordMembers)
	t.Run("DiscordUserToUserDiscordMemberRoles", testDiscordUserToManyUserDiscordMemberRoles)
	t.Run("DiscordGuildToGuildDiscordGuildChannels", testDiscordGuildToManyGuildDiscordGuildChannels)
	t.Run("DiscordGuildToGuildDiscordMembers", testDiscordGuildToManyGuildDiscordMembers)
	t.Run("DiscordGuildToGuildDiscordGuildRoles", testDiscordGuildToManyGuildDiscordGuildRoles)
	t.Run("DiscordGuildToGuildDiscordMemberRoles", testDiscordGuildToManyGuildDiscordMemberRoles)
	t.Run("DiscordGuildRoleToRoleDiscordMemberRoles", testDiscordGuildRoleToManyRoleDiscordMemberRoles)
	t.Run("DiscordMessageToMessageDiscordMessageRevisions", testDiscordMessageToManyMessageDiscordMessageRevisions)
	t.Run("DiscordMessageToMessageDiscordMessageEmbeds", testDiscordMessageToManyMessageDiscordMessageEmbeds)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("DiscordGuildChannelToDiscordGuildUsingGuild", testDiscordGuildChannelToOneSetOpDiscordGuildUsingGuild)
	t.Run("DiscordPrivateChannelToDiscordUserUsingRecipient", testDiscordPrivateChannelToOneSetOpDiscordUserUsingRecipient)
	t.Run("DiscordMemberToDiscordUserUsingUser", testDiscordMemberToOneSetOpDiscordUserUsingUser)
	t.Run("DiscordMemberToDiscordGuildUsingGuild", testDiscordMemberToOneSetOpDiscordGuildUsingGuild)
	t.Run("DiscordGuildRoleToDiscordGuildUsingGuild", testDiscordGuildRoleToOneSetOpDiscordGuildUsingGuild)
	t.Run("DiscordMemberRoleToDiscordUserUsingUser", testDiscordMemberRoleToOneSetOpDiscordUserUsingUser)
	t.Run("DiscordMemberRoleToDiscordGuildUsingGuild", testDiscordMemberRoleToOneSetOpDiscordGuildUsingGuild)
	t.Run("DiscordMemberRoleToDiscordGuildRoleUsingRole", testDiscordMemberRoleToOneSetOpDiscordGuildRoleUsingRole)
	t.Run("DiscordMessageRevisionToDiscordMessageUsingMessage", testDiscordMessageRevisionToOneSetOpDiscordMessageUsingMessage)
	t.Run("DiscordMessageEmbedToDiscordMessageUsingMessage", testDiscordMessageEmbedToOneSetOpDiscordMessageUsingMessage)
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
	t.Run("DiscordUserToRecipientDiscordPrivateChannels", testDiscordUserToManyAddOpRecipientDiscordPrivateChannels)
	t.Run("DiscordUserToUserDiscordMembers", testDiscordUserToManyAddOpUserDiscordMembers)
	t.Run("DiscordUserToUserDiscordMemberRoles", testDiscordUserToManyAddOpUserDiscordMemberRoles)
	t.Run("DiscordGuildToGuildDiscordGuildChannels", testDiscordGuildToManyAddOpGuildDiscordGuildChannels)
	t.Run("DiscordGuildToGuildDiscordMembers", testDiscordGuildToManyAddOpGuildDiscordMembers)
	t.Run("DiscordGuildToGuildDiscordGuildRoles", testDiscordGuildToManyAddOpGuildDiscordGuildRoles)
	t.Run("DiscordGuildToGuildDiscordMemberRoles", testDiscordGuildToManyAddOpGuildDiscordMemberRoles)
	t.Run("DiscordGuildRoleToRoleDiscordMemberRoles", testDiscordGuildRoleToManyAddOpRoleDiscordMemberRoles)
	t.Run("DiscordMessageToMessageDiscordMessageRevisions", testDiscordMessageToManyAddOpMessageDiscordMessageRevisions)
	t.Run("DiscordMessageToMessageDiscordMessageEmbeds", testDiscordMessageToManyAddOpMessageDiscordMessageEmbeds)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsReload)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesReload)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsReload)
	t.Run("DiscordMembers", testDiscordMembersReload)
	t.Run("DiscordUsers", testDiscordUsersReload)
	t.Run("DiscordGuilds", testDiscordGuildsReload)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesReload)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesReload)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesReload)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsReload)
	t.Run("DiscordMessages", testDiscordMessagesReload)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsReloadAll)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesReloadAll)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsReloadAll)
	t.Run("DiscordMembers", testDiscordMembersReloadAll)
	t.Run("DiscordUsers", testDiscordUsersReloadAll)
	t.Run("DiscordGuilds", testDiscordGuildsReloadAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesReloadAll)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesReloadAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesReloadAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsReloadAll)
	t.Run("DiscordMessages", testDiscordMessagesReloadAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsSelect)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesSelect)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsSelect)
	t.Run("DiscordMembers", testDiscordMembersSelect)
	t.Run("DiscordUsers", testDiscordUsersSelect)
	t.Run("DiscordGuilds", testDiscordGuildsSelect)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesSelect)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesSelect)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesSelect)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsSelect)
	t.Run("DiscordMessages", testDiscordMessagesSelect)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsUpdate)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesUpdate)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsUpdate)
	t.Run("DiscordMembers", testDiscordMembersUpdate)
	t.Run("DiscordUsers", testDiscordUsersUpdate)
	t.Run("DiscordGuilds", testDiscordGuildsUpdate)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesUpdate)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesUpdate)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesUpdate)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsUpdate)
	t.Run("DiscordMessages", testDiscordMessagesUpdate)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsSliceUpdateAll)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesSliceUpdateAll)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsSliceUpdateAll)
	t.Run("DiscordMembers", testDiscordMembersSliceUpdateAll)
	t.Run("DiscordUsers", testDiscordUsersSliceUpdateAll)
	t.Run("DiscordGuilds", testDiscordGuildsSliceUpdateAll)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesSliceUpdateAll)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesSliceUpdateAll)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesSliceUpdateAll)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsSliceUpdateAll)
	t.Run("DiscordMessages", testDiscordMessagesSliceUpdateAll)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("DiscordGuildChannels", testDiscordGuildChannelsUpsert)
	t.Run("DiscordChannelOverwrites", testDiscordChannelOverwritesUpsert)
	t.Run("DiscordPrivateChannels", testDiscordPrivateChannelsUpsert)
	t.Run("DiscordMembers", testDiscordMembersUpsert)
	t.Run("DiscordUsers", testDiscordUsersUpsert)
	t.Run("DiscordGuilds", testDiscordGuildsUpsert)
	t.Run("DiscordGuildRoles", testDiscordGuildRolesUpsert)
	t.Run("DiscordMemberRoles", testDiscordMemberRolesUpsert)
	t.Run("DiscordVoiceStates", testDiscordVoiceStatesUpsert)
	t.Run("DiscordMessageRevisions", testDiscordMessageRevisionsUpsert)
	t.Run("DiscordMessages", testDiscordMessagesUpsert)
	t.Run("DiscordMessageEmbeds", testDiscordMessageEmbedsUpsert)
}

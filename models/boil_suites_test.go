package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("DGuilds", testDGuilds)
	t.Run("DGuildRoles", testDGuildRoles)
	t.Run("DMembers", testDMembers)
	t.Run("DChannelOverwrites", testDChannelOverwrites)
	t.Run("DChannels", testDChannels)
	t.Run("DUsers", testDUsers)
	t.Run("DVoiceStates", testDVoiceStates)
	t.Run("DMessageRevisions", testDMessageRevisions)
	t.Run("DMessages", testDMessages)
	t.Run("DMessageEmbeds", testDMessageEmbeds)
	t.Run("DMeta", testDMeta)
}

func TestDelete(t *testing.T) {
	t.Run("DGuilds", testDGuildsDelete)
	t.Run("DGuildRoles", testDGuildRolesDelete)
	t.Run("DMembers", testDMembersDelete)
	t.Run("DChannelOverwrites", testDChannelOverwritesDelete)
	t.Run("DChannels", testDChannelsDelete)
	t.Run("DUsers", testDUsersDelete)
	t.Run("DVoiceStates", testDVoiceStatesDelete)
	t.Run("DMessageRevisions", testDMessageRevisionsDelete)
	t.Run("DMessages", testDMessagesDelete)
	t.Run("DMessageEmbeds", testDMessageEmbedsDelete)
	t.Run("DMeta", testDMetaDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("DGuilds", testDGuildsQueryDeleteAll)
	t.Run("DGuildRoles", testDGuildRolesQueryDeleteAll)
	t.Run("DMembers", testDMembersQueryDeleteAll)
	t.Run("DChannelOverwrites", testDChannelOverwritesQueryDeleteAll)
	t.Run("DChannels", testDChannelsQueryDeleteAll)
	t.Run("DUsers", testDUsersQueryDeleteAll)
	t.Run("DVoiceStates", testDVoiceStatesQueryDeleteAll)
	t.Run("DMessageRevisions", testDMessageRevisionsQueryDeleteAll)
	t.Run("DMessages", testDMessagesQueryDeleteAll)
	t.Run("DMessageEmbeds", testDMessageEmbedsQueryDeleteAll)
	t.Run("DMeta", testDMetaQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("DGuilds", testDGuildsSliceDeleteAll)
	t.Run("DGuildRoles", testDGuildRolesSliceDeleteAll)
	t.Run("DMembers", testDMembersSliceDeleteAll)
	t.Run("DChannelOverwrites", testDChannelOverwritesSliceDeleteAll)
	t.Run("DChannels", testDChannelsSliceDeleteAll)
	t.Run("DUsers", testDUsersSliceDeleteAll)
	t.Run("DVoiceStates", testDVoiceStatesSliceDeleteAll)
	t.Run("DMessageRevisions", testDMessageRevisionsSliceDeleteAll)
	t.Run("DMessages", testDMessagesSliceDeleteAll)
	t.Run("DMessageEmbeds", testDMessageEmbedsSliceDeleteAll)
	t.Run("DMeta", testDMetaSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("DGuilds", testDGuildsExists)
	t.Run("DGuildRoles", testDGuildRolesExists)
	t.Run("DMembers", testDMembersExists)
	t.Run("DChannelOverwrites", testDChannelOverwritesExists)
	t.Run("DChannels", testDChannelsExists)
	t.Run("DUsers", testDUsersExists)
	t.Run("DVoiceStates", testDVoiceStatesExists)
	t.Run("DMessageRevisions", testDMessageRevisionsExists)
	t.Run("DMessages", testDMessagesExists)
	t.Run("DMessageEmbeds", testDMessageEmbedsExists)
	t.Run("DMeta", testDMetaExists)
}

func TestFind(t *testing.T) {
	t.Run("DGuilds", testDGuildsFind)
	t.Run("DGuildRoles", testDGuildRolesFind)
	t.Run("DMembers", testDMembersFind)
	t.Run("DChannelOverwrites", testDChannelOverwritesFind)
	t.Run("DChannels", testDChannelsFind)
	t.Run("DUsers", testDUsersFind)
	t.Run("DVoiceStates", testDVoiceStatesFind)
	t.Run("DMessageRevisions", testDMessageRevisionsFind)
	t.Run("DMessages", testDMessagesFind)
	t.Run("DMessageEmbeds", testDMessageEmbedsFind)
	t.Run("DMeta", testDMetaFind)
}

func TestBind(t *testing.T) {
	t.Run("DGuilds", testDGuildsBind)
	t.Run("DGuildRoles", testDGuildRolesBind)
	t.Run("DMembers", testDMembersBind)
	t.Run("DChannelOverwrites", testDChannelOverwritesBind)
	t.Run("DChannels", testDChannelsBind)
	t.Run("DUsers", testDUsersBind)
	t.Run("DVoiceStates", testDVoiceStatesBind)
	t.Run("DMessageRevisions", testDMessageRevisionsBind)
	t.Run("DMessages", testDMessagesBind)
	t.Run("DMessageEmbeds", testDMessageEmbedsBind)
	t.Run("DMeta", testDMetaBind)
}

func TestOne(t *testing.T) {
	t.Run("DGuilds", testDGuildsOne)
	t.Run("DGuildRoles", testDGuildRolesOne)
	t.Run("DMembers", testDMembersOne)
	t.Run("DChannelOverwrites", testDChannelOverwritesOne)
	t.Run("DChannels", testDChannelsOne)
	t.Run("DUsers", testDUsersOne)
	t.Run("DVoiceStates", testDVoiceStatesOne)
	t.Run("DMessageRevisions", testDMessageRevisionsOne)
	t.Run("DMessages", testDMessagesOne)
	t.Run("DMessageEmbeds", testDMessageEmbedsOne)
	t.Run("DMeta", testDMetaOne)
}

func TestAll(t *testing.T) {
	t.Run("DGuilds", testDGuildsAll)
	t.Run("DGuildRoles", testDGuildRolesAll)
	t.Run("DMembers", testDMembersAll)
	t.Run("DChannelOverwrites", testDChannelOverwritesAll)
	t.Run("DChannels", testDChannelsAll)
	t.Run("DUsers", testDUsersAll)
	t.Run("DVoiceStates", testDVoiceStatesAll)
	t.Run("DMessageRevisions", testDMessageRevisionsAll)
	t.Run("DMessages", testDMessagesAll)
	t.Run("DMessageEmbeds", testDMessageEmbedsAll)
	t.Run("DMeta", testDMetaAll)
}

func TestCount(t *testing.T) {
	t.Run("DGuilds", testDGuildsCount)
	t.Run("DGuildRoles", testDGuildRolesCount)
	t.Run("DMembers", testDMembersCount)
	t.Run("DChannelOverwrites", testDChannelOverwritesCount)
	t.Run("DChannels", testDChannelsCount)
	t.Run("DUsers", testDUsersCount)
	t.Run("DVoiceStates", testDVoiceStatesCount)
	t.Run("DMessageRevisions", testDMessageRevisionsCount)
	t.Run("DMessages", testDMessagesCount)
	t.Run("DMessageEmbeds", testDMessageEmbedsCount)
	t.Run("DMeta", testDMetaCount)
}

func TestInsert(t *testing.T) {
	t.Run("DGuilds", testDGuildsInsert)
	t.Run("DGuilds", testDGuildsInsertWhitelist)
	t.Run("DGuildRoles", testDGuildRolesInsert)
	t.Run("DGuildRoles", testDGuildRolesInsertWhitelist)
	t.Run("DMembers", testDMembersInsert)
	t.Run("DMembers", testDMembersInsertWhitelist)
	t.Run("DChannelOverwrites", testDChannelOverwritesInsert)
	t.Run("DChannelOverwrites", testDChannelOverwritesInsertWhitelist)
	t.Run("DChannels", testDChannelsInsert)
	t.Run("DChannels", testDChannelsInsertWhitelist)
	t.Run("DUsers", testDUsersInsert)
	t.Run("DUsers", testDUsersInsertWhitelist)
	t.Run("DVoiceStates", testDVoiceStatesInsert)
	t.Run("DVoiceStates", testDVoiceStatesInsertWhitelist)
	t.Run("DMessageRevisions", testDMessageRevisionsInsert)
	t.Run("DMessageRevisions", testDMessageRevisionsInsertWhitelist)
	t.Run("DMessages", testDMessagesInsert)
	t.Run("DMessages", testDMessagesInsertWhitelist)
	t.Run("DMessageEmbeds", testDMessageEmbedsInsert)
	t.Run("DMessageEmbeds", testDMessageEmbedsInsertWhitelist)
	t.Run("DMeta", testDMetaInsert)
	t.Run("DMeta", testDMetaInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("DMemberToDUserUsingUser", testDMemberToOneDUserUsingUser)
	t.Run("DChannelOverwriteToDChannelUsingChannel", testDChannelOverwriteToOneDChannelUsingChannel)
	t.Run("DVoiceStateToDChannelUsingChannel", testDVoiceStateToOneDChannelUsingChannel)
	t.Run("DMessageRevisionToDMessageUsingMessage", testDMessageRevisionToOneDMessageUsingMessage)
	t.Run("DMessageEmbedToDMessageUsingMessage", testDMessageEmbedToOneDMessageUsingMessage)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("DChannelToChannelDChannelOverwrites", testDChannelToManyChannelDChannelOverwrites)
	t.Run("DChannelToChannelDVoiceStates", testDChannelToManyChannelDVoiceStates)
	t.Run("DUserToUserDMembers", testDUserToManyUserDMembers)
	t.Run("DMessageToMessageDMessageRevisions", testDMessageToManyMessageDMessageRevisions)
	t.Run("DMessageToMessageDMessageEmbeds", testDMessageToManyMessageDMessageEmbeds)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("DMemberToDUserUsingUser", testDMemberToOneSetOpDUserUsingUser)
	t.Run("DChannelOverwriteToDChannelUsingChannel", testDChannelOverwriteToOneSetOpDChannelUsingChannel)
	t.Run("DVoiceStateToDChannelUsingChannel", testDVoiceStateToOneSetOpDChannelUsingChannel)
	t.Run("DMessageRevisionToDMessageUsingMessage", testDMessageRevisionToOneSetOpDMessageUsingMessage)
	t.Run("DMessageEmbedToDMessageUsingMessage", testDMessageEmbedToOneSetOpDMessageUsingMessage)
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
	t.Run("DChannelToChannelDChannelOverwrites", testDChannelToManyAddOpChannelDChannelOverwrites)
	t.Run("DChannelToChannelDVoiceStates", testDChannelToManyAddOpChannelDVoiceStates)
	t.Run("DUserToUserDMembers", testDUserToManyAddOpUserDMembers)
	t.Run("DMessageToMessageDMessageRevisions", testDMessageToManyAddOpMessageDMessageRevisions)
	t.Run("DMessageToMessageDMessageEmbeds", testDMessageToManyAddOpMessageDMessageEmbeds)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("DGuilds", testDGuildsReload)
	t.Run("DGuildRoles", testDGuildRolesReload)
	t.Run("DMembers", testDMembersReload)
	t.Run("DChannelOverwrites", testDChannelOverwritesReload)
	t.Run("DChannels", testDChannelsReload)
	t.Run("DUsers", testDUsersReload)
	t.Run("DVoiceStates", testDVoiceStatesReload)
	t.Run("DMessageRevisions", testDMessageRevisionsReload)
	t.Run("DMessages", testDMessagesReload)
	t.Run("DMessageEmbeds", testDMessageEmbedsReload)
	t.Run("DMeta", testDMetaReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("DGuilds", testDGuildsReloadAll)
	t.Run("DGuildRoles", testDGuildRolesReloadAll)
	t.Run("DMembers", testDMembersReloadAll)
	t.Run("DChannelOverwrites", testDChannelOverwritesReloadAll)
	t.Run("DChannels", testDChannelsReloadAll)
	t.Run("DUsers", testDUsersReloadAll)
	t.Run("DVoiceStates", testDVoiceStatesReloadAll)
	t.Run("DMessageRevisions", testDMessageRevisionsReloadAll)
	t.Run("DMessages", testDMessagesReloadAll)
	t.Run("DMessageEmbeds", testDMessageEmbedsReloadAll)
	t.Run("DMeta", testDMetaReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("DGuilds", testDGuildsSelect)
	t.Run("DGuildRoles", testDGuildRolesSelect)
	t.Run("DMembers", testDMembersSelect)
	t.Run("DChannelOverwrites", testDChannelOverwritesSelect)
	t.Run("DChannels", testDChannelsSelect)
	t.Run("DUsers", testDUsersSelect)
	t.Run("DVoiceStates", testDVoiceStatesSelect)
	t.Run("DMessageRevisions", testDMessageRevisionsSelect)
	t.Run("DMessages", testDMessagesSelect)
	t.Run("DMessageEmbeds", testDMessageEmbedsSelect)
	t.Run("DMeta", testDMetaSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("DGuilds", testDGuildsUpdate)
	t.Run("DGuildRoles", testDGuildRolesUpdate)
	t.Run("DMembers", testDMembersUpdate)
	t.Run("DChannelOverwrites", testDChannelOverwritesUpdate)
	t.Run("DChannels", testDChannelsUpdate)
	t.Run("DUsers", testDUsersUpdate)
	t.Run("DVoiceStates", testDVoiceStatesUpdate)
	t.Run("DMessageRevisions", testDMessageRevisionsUpdate)
	t.Run("DMessages", testDMessagesUpdate)
	t.Run("DMessageEmbeds", testDMessageEmbedsUpdate)
	t.Run("DMeta", testDMetaUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("DGuilds", testDGuildsSliceUpdateAll)
	t.Run("DGuildRoles", testDGuildRolesSliceUpdateAll)
	t.Run("DMembers", testDMembersSliceUpdateAll)
	t.Run("DChannelOverwrites", testDChannelOverwritesSliceUpdateAll)
	t.Run("DChannels", testDChannelsSliceUpdateAll)
	t.Run("DUsers", testDUsersSliceUpdateAll)
	t.Run("DVoiceStates", testDVoiceStatesSliceUpdateAll)
	t.Run("DMessageRevisions", testDMessageRevisionsSliceUpdateAll)
	t.Run("DMessages", testDMessagesSliceUpdateAll)
	t.Run("DMessageEmbeds", testDMessageEmbedsSliceUpdateAll)
	t.Run("DMeta", testDMetaSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("DGuilds", testDGuildsUpsert)
	t.Run("DGuildRoles", testDGuildRolesUpsert)
	t.Run("DMembers", testDMembersUpsert)
	t.Run("DChannelOverwrites", testDChannelOverwritesUpsert)
	t.Run("DChannels", testDChannelsUpsert)
	t.Run("DUsers", testDUsersUpsert)
	t.Run("DVoiceStates", testDVoiceStatesUpsert)
	t.Run("DMessageRevisions", testDMessageRevisionsUpsert)
	t.Run("DMessages", testDMessagesUpsert)
	t.Run("DMessageEmbeds", testDMessageEmbedsUpsert)
	t.Run("DMeta", testDMetaUpsert)
}

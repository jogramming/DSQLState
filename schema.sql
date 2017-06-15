DROP TABLE IF EXISTS d_users CASCADE;
CREATE TABLE IF NOT EXISTS d_users (
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

CREATE INDEX IF NOT EXISTS d_users_lower_idx ON d_users(lower(username));

DROP TABLE IF EXISTS d_guilds CASCADE;
CREATE TABLE IF NOT EXISTS d_guilds (
	id bigint PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	left_at TIMESTAMP WITH TIME ZONE,
	synced bool NOT NULL,

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

DROP TABLE IF EXISTS d_guild_roles CASCADE;
CREATE TABLE IF NOT EXISTS d_guild_roles (
	id bigint PRIMARY KEY,
	guild_id bigint NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE,
	synced bool NOT NULL,

	name text NOT NULL,
	managed bool NOT NULL,
	mentionable bool NOT NULL,
	hoist bool NOT NULL,
	color int NOT NULL,
	position int NOT NULL,
	permissions int NOT NULL
);

DROP TABLE IF EXISTS d_channels CASCADE;
CREATE TABLE IF NOT EXISTS d_channels (
	id bigint PRIMARY KEY,
	guild_id bigint,
	recipient_id bigint,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE,
	synced bool NOT NULL,

	name text NOT NULL,
	topic text NOT NULL,
	type text NOT NULL,
	last_message_id bigint NOT NULL,
	position int NOT NULL,
	bitrate int NOT NULL
);

CREATE INDEX IF NOT EXISTS d_channels_guild_idx ON d_channels(guild_id);
CREATE INDEX IF NOT EXISTS d_channels_recipient_idx ON d_channels(recipient_id);

DROP TABLE IF EXISTS d_channel_overwrites CASCADE;
CREATE TABLE IF NOT EXISTS d_channel_overwrites (
	id bigint NOT NULL,
	channel_id bigint references d_channels(id) NOT NULL,

	type varchar(10) NOT NULL,
	allow int NOT NULL,
	deny int NOT NULL,

	PRIMARY KEY(channel_id, id)
);

CREATE INDEX IF NOT EXISTS d_channel_overwrites_channel_idx ON d_channel_overwrites(channel_id);
CREATE INDEX IF NOT EXISTS d_channel_overwrites_idx ON d_channel_overwrites(id);

DROP TABLE IF EXISTS d_members;

CREATE TABLE IF NOT EXISTS d_members (
	user_id bigint references d_users(id) NOT NULL,
	guild_id bigint NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	synced bool NOT NULL,

	left_at TIMESTAMP WITH TIME ZONE,

	joined_at TIMESTAMP WITH TIME ZONE NOT NULL,
	nick varchar(32) NOT NULL,
	deaf bool NOT NULL,
	mute bool NOT NULL,
	roles bigint[] NOT NULL,

	PRIMARY KEY(user_id, guild_id)
);

CREATE INDEX IF NOT EXISTS d_members_user_idx ON d_members(user_id);
CREATE INDEX IF NOT EXISTS d_members_guild_idx ON d_members(guild_id);

DROP TABLE IF EXISTS d_voice_states;
CREATE TABLE IF NOT EXISTS d_voice_states (
	user_id bigint NOT NULL,
	guild_id bigint,
	channel_id bigint references d_channels(id) NOT NULL,
	session_id text NOT NULL,

	surpress bool NOT NULL,
	self_mute bool NOT NULL,
	self_deaf bool NOT NULL,
	mute bool NOT NULL,
	deaf bool NOT NULL,

	PRIMARY KEY(guild_id, user_id)
);

CREATE INDEX IF NOT EXISTS d_voice_states_guild_idx ON d_voice_states(guild_id);
CREATE INDEX IF NOT EXISTS d_voice_states_channel_idx ON d_voice_states(channel_id);

DROP TABLE IF EXISTS d_messages CASCADE;
CREATE TABLE IF NOT EXISTS d_messages (
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

CREATE INDEX IF NOT EXISTS d_messages_channel_idx ON d_messages(channel_id);

DROP TABLE IF EXISTS d_message_revisions CASCADE;
CREATE TABLE IF NOT EXISTS d_message_revisions (
	revision_num int,
	message_id bigint references d_messages(id) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,

	content text NOT NULL,
	embeds bigint[] NOT NULL,

	mentions bigint[] NOT NULL,
	mention_roles bigint[] NOT NULL,

	PRIMARY KEY(message_id, revision_num)
);

CREATE INDEX IF NOT EXISTS d_message_revisions_message_idx ON d_message_revisions(message_id);

DROP TABLE IF EXISTS d_message_embeds;
CREATE TABLE IF NOT EXISTS d_message_embeds (
	id bigserial PRIMARY KEY,
	message_id bigint references d_messages(id) NOT NULL,
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

CREATE INDEX IF NOT EXISTS d_message_embeds_message_idx ON d_message_embeds(message_id);
 
DROP TABLE IF EXISTS d_change_logs;
CREATE TABLE IF NOT EXISTS d_change_logs (
	id bigserial PRIMARY KEY,
	field int NOT NULL,
	valueInt bigint,
	valueString text,
	valueBOol bool
)

DROP TABLE IF EXISTS d_meta;
CREATE TABLE IF NOT EXISTS d_meta (
	key text PRIMARY KEY,
	value bytea NOT NULL 
);
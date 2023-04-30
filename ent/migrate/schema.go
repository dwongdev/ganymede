// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChannelsColumns holds the columns for the "channels" table.
	ChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "ext_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "display_name", Type: field.TypeString, Unique: true},
		{Name: "image_path", Type: field.TypeString},
		{Name: "retention", Type: field.TypeBool, Default: false},
		{Name: "retention_days", Type: field.TypeInt64, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ChannelsTable holds the schema information for the "channels" table.
	ChannelsTable = &schema.Table{
		Name:       "channels",
		Columns:    ChannelsColumns,
		PrimaryKey: []*schema.Column{ChannelsColumns[0]},
	}
	// LivesColumns holds the columns for the "lives" table.
	LivesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "watch_live", Type: field.TypeBool, Default: true},
		{Name: "watch_vod", Type: field.TypeBool, Default: false},
		{Name: "download_archives", Type: field.TypeBool, Default: false},
		{Name: "download_highlights", Type: field.TypeBool, Default: false},
		{Name: "download_uploads", Type: field.TypeBool, Default: false},
		{Name: "download_sub_only", Type: field.TypeBool, Default: false},
		{Name: "is_live", Type: field.TypeBool, Default: false},
		{Name: "archive_chat", Type: field.TypeBool, Default: true},
		{Name: "resolution", Type: field.TypeString, Nullable: true, Default: "best"},
		{Name: "last_live", Type: field.TypeTime},
		{Name: "render_chat", Type: field.TypeBool, Default: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "channel_live", Type: field.TypeUUID},
	}
	// LivesTable holds the schema information for the "lives" table.
	LivesTable = &schema.Table{
		Name:       "lives",
		Columns:    LivesColumns,
		PrimaryKey: []*schema.Column{LivesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lives_channels_live",
				Columns:    []*schema.Column{LivesColumns[14]},
				RefColumns: []*schema.Column{ChannelsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LiveCategoriesColumns holds the columns for the "live_categories" table.
	LiveCategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "live_id", Type: field.TypeUUID},
	}
	// LiveCategoriesTable holds the schema information for the "live_categories" table.
	LiveCategoriesTable = &schema.Table{
		Name:       "live_categories",
		Columns:    LiveCategoriesColumns,
		PrimaryKey: []*schema.Column{LiveCategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "live_categories_lives_categories",
				Columns:    []*schema.Column{LiveCategoriesColumns[2]},
				RefColumns: []*schema.Column{LivesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PlaybacksColumns holds the columns for the "playbacks" table.
	PlaybacksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "vod_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "time", Type: field.TypeInt, Default: 0},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"in_progress", "finished"}, Default: "in_progress"},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// PlaybacksTable holds the schema information for the "playbacks" table.
	PlaybacksTable = &schema.Table{
		Name:       "playbacks",
		Columns:    PlaybacksColumns,
		PrimaryKey: []*schema.Column{PlaybacksColumns[0]},
	}
	// PlaylistsColumns holds the columns for the "playlists" table.
	PlaylistsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "thumbnail_path", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// PlaylistsTable holds the schema information for the "playlists" table.
	PlaylistsTable = &schema.Table{
		Name:       "playlists",
		Columns:    PlaylistsColumns,
		PrimaryKey: []*schema.Column{PlaylistsColumns[0]},
	}
	// QueuesColumns holds the columns for the "queues" table.
	QueuesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "live_archive", Type: field.TypeBool, Default: false},
		{Name: "on_hold", Type: field.TypeBool, Default: false},
		{Name: "video_processing", Type: field.TypeBool, Default: true},
		{Name: "chat_processing", Type: field.TypeBool, Default: true},
		{Name: "processing", Type: field.TypeBool, Default: true},
		{Name: "task_vod_create_folder", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "running", "pending", "failed"}, Default: "pending"},
		{Name: "task_vod_download_thumbnail", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "running", "pending", "failed"}, Default: "pending"},
		{Name: "task_vod_save_info", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "running", "pending", "failed"}, Default: "pending"},
		{Name: "task_video_download", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "running", "pending", "failed"}, Default: "pending"},
		{Name: "task_video_convert", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "running", "pending", "failed"}, Default: "pending"},
		{Name: "task_video_move", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "running", "pending", "failed"}, Default: "pending"},
		{Name: "task_chat_download", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "running", "pending", "failed"}, Default: "pending"},
		{Name: "task_chat_convert", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "running", "pending", "failed"}, Default: "pending"},
		{Name: "task_chat_render", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "running", "pending", "failed"}, Default: "pending"},
		{Name: "task_chat_move", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "running", "pending", "failed"}, Default: "pending"},
		{Name: "chat_start", Type: field.TypeTime, Nullable: true},
		{Name: "render_chat", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "vod_queue", Type: field.TypeUUID, Unique: true},
	}
	// QueuesTable holds the schema information for the "queues" table.
	QueuesTable = &schema.Table{
		Name:       "queues",
		Columns:    QueuesColumns,
		PrimaryKey: []*schema.Column{QueuesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "queues_vods_queue",
				Columns:    []*schema.Column{QueuesColumns[20]},
				RefColumns: []*schema.Column{VodsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TwitchCategoriesColumns holds the columns for the "twitch_categories" table.
	TwitchCategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "box_art_url", Type: field.TypeString, Nullable: true},
		{Name: "igdb_id", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// TwitchCategoriesTable holds the schema information for the "twitch_categories" table.
	TwitchCategoriesTable = &schema.Table{
		Name:       "twitch_categories",
		Columns:    TwitchCategoriesColumns,
		PrimaryKey: []*schema.Column{TwitchCategoriesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "sub", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "oauth", Type: field.TypeBool, Default: false},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "editor", "archiver", "user"}, Default: "user"},
		{Name: "webhook", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VodsColumns holds the columns for the "vods" table.
	VodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "ext_id", Type: field.TypeString},
		{Name: "platform", Type: field.TypeEnum, Enums: []string{"twitch", "youtube"}, Default: "twitch"},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"archive", "live", "highlight", "upload", "clip"}, Default: "archive"},
		{Name: "title", Type: field.TypeString},
		{Name: "duration", Type: field.TypeInt, Default: 1},
		{Name: "views", Type: field.TypeInt, Default: 1},
		{Name: "resolution", Type: field.TypeString, Nullable: true},
		{Name: "processing", Type: field.TypeBool, Default: false},
		{Name: "thumbnail_path", Type: field.TypeString, Nullable: true},
		{Name: "web_thumbnail_path", Type: field.TypeString},
		{Name: "video_path", Type: field.TypeString},
		{Name: "chat_path", Type: field.TypeString, Nullable: true},
		{Name: "chat_video_path", Type: field.TypeString, Nullable: true},
		{Name: "info_path", Type: field.TypeString, Nullable: true},
		{Name: "caption_path", Type: field.TypeString, Nullable: true},
		{Name: "folder_name", Type: field.TypeString, Nullable: true},
		{Name: "file_name", Type: field.TypeString, Nullable: true},
		{Name: "locked", Type: field.TypeBool, Default: false},
		{Name: "streamed_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "channel_vods", Type: field.TypeUUID},
	}
	// VodsTable holds the schema information for the "vods" table.
	VodsTable = &schema.Table{
		Name:       "vods",
		Columns:    VodsColumns,
		PrimaryKey: []*schema.Column{VodsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vods_channels_vods",
				Columns:    []*schema.Column{VodsColumns[22]},
				RefColumns: []*schema.Column{ChannelsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PlaylistVodsColumns holds the columns for the "playlist_vods" table.
	PlaylistVodsColumns = []*schema.Column{
		{Name: "playlist_id", Type: field.TypeUUID},
		{Name: "vod_id", Type: field.TypeUUID},
	}
	// PlaylistVodsTable holds the schema information for the "playlist_vods" table.
	PlaylistVodsTable = &schema.Table{
		Name:       "playlist_vods",
		Columns:    PlaylistVodsColumns,
		PrimaryKey: []*schema.Column{PlaylistVodsColumns[0], PlaylistVodsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "playlist_vods_playlist_id",
				Columns:    []*schema.Column{PlaylistVodsColumns[0]},
				RefColumns: []*schema.Column{PlaylistsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "playlist_vods_vod_id",
				Columns:    []*schema.Column{PlaylistVodsColumns[1]},
				RefColumns: []*schema.Column{VodsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChannelsTable,
		LivesTable,
		LiveCategoriesTable,
		PlaybacksTable,
		PlaylistsTable,
		QueuesTable,
		TwitchCategoriesTable,
		UsersTable,
		VodsTable,
		PlaylistVodsTable,
	}
)

func init() {
	LivesTable.ForeignKeys[0].RefTable = ChannelsTable
	LiveCategoriesTable.ForeignKeys[0].RefTable = LivesTable
	QueuesTable.ForeignKeys[0].RefTable = VodsTable
	VodsTable.ForeignKeys[0].RefTable = ChannelsTable
	PlaylistVodsTable.ForeignKeys[0].RefTable = PlaylistsTable
	PlaylistVodsTable.ForeignKeys[1].RefTable = VodsTable
}

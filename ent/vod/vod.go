// Code generated by ent, DO NOT EDIT.

package vod

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/internal/utils"
)

const (
	// Label holds the string label denoting the vod type in the database.
	Label = "vod"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldExtID holds the string denoting the ext_id field in the database.
	FieldExtID = "ext_id"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldViews holds the string denoting the views field in the database.
	FieldViews = "views"
	// FieldResolution holds the string denoting the resolution field in the database.
	FieldResolution = "resolution"
	// FieldProcessing holds the string denoting the processing field in the database.
	FieldProcessing = "processing"
	// FieldThumbnailPath holds the string denoting the thumbnail_path field in the database.
	FieldThumbnailPath = "thumbnail_path"
	// FieldWebThumbnailPath holds the string denoting the web_thumbnail_path field in the database.
	FieldWebThumbnailPath = "web_thumbnail_path"
	// FieldVideoPath holds the string denoting the video_path field in the database.
	FieldVideoPath = "video_path"
	// FieldChatPath holds the string denoting the chat_path field in the database.
	FieldChatPath = "chat_path"
	// FieldChatVideoPath holds the string denoting the chat_video_path field in the database.
	FieldChatVideoPath = "chat_video_path"
	// FieldInfoPath holds the string denoting the info_path field in the database.
	FieldInfoPath = "info_path"
	// FieldCaptionPath holds the string denoting the caption_path field in the database.
	FieldCaptionPath = "caption_path"
	// FieldFolderName holds the string denoting the folder_name field in the database.
	FieldFolderName = "folder_name"
	// FieldFileName holds the string denoting the file_name field in the database.
	FieldFileName = "file_name"
	// FieldLocked holds the string denoting the locked field in the database.
	FieldLocked = "locked"
	// FieldStreamedAt holds the string denoting the streamed_at field in the database.
	FieldStreamedAt = "streamed_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeChannel holds the string denoting the channel edge name in mutations.
	EdgeChannel = "channel"
	// EdgeQueue holds the string denoting the queue edge name in mutations.
	EdgeQueue = "queue"
	// EdgePlaylists holds the string denoting the playlists edge name in mutations.
	EdgePlaylists = "playlists"
	// Table holds the table name of the vod in the database.
	Table = "vods"
	// ChannelTable is the table that holds the channel relation/edge.
	ChannelTable = "vods"
	// ChannelInverseTable is the table name for the Channel entity.
	// It exists in this package in order to avoid circular dependency with the "channel" package.
	ChannelInverseTable = "channels"
	// ChannelColumn is the table column denoting the channel relation/edge.
	ChannelColumn = "channel_vods"
	// QueueTable is the table that holds the queue relation/edge.
	QueueTable = "queues"
	// QueueInverseTable is the table name for the Queue entity.
	// It exists in this package in order to avoid circular dependency with the "queue" package.
	QueueInverseTable = "queues"
	// QueueColumn is the table column denoting the queue relation/edge.
	QueueColumn = "vod_queue"
	// PlaylistsTable is the table that holds the playlists relation/edge. The primary key declared below.
	PlaylistsTable = "playlist_vods"
	// PlaylistsInverseTable is the table name for the Playlist entity.
	// It exists in this package in order to avoid circular dependency with the "playlist" package.
	PlaylistsInverseTable = "playlists"
)

// Columns holds all SQL columns for vod fields.
var Columns = []string{
	FieldID,
	FieldExtID,
	FieldPlatform,
	FieldType,
	FieldTitle,
	FieldDuration,
	FieldViews,
	FieldResolution,
	FieldProcessing,
	FieldThumbnailPath,
	FieldWebThumbnailPath,
	FieldVideoPath,
	FieldChatPath,
	FieldChatVideoPath,
	FieldInfoPath,
	FieldCaptionPath,
	FieldFolderName,
	FieldFileName,
	FieldLocked,
	FieldStreamedAt,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "vods"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"channel_vods",
}

var (
	// PlaylistsPrimaryKey and PlaylistsColumn2 are the table columns denoting the
	// primary key for the playlists relation (M2M).
	PlaylistsPrimaryKey = []string{"playlist_id", "vod_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDuration holds the default value on creation for the "duration" field.
	DefaultDuration int
	// DefaultViews holds the default value on creation for the "views" field.
	DefaultViews int
	// DefaultProcessing holds the default value on creation for the "processing" field.
	DefaultProcessing bool
	// DefaultLocked holds the default value on creation for the "locked" field.
	DefaultLocked bool
	// DefaultStreamedAt holds the default value on creation for the "streamed_at" field.
	DefaultStreamedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

const DefaultPlatform utils.VodPlatform = "twitch"

// PlatformValidator is a validator for the "platform" field enum values. It is called by the builders before save.
func PlatformValidator(pl utils.VodPlatform) error {
	switch pl {
	case "twitch", "youtube":
		return nil
	default:
		return fmt.Errorf("vod: invalid enum value for platform field: %q", pl)
	}
}

const DefaultType utils.VodType = "archive"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type utils.VodType) error {
	switch _type {
	case "archive", "live", "highlight", "upload", "clip":
		return nil
	default:
		return fmt.Errorf("vod: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Vod queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByExtID orders the results by the ext_id field.
func ByExtID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtID, opts...).ToFunc()
}

// ByPlatform orders the results by the platform field.
func ByPlatform(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlatform, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByViews orders the results by the views field.
func ByViews(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldViews, opts...).ToFunc()
}

// ByResolution orders the results by the resolution field.
func ByResolution(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResolution, opts...).ToFunc()
}

// ByProcessing orders the results by the processing field.
func ByProcessing(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProcessing, opts...).ToFunc()
}

// ByThumbnailPath orders the results by the thumbnail_path field.
func ByThumbnailPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThumbnailPath, opts...).ToFunc()
}

// ByWebThumbnailPath orders the results by the web_thumbnail_path field.
func ByWebThumbnailPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWebThumbnailPath, opts...).ToFunc()
}

// ByVideoPath orders the results by the video_path field.
func ByVideoPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVideoPath, opts...).ToFunc()
}

// ByChatPath orders the results by the chat_path field.
func ByChatPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChatPath, opts...).ToFunc()
}

// ByChatVideoPath orders the results by the chat_video_path field.
func ByChatVideoPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChatVideoPath, opts...).ToFunc()
}

// ByInfoPath orders the results by the info_path field.
func ByInfoPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInfoPath, opts...).ToFunc()
}

// ByCaptionPath orders the results by the caption_path field.
func ByCaptionPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCaptionPath, opts...).ToFunc()
}

// ByFolderName orders the results by the folder_name field.
func ByFolderName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFolderName, opts...).ToFunc()
}

// ByFileName orders the results by the file_name field.
func ByFileName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileName, opts...).ToFunc()
}

// ByLocked orders the results by the locked field.
func ByLocked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocked, opts...).ToFunc()
}

// ByStreamedAt orders the results by the streamed_at field.
func ByStreamedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStreamedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByChannelField orders the results by channel field.
func ByChannelField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChannelStep(), sql.OrderByField(field, opts...))
	}
}

// ByQueueField orders the results by queue field.
func ByQueueField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQueueStep(), sql.OrderByField(field, opts...))
	}
}

// ByPlaylistsCount orders the results by playlists count.
func ByPlaylistsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlaylistsStep(), opts...)
	}
}

// ByPlaylists orders the results by playlists terms.
func ByPlaylists(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaylistsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newChannelStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChannelInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ChannelTable, ChannelColumn),
	)
}
func newQueueStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QueueInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, QueueTable, QueueColumn),
	)
}
func newPlaylistsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaylistsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PlaylistsTable, PlaylistsPrimaryKey...),
	)
}

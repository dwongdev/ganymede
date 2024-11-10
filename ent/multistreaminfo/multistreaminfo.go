// Code generated by ent, DO NOT EDIT.

package multistreaminfo

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the multistreaminfo type in the database.
	Label = "multistream_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDelayMs holds the string denoting the delay_ms field in the database.
	FieldDelayMs = "delay_ms"
	// EdgeVod holds the string denoting the vod edge name in mutations.
	EdgeVod = "vod"
	// EdgePlaylist holds the string denoting the playlist edge name in mutations.
	EdgePlaylist = "playlist"
	// Table holds the table name of the multistreaminfo in the database.
	Table = "multistream_infos"
	// VodTable is the table that holds the vod relation/edge.
	VodTable = "multistream_infos"
	// VodInverseTable is the table name for the Vod entity.
	// It exists in this package in order to avoid circular dependency with the "vod" package.
	VodInverseTable = "vods"
	// VodColumn is the table column denoting the vod relation/edge.
	VodColumn = "multistream_info_vod"
	// PlaylistTable is the table that holds the playlist relation/edge.
	PlaylistTable = "multistream_infos"
	// PlaylistInverseTable is the table name for the Playlist entity.
	// It exists in this package in order to avoid circular dependency with the "playlist" package.
	PlaylistInverseTable = "playlists"
	// PlaylistColumn is the table column denoting the playlist relation/edge.
	PlaylistColumn = "playlist_multistream_info"
)

// Columns holds all SQL columns for multistreaminfo fields.
var Columns = []string{
	FieldID,
	FieldDelayMs,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "multistream_infos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"multistream_info_vod",
	"playlist_multistream_info",
}

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

// OrderOption defines the ordering options for the MultistreamInfo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDelayMs orders the results by the delay_ms field.
func ByDelayMs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelayMs, opts...).ToFunc()
}

// ByVodField orders the results by vod field.
func ByVodField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVodStep(), sql.OrderByField(field, opts...))
	}
}

// ByPlaylistField orders the results by playlist field.
func ByPlaylistField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaylistStep(), sql.OrderByField(field, opts...))
	}
}
func newVodStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VodInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, VodTable, VodColumn),
	)
}
func newPlaylistStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaylistInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PlaylistTable, PlaylistColumn),
	)
}

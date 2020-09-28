package zingmp3

// SongInfoResponse is reponse from zingmp3 get song info api
type SongInfoResponse struct {
	Err  int                  `json:"err,omitempty"`
	Msg  string               `json:"msg"`
	Data SongInfoResponseData `json:"data"`
}

// SongInfoResponseData ...
type SongInfoResponseData struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	ArtistName      string    `json:"artist_names"`
	RawID           int64     `json:"raw_id"`
	Alias           string    `json:"alias"`
	IsZMA           bool      `json:"is_zma"`
	Link            string    `json:"link"`
	Thumbnail       string    `json:"thumbnail"`
	ThumbnailMedium string    `json:"thumbnail_medium"`
	Lyric           string    `json:"lyric"`
	Gernes          []*Gerne  `json:"gernes"`
	Artists         []*Artist `json:"artists"`
	Streaming       Streaming `json:"streaming"`
	Duration        int32     `json:"duration"`
}

// Album ...
type Album struct {
	RawID           int64    `json:"raw_id"`
	ID              string   `json:"id"`
	Link            string   `json:"link"`
	Title           string   `json:"title"`
	ArtistName      string   `json:"artist_name"`
	Artist          []Artist `json:"artists"`
	Thumbnail       string   `json:"thumbnail"`
	ThumbnailMedium string   `json:"thumbnail_medium"`
}

// Gerne ...
type Gerne struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Alias string `json:"au-my"`
	Link  string `json:"link"`
}

// Artist ...
type Artist struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	Cover      string `json:"cover"`
	Thumbnail  string `json:"thumbnail"`
	Spotlight  bool   `json:"spotlight"`
	Follow     int64  `json:"follow"`
	PlaylistID string `json:"playlistId"`
}

// Streaming ...
type Streaming struct {
	Audio AudioFormat `json:"default"`
}

// AudioFormat ...
type AudioFormat struct {
	Format128 string `json:"128"`
	Format320 string `json:"320"`
	Msg       string `json:"msg"`
}

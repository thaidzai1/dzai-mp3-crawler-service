package downloadsongs

// DownloadSongSchema ...
type DownloadSongSchema struct {
	Web         string   `yaml:"web"`
	DownloadDir string   `yaml:"downloadDir"`
	SongsUrls   []string `yaml:"songUrls"`
}

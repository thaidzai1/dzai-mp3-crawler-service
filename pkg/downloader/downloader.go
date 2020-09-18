package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/crawler"
	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/httpreq/types/zingmp3"
)

// DownloadFile will download file from url to filePath
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// DownloadSongs will download song of each source web
func DownloadSongs(source string, urls []string, folderPath string) error {
	if source == "zingmp3" {
		var zingMp3Responses []zingmp3.SongInfoResponse
		for _, url := range urls {
			zingMp3Res, err := crawler.CrawlZingMp3Song(url)
			if err != nil {
				color.Red("%s is failed!!!", url)
			}
			zingMp3Responses = append(zingMp3Responses, *zingMp3Res)
		}
		color.Cyan("Downloading %s...", zingMp3Responses[0].Data.Title)
		filePath := fmt.Sprintf("%s/%s.mp3", folderPath, zingMp3Responses[0].Data.Title)
		err := DownloadFile(filePath, "https:"+zingMp3Responses[0].Data.Streaming.Audio.Format128)
		if err != nil {
			return err
		}
		color.Green("Download %s done!", zingMp3Responses[0].Data.Title)
		return nil
	}

	return nil
}

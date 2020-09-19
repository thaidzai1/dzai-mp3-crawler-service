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
func DownloadFile(folderPath string, filename string, url string, noti chan string) {
	filePath := fmt.Sprintf("%s/%s.mp3", folderPath, filename)
	errorMsg := fmt.Sprintf("Download %s failed", filename)
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		noti <-errorMsg
		return
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		noti <-errorMsg
		return
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		noti <-errorMsg
		return
	}
	noti <-fmt.Sprintf("Download %s successfully!", filename)
	return
}

// DownloadSongs will download song of each source web
func DownloadSongs(source string, urls []string, folderPath string) error {
	if source == "zingmp3" {
		var zingMp3Responses []*zingmp3.SongInfoResponse
		for _, url := range urls {
			zingMp3Res, err := crawler.CrawlZingMp3Song(url)
			if err != nil {
				color.Red("%s is failed!!!", url)
			}
			zingMp3Responses = append(zingMp3Responses, zingMp3Res...)
		}
		fmt.Println(len(zingMp3Responses))
		downloaderChan := make(chan string)
		for _, zingRes := range zingMp3Responses {
			color.Cyan("Downloading %s...", zingRes.Data.Title)
			go DownloadFile(folderPath, zingRes.Data.Title, "https:"+zingRes.Data.Streaming.Audio.Format128, downloaderChan)
		}

		downloadedFile := 0
		for {
			select {
			case message := <- downloaderChan: {
				color.Cyan(message)
				downloadedFile++
			}
			}
			if downloadedFile == len(zingMp3Responses) {
				break
			}
		}
		color.Green("DONE!")
		return nil
	}

	return nil
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/assemble"
	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/crawler"
	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/httpreq"
	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/httpreq/types/zingmp3"
)

func main() {
	dir, err := ioutil.TempDir("", "chromedp-crawler")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		// use headless browser on production
		chromedp.Flag("headless", false),
		chromedp.UserDataDir(dir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create chromedp context
	chromeCtx, chromeCtxCancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer chromeCtxCancel()

	// ensure that the browser process is started
	if err := chromedp.Run(chromeCtx); err != nil {
		panic(err)
	}

	// start scraping zingmp3
	zingHTML := crawler.ScrapingZingMp3(chromeCtx, "https://zingmp3.vn/album/Today-s-Top-Hits-Various-Artists/ZWZC0WCE.html", "#main-body")

	songCodes := crawler.GetZingMp3SongCodes(zingHTML)
	fmt.Printf("songCodes: %s\n", songCodes)

	songInfoAPIs := assemble.ZingMp3SongAPIs(songCodes, "/song/get-song-info")
	fmt.Printf("songInfoAPIs: %s\n", songInfoAPIs)

	zingMp3ResBytes, err := httpreq.Call(songInfoAPIs[0], "GET", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("zingMp3Res: %s \n", string(zingMp3ResBytes))

	zingMp3Res := &zingmp3.SongInfoResponse{}
	err = json.Unmarshal(zingMp3ResBytes, zingMp3Res)
	if err != nil {
		panic(err)
	}
	fmt.Printf("zingMp3Res: %v \n", zingMp3Res.Data.Streaming)
}

package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/assemble"
	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/httpreq"
	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/httpreq/types/zingmp3"
)

// ScrapingZingMp3 will return body html of zingmp3
func ScrapingZingMp3(ctx context.Context, url string, selector string) string {
	var html string
	var scrapingTasks = chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.WaitVisible(`button[type="submit"]`),
		chromedp.Click(`button[type="submit"]`, chromedp.NodeVisible),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}

			html, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	}
	chromedp.Run(ctx, scrapingTasks...)
	return html
}

// GetZingMp3SongCodes will return list of song code from zingmp3
func GetZingMp3SongCodes(str string) []string {
	uniqPrefixIdentification := `data-for="more-`
	songCodeRegex := regexp.MustCompile(fmt.Sprintf(`%s\w+"`, uniqPrefixIdentification))
	strContainSongCodes := songCodeRegex.FindAllString(str, -1)
	if strContainSongCodes == nil {
		return nil
	}
	var songCodes []string
	for _, str := range strContainSongCodes {
		songCode := str[len(uniqPrefixIdentification) : len(str)-1]
		songCodes = append(songCodes, songCode)
	}
	return songCodes
}

// CrawlZingMp3Song will crawl zingmp3 page by url
func CrawlZingMp3Song(url string) (*zingmp3.SongInfoResponse, error) {
	dir, err := ioutil.TempDir("", "chromedp-crawler")
	if err != nil {
		return nil, err
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
	zingHTML := ScrapingZingMp3(chromeCtx, url, "#main-body")
	codes := GetZingMp3SongCodes(zingHTML)
	apis := assemble.ZingMp3SongAPIs(codes, "/song/get-song-info")
	zingMp3ResBytes, err := httpreq.Call(apis[0], "GET", nil)
	if err != nil {
		return nil, err
	}

	zingMp3Res := &zingmp3.SongInfoResponse{}
	err = json.Unmarshal(zingMp3ResBytes, zingMp3Res)
	if err != nil {
		return nil, err
	}
	return zingMp3Res, nil
}
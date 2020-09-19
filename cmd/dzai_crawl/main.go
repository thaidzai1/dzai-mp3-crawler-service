package main

import (
	"fmt"

	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/crawler"
)

func main() {
	zingMp3Res, err := crawler.CrawlZingMp3Song("https://zingmp3.vn/album/Today-s-Top-Hits-Various-Artists/ZWZC0WCE.html")
	fmt.Println(zingMp3Res, err)
}

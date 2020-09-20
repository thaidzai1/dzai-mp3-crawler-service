package main

import (
	"context"
	"fmt"

	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/crawler"
)

var (
	ctxCancel context.CancelFunc
	ctx context.Context
)

func main() {
	ctx, ctxCancel = context.WithCancel(context.Background())
	zingMp3Res, err := crawler.CrawlZingMp3Song(ctx, "https://zingmp3.vn/album/Today-s-Top-Hits-Various-Artists/ZWZC0WCE.html")
	fmt.Println(zingMp3Res, err)

	<-ctx.Done()
}

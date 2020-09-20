package main

import (
	"context"
	"flag"
	"os"

	"github.com/fatih/color"
	"github.com/thaidzai285/dzai-mp3-crawler-service/internal/pkg/schema"
	"github.com/thaidzai285/dzai-mp3-crawler-service/internal/pkg/validator"
	"github.com/thaidzai285/dzai-mp3-crawler-service/pkg/downloader"
)

var (
	ctxCancel   context.CancelFunc
	ctx         context.Context
	flSongsFile = flag.String("download", "", "Path to list songs want to download")
)

func main() {
	ctx, ctxCancel = context.WithCancel(context.Background())
	flag.Parse()
	if *flSongsFile == "" {
		panic("Error songs file not found")
	}

	func() {
		defer ctxCancel()
		color.Cyan("Parsing config file...")
		config := schema.LoadDownloadSongsConfig(*flSongsFile)
		validator.Schema(config)
		_, err := os.Stat(config.DownloadDir)
		if os.IsNotExist(err) {
			color.Yellow("%s not found!\nSystem is creating folder %s...\n", config.DownloadDir, config.DownloadDir)
			os.MkdirAll(config.DownloadDir, os.ModePerm)
			_, err = os.Stat(config.DownloadDir)
			if os.IsNotExist(err) {
				color.Red("System error!")
				os.Exit(0)
			}
			color.Cyan("%s is created successfully!\n", config.DownloadDir)
		}
		color.Cyan("Start downloading... Please patient!")
		err = downloader.DownloadSongs(ctx, config.Web, config.SongsUrls, config.DownloadDir)
		if err != nil {
			color.Red("System error!")
			os.Exit(0)
		}
		color.Green("DONE!")
	}()
	<-ctx.Done()
}

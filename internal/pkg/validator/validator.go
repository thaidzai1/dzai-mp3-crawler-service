package validator

import (
	"fmt"
	"os"
	"regexp"

	"github.com/thaidzai285/dzai-mp3-crawler-service/internal/types/downloadsongs"
)

var (
	supportedWeb = map[string]string{"zingmp3": "https://zingmp3.vn"}
	httpRegex    = regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
)

// Schema will validate schema
func Schema(schema interface{}) {
	switch v := schema.(type) {
	case *downloadsongs.DownloadSongSchema:
		if supportedWeb[v.Web] == "" {
			fmt.Printf("Web %s is not supported so far\n", v.Web)
			os.Exit(0)
		}
		for _, url := range v.SongsUrls {
			if !httpRegex.Match([]byte(url)) {
				fmt.Printf("Url %s is not supported so far \n", url)
				os.Exit(0)
			}
			if len(v.SongsUrls) > 10 {
				fmt.Printf("Max Urls can be crawled is 10")
				os.Exit(0)
			}
		}
	default:
		fmt.Println("default")
	}
}

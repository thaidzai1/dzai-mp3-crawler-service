package schema

import (
	"io/ioutil"

	"github.com/thaidzai285/dzai-mp3-crawler-service/internal/types/downloadsongs"
	"gopkg.in/yaml.v2"
)

// LoadDownloadSongsConfig will return download songs configuartion from users
func LoadDownloadSongsConfig(filePath string) *downloadsongs.DownloadSongSchema {
	schema := &downloadsongs.DownloadSongSchema{}
	bytesSchema, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bytesSchema, schema)
	if err != nil {
		panic(err)
	}
	return schema
}

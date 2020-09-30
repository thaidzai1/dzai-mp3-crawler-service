package dcrawl

import (
	"context"
	"log"

	"github.com/thaidzai285/dzai-mp3-crawler-service/internal/pkg/crawler"
	"github.com/thaidzai285/dzai-mp3-protobuf/pkg/pb"
)

// CrawlerService ...
type CrawlerService struct{}

// NewCrawlerService ...
func NewCrawlerService() *CrawlerService {
	return &CrawlerService{}
}

// Crawl ...
func (c *CrawlerService) Crawl(ctx context.Context, in *pb.CrawlRequest) (*pb.CrawlResponse, error) {
	log.Println("in", in)
	if in.Source == "zingmp3" {
		// https://zingmp3.vn/album/Today-s-Top-Hits-Various-Artists/ZWZC0WCE.html
		zingMp3Res, err := crawler.CrawlZingMp3Song(ctx, in.Urls[0])
		if err != nil {
			return &pb.CrawlResponse{
				Success: false,
				Message: err.Error(),
			}, err
		}
		var songsData []*pb.Song
		for _, res := range zingMp3Res {
			var genres []*pb.Genre
			for _, genre := range res.Data.Gernes {
				genres = append(genres, &pb.Genre{
					Alias: genre.Alias,
					Id:    genre.ID,
					Link:  genre.Link,
					Name:  genre.Name,
				})
			}
			songsData = append(songsData, &pb.Song{
				Id:              res.Data.ID,
				Title:           res.Data.Title,
				ArtistsNames:    res.Data.ArtistName,
				Genres:          genres,
				Duration:        res.Data.Duration,
				Thumbnail:       res.Data.Thumbnail,
				ThumbnailMedium: res.Data.ThumbnailMedium,
			})
		}
		crawlResponse := &pb.CrawlResponse{
			Success: true,
			Message: "success",
			Data:    songsData,
		}
		return crawlResponse, nil
	}
	return &pb.CrawlResponse{}, nil
}

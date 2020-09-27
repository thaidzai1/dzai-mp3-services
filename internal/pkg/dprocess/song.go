package dprocess

import (
	"context"
	"log"

	"github.com/thaidzai285/dzai-mp3-protobuf/pkg/pb"
)

// SongService ...
type SongService struct {
	pb.CrawlerClient
}

// NewSongService ...
func NewSongService(dzaiCrawler pb.CrawlerClient) *SongService {
	return &SongService{dzaiCrawler}
}

var genres = map[string][]string{
	"usuk": {
		"https://zingmp3.vn/album/Top-100-Pop-Au-My-Hay-Nhat-Various-Artists/ZWZB96AB.html",
		"https://zingmp3.vn/album/Top-100-Nhac-Rap-Hip-Hop-Au-My-Hay-Nhat-Various-Artists/ZWZB96AD.html",
		"https://zingmp3.vn/album/Top-100-Nhac-Electronic-Dance-Au-My-Hay-Nhat-Various-Artists/ZWZB96C7.html",
	},
	"kpop": {
		"https://zingmp3.vn/album/Top-100-Nhac-Han-Quoc-Hay-NhatVarious-Artists/ZWZB96DC.html",
	},
	"vpop": {
		"https://zingmp3.vn/album/Top-100-Nhac-Rap-Viet-Nam-Hay-Nhat-Various-Artists/ZWZB96AI.html",
		"https://zingmp3.vn/album/Top-100-Bai-Hat-Nhac-Tre-Hay-Nhat-Various-Artists/ZWZB969E.html",
	},
}

// GetSongs ...
func (s *SongService) GetSongs(ctx context.Context, in *pb.SongRequestParams) (*pb.GetSongsResponse, error) {
	log.Printf("Received: %v", in)
	urls := genres[in.Genre]
	if in.Genre == "" || len(urls) == 0 {
		return &pb.GetSongsResponse{
			Success: false,
			Message: "Unsupported genre",
		}, nil
	}

	crawlResponse, err := s.Crawl(ctx, &pb.CrawlRequest{Source: "zingmp3", Urls: urls})
	if err != nil {
		return &pb.GetSongsResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.GetSongsResponse{
		Success: true,
		Message: "success",
		Data:    crawlResponse.Data}, nil
}

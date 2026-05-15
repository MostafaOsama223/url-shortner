package service

import (
	"fmt"
	"log"
	"shortner-service/repo"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UrlService struct {
	repo *repo.UrlRepo
}

func (s *UrlService) Repo() *repo.UrlRepo {
	return s.repo
}

func NewUrlService(urlRepo *repo.UrlRepo) *UrlService {
	return &UrlService{
		repo: urlRepo,
	}
}

func (s *UrlService) ShortenUrl(url string) (string, error) {
	id := Base10ToBase62(Base16ToBase10(bson.NewObjectID().Hex()))
	shortUrl := id

	record, err := s.repo.InsertUrlRecord(&repo.UrlRecord{
		ID:          id,
		OriginalUrl: url,
		ShortUrl:    shortUrl,
	})

	if err != nil {
		log.Printf("ERROR: Failed to insert URL record: %v", err)
		return "", fmt.Errorf("failed to store URL in database: %w", err)
	}
	log.Printf("Successfully stored URL: original=%s, short=%s", record.OriginalUrl, record.ShortUrl)
	return record.ShortUrl, nil
}

func (s *UrlService) GetOriginalUrl(shortUrl string) (string, error) {
	originalUrl, err := s.repo.FindOriginalUrl(shortUrl)
	if err != nil {
		log.Printf("ERROR: Failed to find original URL for short URL '%s': %v", shortUrl, err)
		return "", fmt.Errorf("failed to retrieve original URL: %w", err)
	}

	log.Printf("Successfully retrieved original URL for short URL '%s': %s", shortUrl, originalUrl)
	return originalUrl, nil
}

package service

import (
	"fmt"
	"log"
	"math/rand"
	"shortner-service/repo"
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

// TODO: This is a very basic implementation.
func (s *UrlService) generateShortUrl() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (s *UrlService) ShortenUrl(url string) (string, error) {
	shortUrl := s.generateShortUrl()
	record, err := s.repo.InsertUrlRecord(&repo.UrlRecord{
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

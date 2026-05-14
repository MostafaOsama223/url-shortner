package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UrlAPIHandler) shortenUrlHandler(c *gin.Context) {
	url := RequestBodyToStruct(c, &ShortenUrlRequest{})

	if url.URL == "" {
		log.Printf("ERROR: Empty URL in request")
		Fail(c, http.StatusBadRequest, "invalid_url", "URL cannot be empty")
		return
	}

	shortenedUrl, err := h.urlService.ShortenUrl(url.URL)

	if err != nil {
		log.Printf("ERROR: Failed to shorten URL: %v", err)
		Fail(c, http.StatusInternalServerError, "shorten_failed", err.Error())
		return
	}

	c.JSON(http.StatusCreated, ShortenUrlResponse{ShortUrl: shortenedUrl})
}

func (h *UrlAPIHandler) redirectHandler(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	if shortUrl == "" {
		log.Printf("ERROR: Empty short URL in request")
		Fail(c, http.StatusBadRequest, "invalid_short_url", "Short URL cannot be empty")
		return
	}

	originalUrl, err := h.urlService.GetOriginalUrl(shortUrl)
	if err != nil {
		log.Printf("ERROR: Failed to retrieve original URL for short URL '%s': %v", shortUrl, err)
		Fail(c, http.StatusNotFound, "not_found", "Original URL not found")
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalUrl)
}

package api

type ShortenUrlRequest struct {
	URL string `json:"url"`
}

type ShortenUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

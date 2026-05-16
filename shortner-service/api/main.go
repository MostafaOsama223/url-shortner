package api

import (
	"github.com/MostafaOsama223/shortner-service/service"

	"github.com/gin-gonic/gin"
)

type UrlAPIHandler struct {
	router     *gin.Engine
	urlService *service.UrlService
}

func NewUrlAPIHandler(urlService *service.UrlService) *UrlAPIHandler {
	router := gin.Default()
	handler := &UrlAPIHandler{
		router:     router,
		urlService: urlService,
	}

	handler.InitRoutes()

	return handler
}

func (h *UrlAPIHandler) InitRoutes() {
	{
		v1 := h.Group("/api/v1/url")
		v1.POST("/shorten", h.shortenUrlHandler)
		v1.GET("/:shortUrl", h.redirectHandler)
	}
}

func (h *UrlAPIHandler) Run() {
	h.router.Run()
}

func (h *UrlAPIHandler) GET(path string, handler gin.HandlerFunc) {
	h.router.GET(path, handler)
}

func (h *UrlAPIHandler) Group(path string) *gin.RouterGroup {
	return h.router.Group(path)
}

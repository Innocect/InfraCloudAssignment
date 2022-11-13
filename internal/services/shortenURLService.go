package services

import (
	"context"
	"net/http"
	"shortlink/internal/config"
	"sync"
)

var shortenURLSvcStruct ShortenURLService
var shortenURLServiceOnce sync.Once

type ShortenURLService interface {
	ValidateRequest(ctx context.Context, req *http.Request) (
		validationErrors error)
	ProcessRequest(ctx context.Context, req *http.Request, email string, name string, config *config.RouterConfig) (
		*http.Response, error)
}

type shortenURLService struct {
	config *config.RouterConfig
}

// InitShortenURLService ...
func InitShortenURLService(config *config.RouterConfig) ShortenURLService {
	shortenURLServiceOnce.Do(func() {
		shortenURLSvcStruct = &shortenURLService{config: config}
	})
	return shortenURLSvcStruct
}

// GetShortenURLService ...
func GetShortenURLService() ShortenURLService {
	if shortenURLSvcStruct == nil {
		panic("ShortenURLService not initialized")
	}
	return shortenURLSvcStruct
}

func (service *shortenURLService) ValidateRequest(ctx context.Context, req *http.Request) (validationError error) {
	return nil
}

func (service *shortenURLService) ProcessRequest(ctx context.Context, req *http.Request, email string, name string, config *config.RouterConfig) (*http.Response, error) {
	return nil, nil
}

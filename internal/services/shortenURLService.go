package services

import (
	"context"
	"errors"
	"fmt"
	"shortlink/internal/config"
	"shortlink/internal/dao"
	"shortlink/internal/models"
	"shortlink/internal/utils"
	"sync"
)

var shortenURLSvcStruct ShortenURLService
var shortenURLServiceOnce sync.Once

type ShortenURLService interface {
	ValidateRequest(ctx context.Context, req *models.ShortLinkRedisData) (
		validationErrors error)
	ProcessRequest(ctx context.Context, req *models.ShortLinkRedisData) (*models.ShortLinkRedisData, error)
}

type shortenURLService struct {
	config *config.RedisConfig
}

// InitShortenURLService ...
func InitShortenURLService(config *config.RedisConfig) ShortenURLService {
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

func (service *shortenURLService) ValidateRequest(ctx context.Context, req *models.ShortLinkRedisData) (validationError error) {
	if req.LongURL == "" {
		return errors.New("empty long URL provided")
	}

	if len(req.ShortURL) != 0 {
		return errors.New("please do not provide the short-url")
	}

	if !utils.IsValidUrl(req.LongURL) {
		return errors.New("please enter a valid long-url")
	}
	return nil
}

func (service *shortenURLService) ProcessRequest(ctx context.Context, req *models.ShortLinkRedisData) (*models.ShortLinkRedisData, error) {
	// check if LongURL Exist

	longURL := req.LongURL
	shortURL, err := dao.GetShortURL(ctx, longURL)
	if err != nil {
		return nil, err
	}

	if len(shortURL) != 0 {
		fmt.Println("Successfully found the short-link for the given long-link.")
		return &models.ShortLinkRedisData{
			LongURL:  longURL,
			ShortURL: shortURL,
		}, nil
	}

	// If shortURL does not exist generate and persist in DB
	shortURL, err = service.generateShortURL(longURL)
	if err != nil {
		fmt.Print("Service Error : failed to generate short-url")
		return nil, err
	}

	if len(shortURL) == 0 {
		return nil, errors.New("Service Error : failed to generate short-url")
	}

	err = dao.SetShortURL(ctx, shortURL, longURL)
	if err != nil {
		return nil, err
	}

	return &models.ShortLinkRedisData{
		ShortURL: shortURL,
		LongURL:  longURL,
	}, nil
}

func (service *shortenURLService) generateShortURL(longURL string) (string, error) {
	return "done", nil
}

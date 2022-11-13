package services

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"shortlink/internal/config"
	"shortlink/internal/dao"
	"shortlink/internal/models"
	"shortlink/internal/utils"
	"sync"
	"time"

	"github.com/go-redis/redis"
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
	if err != nil && err != redis.Nil {
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
	shortURL, err = service.generateShortURL(ctx, longURL)
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

func (service *shortenURLService) generateShortURL(ctx context.Context, longURL string) (string, error) {

	baseURL := "https://innocect/"
	shortSuffix := ""
	shortSuffixLength := 4

	for {
		shortSuffix = service.generate(shortSuffixLength)
		if Suffixexists, _ := dao.GetShortURL(ctx, longURL); len(Suffixexists) == 0 {
			break
		}
		shortSuffixLength = shortSuffixLength + 1
		if shortSuffixLength == 9 {
			fmt.Print("Service Error : Max number of retries atttempted. Try again")
			return "", errors.New("Service Error : Max number of retries atttempted. Try again")
		}
	}

	return baseURL + shortSuffix, nil
}

func (service *shortenURLService) generate(shortSuffixLength int) string {
	var characters = []rune("23456789abcdefghjkmnpqrtuvwxyzACDEFGHJKMNPQRTUVWXYZ")

	// Randomly shuffling characters
	rand.Shuffle(len(characters), func(i, j int) {
		characters[i], characters[j] = characters[j], characters[i]
	})

	shortSuffix := make([]rune, shortSuffixLength)
	newRandomSource := rand.NewSource(time.Now().UnixNano())
	randomNumberBasedOnSource := rand.New(newRandomSource)

	for i := range shortSuffix {
		shortSuffix[i] = characters[randomNumberBasedOnSource.Intn(len(characters))]
	}

	return string(shortSuffix)
}

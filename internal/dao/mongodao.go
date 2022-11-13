package dao

import (
	"context"
	"shortlink/internal/config"
	"shortlink/internal/models"
	"sync"
)

type ShortlinkMongoDAO interface {
	SetShortURL(ctx context.Context, longURL string, model *models.ShortLinkMongoData) error
}

var shortlinkMongoDAOStruct ShortlinkMongoDAO
var shortlinkMongoDAOOnce sync.Once

// shortlinkMongoDAO implements ShortlinkMongoDAO inteface and uses mongo
type shortlinkMongoDAO struct {
	config *config.RouterConfig
}

// InitShortlinkMongoDAO returns pointer to instance of ShortlinkMongoDAO implemenation
func InitShortlinkMongoDAO(conf *config.RouterConfig) ShortlinkMongoDAO {
	shortlinkMongoDAOOnce.Do(func() {
		shortlinkMongoDAOStruct = &shortlinkMongoDAO{config: conf}
	})
	return shortlinkMongoDAOStruct
}

// GetShortlinkMongoDAO ..
func GetShortlinkMongoDAO() ShortlinkMongoDAO {
	if shortlinkMongoDAOStruct == nil {
		panic("ShortlinkMongoDAO not initialized")
	}
	return shortlinkMongoDAOStruct
}

func (dao *shortlinkMongoDAO) SetShortURL(ctx context.Context, longURL string, model *models.ShortLinkMongoData) error {
	return nil
}

package global

import (
	"github.com/paiz4ever/go-web-framework/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG config.Config
	DB     *gorm.DB
	REDIS  redis.UniversalClient
	LOG    *zap.Logger
)

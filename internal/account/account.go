package account

import (
	"github.com/hatlonely/account/internal/mysqldb"
	"github.com/hatlonely/account/internal/rediscache"
	"github.com/sirupsen/logrus"
)

var InfoLog *logrus.Logger
var WarnLog *logrus.Logger
var AccessLog *logrus.Logger

func init() {
	InfoLog = logrus.New()
	WarnLog = logrus.New()
	AccessLog = logrus.New()
}

type Service struct {
	db    *mysqldb.MysqlDB
	cache *rediscache.RedisCache
}

func NewService(db *mysqldb.MysqlDB, cache *rediscache.RedisCache) *Service {
	return &Service{
		db:    db,
		cache: cache,
	}
}

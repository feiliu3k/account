package account

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hatlonely/account/internal/mysqldb"
	"github.com/hatlonely/account/internal/rediscache"
	"github.com/satori/go.uuid"
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

func (s *Service) Login(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")

	ok, token, err := s.login(username, password)
	res := gin.H{
		"valid": ok,
		"err":   err,
		"token": token,
	}

	AccessLog.WithFields(logrus.Fields{
		"request":  fmt.Sprintf("%v%v", c.Request.Host, c.Request.URL),
		"response": res,
	}).Info()

	c.JSON(200, res)
}

func (s *Service) login(username string, password string) (bool, string, error) {
	account, err := s.db.SelectAccountByUsernameOrTelephoneOrEmail(username)
	if err != nil {
		WarnLog.WithField("err", err).Warn("SelectAccountByUsernameOrTelephoneOrEmail failed")
		return false, "", err
	}

	if account.Password != password {
		return false, "", nil
	}

	uid := uuid.NewV4()
	buf := make([]byte, 32)
	hex.Encode(buf, uid.Bytes())
	token := string(buf)

	if err := s.cache.SetAccount(token, account); err != nil {
		return false, "", err
	}

	return true, token, nil
}

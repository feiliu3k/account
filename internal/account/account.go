package account

import (
	"encoding/hex"
	"github.com/hatlonely/account/internal/mysqldb"
	"github.com/hatlonely/account/internal/rediscache"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"regexp"
)

var InfoLog *logrus.Logger
var WarnLog *logrus.Logger
var AccessLog *logrus.Logger
var EmailRegex *regexp.Regexp
var TelephoneRegex *regexp.Regexp

func init() {
	InfoLog = logrus.New()
	WarnLog = logrus.New()
	AccessLog = logrus.New()

	TelephoneRegex = regexp.MustCompile(`^1[345789][0-9]{9}$`)
	EmailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
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

func NewToken() string {
	uid := uuid.NewV4()
	buf := make([]byte, 32)
	hex.Encode(buf, uid.Bytes())
	return string(buf)
}

func ValidateTelephone(telephone string) bool {
	return TelephoneRegex.MatchString(telephone)
}

func ValidateEmail(email string) bool {
	return EmailRegex.MatchString(email)
}

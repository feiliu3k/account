package account

import (
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LoginReqBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResBody struct {
	Valid bool   `json:"valid"`
	Token string `json:"token"`
	Err   error  `json:"err"`
}

func (s *Service) Login(c *gin.Context) {
	req := &LoginReqBody{}
	if err := json.NewDecoder(c.Request.Body).Decode(req); err != nil {
		WarnLog.WithField("err", err).Warn("decode request body failed")
		c.String(http.StatusBadRequest, "")
		return
	}
	res, err := s.login(req)
	if err != nil {
		WarnLog.WithField("err", err).Warn("login failed")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	AccessLog.WithFields(logrus.Fields{
		"host": c.Request.Host,
		"url":  c.Request.URL.String(),
		"req":  req,
		"res":  res,
	}).Info()

	c.JSON(http.StatusOK, res)
}

func (s *Service) login(req *LoginReqBody) (*LoginResBody, error) {
	account, err := s.db.SelectAccountByUsernameOrTelephoneOrEmail(req.Username)
	if err != nil {
		return nil, err
	}

	if account.Password != req.Password {
		return &LoginResBody{Valid: false}, nil
	}

	uid := uuid.NewV4()
	buf := make([]byte, 32)
	hex.Encode(buf, uid.Bytes())
	token := string(buf)

	if err := s.cache.SetAccount(token, account); err != nil {
		return nil, err
	}

	return &LoginResBody{Valid: true, Token: token}, nil
}

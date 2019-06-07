package account

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LoginReqBody struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResBody struct {
	Valid bool   `json:"valid,omitempty"`
	Token string `json:"token,omitempty"`
}

func (s *Service) Login(c *gin.Context) {
	rid := c.DefaultQuery("rid", NewToken())
	req := &LoginReqBody{}
	var res *LoginResBody
	var err error
	var buf []byte
	status := http.StatusOK

	defer func() {
		AccessLog.WithFields(logrus.Fields{
			"host":   c.Request.Host,
			"body":   string(buf),
			"url":    c.Request.URL.String(),
			"req":    req,
			"res":    res,
			"rid":    rid,
			"err":    err,
			"status": status,
		}).Info()
	}()

	buf, err = c.GetRawData()
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("get raw data failed")
		status = http.StatusBadRequest
		c.String(status, "")
		return
	}

	if err = json.Unmarshal(buf, req); err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("decode request body failed")
		status = http.StatusBadRequest
		c.String(status, "")
		return
	}

	res, err = s.login(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("login failed")
		status = http.StatusInternalServerError
		c.String(status, err.Error())
		return
	}

	status = http.StatusOK
	c.JSON(status, res)
}

func (s *Service) login(req *LoginReqBody) (*LoginResBody, error) {
	account, err := s.db.SelectAccountByUsernameOrTelephoneOrEmail(req.Username)
	if err != nil {
		return nil, err
	}

	if account.Password != req.Password {
		return &LoginResBody{Valid: false}, nil
	}

	token := NewToken()
	if err := s.cache.SetAccount(token, account); err != nil {
		return nil, err
	}

	return &LoginResBody{Valid: true, Token: token}, nil
}

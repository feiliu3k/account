package account

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LoginReqBody struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResBody struct {
	Valid bool   `json:"valid"`
	Token string `json:"token"`
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
		err = fmt.Errorf("get raw data failed, err: [%v]", err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	if err = json.Unmarshal(buf, req); err != nil {
		err = fmt.Errorf("json unmarshal body failed. body: [%v], err: [%v]", string(buf), err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	if req.Username == "" || req.Password == "" {
		err = fmt.Errorf("username or password is empty")
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, "username or password is empty")
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

	if account == nil {
		return &LoginResBody{Valid: false}, nil
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

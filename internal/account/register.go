package account

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hatlonely/account/internal/mysqldb"
	"github.com/sirupsen/logrus"
	"net/http"
)

type RegisterReqBody struct {
	Username  string `json:"username,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

type RegisterResBody struct {
	Success bool `json:"success,omitempty"`
}

func (s *Service) Register(c *gin.Context) {
	rid := c.DefaultQuery("rid", NewToken())
	req := &RegisterReqBody{}
	var res *RegisterResBody
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

	res, err = s.register(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("login failed")
		status = http.StatusInternalServerError
		c.String(status, err.Error())
		return
	}

	status = http.StatusOK
	c.JSON(status, res)
}

func (s *Service) register(req *RegisterReqBody) (*RegisterResBody, error) {
	ok, err := s.db.InsertAccount(&mysqldb.Account{
		Username:  req.Username,
		Telephone: req.Telephone,
		Email:     req.Email,
		Password:  req.Password,
	})

	return &RegisterResBody{Success: ok}, err
}

package account

import (
	"github.com/gin-gonic/gin"
	"github.com/hatlonely/account/internal/mysqldb"
)

func (s *Service) Register(c *gin.Context) {

}

func (s *Service) register(username string, telephone string, email string, password string) (bool, error) {
	ok, err := s.db.InsertAccount(&mysqldb.Account{
		Username:  username,
		Telephone: telephone,
		Email:     email,
		Password:  password,
	})

	return ok, err
}

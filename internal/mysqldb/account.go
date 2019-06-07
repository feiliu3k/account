package mysqldb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type Account struct {
	ID        int    `gorm:"type:bigint(20) auto_increment;primary_key" json:"id"`
	Username  string `gorm:"type:varchar(64);not null;unique_index:username_idx" json:"username"`
	Telephone string `gorm:"type:varchar(64) default '';not null;index:telephone_idx" json:"telephone"`
	Email     string `gorm:"type:varchar(64) default '';not null;index:email_idx" json:"email"`
	Password  string `gorm:"type:varchar(32) default '';not null" json:"password"`
	Role      int    `gorm:"type:bigint(20) default 0;not null" json:"role"`
}

func (m *MysqlDB) SelectAccountByUsernameOrTelephoneOrEmail(key string) (*Account, error) {
	account := &Account{}
	if key == "" {
		return nil, fmt.Errorf("account key is null")
	}
	if err := m.db.Where("username=?", key).Or("telephone=?", key).Or("email=?", key).First(account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return account, nil
}

func (m *MysqlDB) InsertAccount(account *Account) (bool, error) {
	if account.Username == "" {
		return false, fmt.Errorf("username is null, account [%#v]", account)
	}
	if account.Email == "" && account.Telephone == "" {
		return false, fmt.Errorf("email and telephone are both null, account [%#v]", account)
	}

	accountDB := &Account{}
	var conditions []string
	if account.ID != 0 {
		conditions = append(conditions, fmt.Sprintf("id=%v", account.ID))
	}
	if account.Username != "" {
		conditions = append(conditions, fmt.Sprintf("username='%v'", account.Username))
	}
	if account.Telephone != "" {
		conditions = append(conditions, fmt.Sprintf("telephone='%v'", account.Telephone))
	}
	if account.Email != "" {
		conditions = append(conditions, fmt.Sprintf("email='%v'", account.Email))
	}
	err := m.db.Where(strings.Join(conditions, " OR ")).First(accountDB).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if err != gorm.ErrRecordNotFound {
		if accountDB.ID == account.ID {
			return false, fmt.Errorf("accountID [%v] is already exists", accountDB.ID)
		}
		if accountDB.Username == account.Username {

			return false, fmt.Errorf("username [%v] is already exists", accountDB.Username)
		}
		if accountDB.Telephone == account.Telephone {
			return false, fmt.Errorf("telephone [%v] is already exists", accountDB.Telephone)
		}
		if accountDB.Email == account.Email {
			return false, fmt.Errorf("email [%v] is already exists", accountDB.Email)
		}
	}

	if err := m.db.Create(account).Error; err != nil {
		return false, err
	}

	return true, nil
}

package mysqldb

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMysqlDB_InsertAccount(t *testing.T) {
	m, err := NewMysqlDB("hatlonely:keaiduo1@/hads?charset=utf8&parseTime=True&loc=Local")
	Convey("test mysqldb insert account", t, func() {
		So(err, ShouldBeNil)
		So(m, ShouldNotBeNil)

		m.db.Where("username='hatlonely1'").
			Or("email='hatlonely1@foxmail.com'").
			Or("telephone='13145678901'").Delete(&Account{})

		Convey("insert account without username", func() {
			ok, err := m.InsertAccount(&Account{
				Email:     "hatlonely1@foxmail.com",
				Telephone: "13145678901",
				Password:  "e010597fcf126d58fdfa36e636f8fc9e",
			})
			So(err, ShouldNotBeNil)
			So(ok, ShouldBeFalse)
		})

		Convey("insert account without email and telephone", func() {
			ok, err := m.InsertAccount(&Account{
				Username: "hatlonely1",
				Password: "e010597fcf126d58fdfa36e636f8fc9e",
			})
			So(err, ShouldNotBeNil)
			So(ok, ShouldBeFalse)
		})

		Convey("insert account with username and email", func() {
			ok, err := m.InsertAccount(&Account{
				Username: "hatlonely1",
				Email:    "hatlonely1@foxmail.com",
				Password: "e010597fcf126d58fdfa36e636f8fc9e",
			})
			So(err, ShouldBeNil)
			So(ok, ShouldBeTrue)

			account, err := m.SelectAccountByUsernameOrTelephoneOrEmail("hatlonely1")
			So(err, ShouldBeNil)
			So(account.Username, ShouldEqual, "hatlonely1")
			So(account.Email, ShouldEqual, "hatlonely1@foxmail.com")
			So(account.Telephone, ShouldEqual, "")
			So(account.Password, ShouldEqual, "e010597fcf126d58fdfa36e636f8fc9e")
		})

		Convey("insert account with username and telephone", func() {
			ok, err := m.InsertAccount(&Account{
				Username:  "hatlonely1",
				Telephone: "13145678901",
				Password:  "e010597fcf126d58fdfa36e636f8fc9e",
			})
			So(err, ShouldBeNil)
			So(ok, ShouldBeTrue)

			account, err := m.SelectAccountByUsernameOrTelephoneOrEmail("hatlonely1")
			So(err, ShouldBeNil)
			So(account.Username, ShouldEqual, "hatlonely1")
			So(account.Email, ShouldEqual, "")
			So(account.Telephone, ShouldEqual, "13145678901")
			So(account.Password, ShouldEqual, "e010597fcf126d58fdfa36e636f8fc9e")
		})

		Convey("insert account with username and telephone and email", func() {
			ok, err := m.InsertAccount(&Account{
				Username:  "hatlonely1",
				Email:     "hatlonely1@foxmail.com",
				Telephone: "13145678901",
				Password:  "e010597fcf126d58fdfa36e636f8fc9e",
			})
			So(err, ShouldBeNil)
			So(ok, ShouldBeTrue)

			account, err := m.SelectAccountByUsernameOrTelephoneOrEmail("hatlonely1")
			So(err, ShouldBeNil)
			So(account.Username, ShouldEqual, "hatlonely1")
			So(account.Email, ShouldEqual, "hatlonely1@foxmail.com")
			So(account.Telephone, ShouldEqual, "13145678901")
			So(account.Password, ShouldEqual, "e010597fcf126d58fdfa36e636f8fc9e")
		})

		Convey("insert account dup usename or email or telephone", func() {
			ok, err := m.InsertAccount(&Account{
				Username:  "hatlonely1",
				Email:     "hatlonely1@foxmail.com",
				Telephone: "13145678901",
				Password:  "e010597fcf126d58fdfa36e636f8fc9e",
			})
			So(err, ShouldBeNil)
			So(ok, ShouldBeTrue)

			Convey("insert account dup username", func() {
				ok, err = m.InsertAccount(&Account{
					Username:  "hatlonely1",
					Email:     "hatlonely12@foxmail.com",
					Telephone: "13145678902",
					Password:  "e010597fcf126d58fdfa36e636f8fc9e",
				})
				So(err, ShouldNotBeNil)
				So(ok, ShouldBeFalse)
			})

			Convey("insert account dup telephone", func() {
				ok, err = m.InsertAccount(&Account{
					Username:  "hatlonely12",
					Email:     "hatlonely12@foxmail.com",
					Telephone: "13145678901",
					Password:  "e010597fcf126d58fdfa36e636f8fc9e",
				})
				So(err, ShouldNotBeNil)
				So(ok, ShouldBeFalse)
			})

			Convey("insert account dup email", func() {
				ok, err = m.InsertAccount(&Account{
					Username:  "hatlonely12",
					Email:     "hatlonely1@foxmail.com",
					Telephone: "+8612345678902",
					Password:  "e010597fcf126d58fdfa36e636f8fc9e",
				})
				So(err, ShouldNotBeNil)
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestMysqlDB_SelectAccountByUsernameOrTelephoneOrEmail(t *testing.T) {
	m, err := NewMysqlDB("hatlonely1:keaiduo1@/hads?charset=utf8&parseTime=True&loc=Local")
	Convey("test mysqldb select account by username or telephone or email", t, func() {
		So(err, ShouldBeNil)
		So(m, ShouldNotBeNil)

		m.db.Where("username='hatlonely1'").
			Or("email='hatlonely1@foxmail.com'").
			Or("telephone='13145678901'").Delete(&Account{})

		Convey("select account use empty key", func() {
			account, err := m.SelectAccountByUsernameOrTelephoneOrEmail("")
			So(account, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})

		Convey("select account use nonexists key", func() {
			account, err := m.SelectAccountByUsernameOrTelephoneOrEmail("hatlonely1")
			So(account, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		ok, err := m.InsertAccount(&Account{
			Username:  "hatlonely1",
			Email:     "hatlonely1@foxmail.com",
			Telephone: "13145678901",
			Password:  "e010597fcf126d58fdfa36e636f8fc9e",
		})
		So(err, ShouldBeNil)
		So(ok, ShouldBeTrue)

		Convey("select account by username", func() {
			account, err := m.SelectAccountByUsernameOrTelephoneOrEmail("hatlonely1")
			So(err, ShouldBeNil)
			So(account.Username, ShouldEqual, "hatlonely1")
			So(account.Email, ShouldEqual, "hatlonely1@foxmail.com")
			So(account.Telephone, ShouldEqual, "13145678901")
			So(account.Password, ShouldEqual, "e010597fcf126d58fdfa36e636f8fc9e")
		})

		Convey("select account by telephone", func() {
			account, err := m.SelectAccountByUsernameOrTelephoneOrEmail("hatlonely1@foxmail.com")
			So(err, ShouldBeNil)
			So(account.Username, ShouldEqual, "hatlonely1")
			So(account.Email, ShouldEqual, "hatlonely1@foxmail.com")
			So(account.Telephone, ShouldEqual, "13145678901")
			So(account.Password, ShouldEqual, "e010597fcf126d58fdfa36e636f8fc9e")
		})

		Convey("select account by email", func() {
			account, err := m.SelectAccountByUsernameOrTelephoneOrEmail("13145678901")
			So(err, ShouldBeNil)
			So(account.Username, ShouldEqual, "hatlonely1")
			So(account.Email, ShouldEqual, "hatlonely1@foxmail.com")
			So(account.Telephone, ShouldEqual, "13145678901")
			So(account.Password, ShouldEqual, "e010597fcf126d58fdfa36e636f8fc9e")
		})
	})
}

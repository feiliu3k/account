module github.com/hatlonely/account

go 1.12

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.0.0-20190603211518-c8433c9aaceb
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190602015325-4c4f7f33c9ed
	golang.org/x/text => github.com/golang/text v0.3.2
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190605220351-eb0b1bdb6ae6
// google.golang.org/api => github.com/googleapis/google-api-go-client master
)

require (
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.4.0
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jinzhu/gorm v1.9.8
	github.com/lestrrat-go/envload v0.0.0-20180220234015-a3eb8ddeffcc // indirect
	github.com/lestrrat-go/file-rotatelogs v2.2.0+incompatible
	github.com/lestrrat-go/strftime v0.0.0-20180821113735-8b31f9c59b0f // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	github.com/spf13/viper v1.4.0
	github.com/tebeka/strftime v0.0.0-20140926081919-3f9c7761e312 // indirect
)

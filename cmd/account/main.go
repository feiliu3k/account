package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hatlonely/account/internal/account"
	"github.com/hatlonely/account/internal/logger"
	"github.com/hatlonely/account/internal/mysqldb"
	"github.com/spf13/viper"
	"os"
)

// AppVersion name
var AppVersion = "unknown"

func main() {
	version := flag.Bool("v", false, "print current version")
	configfile := flag.String("c", "configs/account.json", "config file path")
	flag.Parse()
	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	config := viper.New()
	config.SetConfigType("json")
	fp, err := os.Open(*configfile)
	if err != nil {
		panic(err)
	}
	err = config.ReadConfig(fp)
	if err != nil {
		panic(err)
	}

	infoLog, err := logger.NewTextLoggerWithViper(config.Sub("logger.infoLog"))
	if err != nil {
		panic(err)
	}
	warnLog, err := logger.NewTextLoggerWithViper(config.Sub("logger.warnLog"))
	if err != nil {
		panic(err)
	}
	accessLog, err := logger.NewJsonLoggerWithViper(config.Sub("logger.accessLog"))
	account.InfoLog = infoLog
	account.WarnLog = warnLog
	account.AccessLog = accessLog

	infoLog.Infof("%v init success, port[%v]", os.Args[0], config.GetString("service.port"))

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	db, err := mysqldb.NewMysqlDB(config.GetString("mysqldb.uri"))
	if err != nil {
		panic(err)
	}
	service := account.NewService(db)

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})
	r.GET("/login", service.Login)

	if err := r.Run(config.GetString("service.port")); err != nil {
		panic(err)
	}
}

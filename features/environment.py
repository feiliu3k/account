#!/usr/bin/env python3

import pymysql
import redis
import subprocess
import time
import requests
import datetime
import json


config = {
    "port": 6061,
    "prefix": "build/account",
    "mysqldb": {
        "host": "127.0.0.1",
        "port": 3306,
        "user": "hatlonely",
        "password": "keaiduo1",
        "db": "hads"
    },
    "redis": {
        "host": "127.0.0.1",
        "port": 6379
    }
}


def deploy():
    fp = open("{}/configs/account.json".format(config["prefix"]))
    cf = json.loads(fp.read())
    fp.close()
    cf["mysqldb"]["uri"] = "{user}:{password}@tcp({host}:{port})/{db}?charset=utf8&parseTime=True&loc=Local".format(
        user=config["mysqldb"]["user"],
        password=config["mysqldb"]["password"],
        db=config["mysqldb"]["db"],
        host=config["mysqldb"]["host"],
        port=config["mysqldb"]["port"],
    )
    cf["rediscache"]["address"] = "{host}:{port}".format(
        host=config["redis"]["host"],
        port=config["redis"]["port"],
    )
    cf["service"]["port"] = ":{}".format(config["port"])
    fp = open("{}/configs/account.json".format(config["prefix"]), "w")
    fp.write(json.dumps(cf, indent=4))
    fp.close()


def start():
    subprocess.Popen(
        "cd {} && nohup bin/account &".format(config["prefix"]),  shell=True
    )

    now = datetime.datetime.now()
    while datetime.datetime.now() - now < datetime.timedelta(seconds=5):
        try:
            res = requests.get("{}/ping".format(config["url"]))
            if res.status_code == 200:
                break
        except Exception as e:
            time.sleep(0.1)


def stop():
    subprocess.getstatusoutput(
        "ps aux | grep bin/account | grep -v grep | awk '{print $2}' | xargs kill"
    )


def before_all(context):
    config["url"] = "http://127.0.0.1:{}".format(config["port"])
    context.mysql_conn = pymysql.connect(
        host=config["mysqldb"]["host"],
        user=config["mysqldb"]["user"],
        port=config["mysqldb"]["port"],
        password=config["mysqldb"]["password"],
        db=config["mysqldb"]["db"],
        charset="utf8",
        cursorclass=pymysql.cursors.DictCursor
    )
    context.redis_client = redis.Redis(host="localhost", port=6379, db=0)
    deploy()
    start()
    context.config = config


def after_all(context):
    stop()


def after_scenario(context, scenario):
    if not context.cleanup:
        return

    if "sql" in context.cleanup:
        with context.mysql_conn.cursor() as cursor:
            cursor.execute(context.cleanup["sql"])
        context.mysql_conn.commit()

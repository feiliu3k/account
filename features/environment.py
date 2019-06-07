#!/usr/bin/env python3

import pymysql
import redis
import subprocess
import time
import requests
import datetime


host = '127.0.0.1'
user = 'hatlonely'
password = 'keaiduo1'
db = 'hads'

url = "http://127.0.0.1:6060"


def start():
    subprocess.Popen("cd build/account && nohup bin/account &",  shell=True)

    now = datetime.datetime.now()
    while datetime.datetime.now() - now < datetime.timedelta(seconds=5):
        try:
            res = requests.get(url+'/ping')
            if res.status_code == 200:
                break
        except Exception as e:
            time.sleep(0.1)


def stop():
    subprocess.getstatusoutput(
        "ps aux | grep bin/account | grep -v grep | awk '{print $2}' | xargs kill"
    )


def before_all(context):
    context.mysql_conn = pymysql.connect(
        host=host,
        user=user,
        password=password,
        db=db,
        charset='utf8',
        cursorclass=pymysql.cursors.DictCursor
    )
    context.redis_client = redis.Redis(host="localhost", port=6379, db=0)
    start()


def after_all(context):
    stop()


def after_scenario(context, scenario):
    if not context.cleanup:
        return

    if "sql" in context.cleanup:
        with context.mysql_conn.cursor() as cursor:
            cursor.execute(context.cleanup["sql"])
        context.mysql_conn.commit()

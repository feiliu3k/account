#!/usr/bin/env python3

import pymysql
import redis


host = '127.0.0.1'
user = 'hatlonely'
password = 'keaiduo1'
db = 'hads'


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


def after_scenario(context, scenario):
    if not context.cleanup:
        return

    if "sql" in context.cleanup:
        with context.mysql_conn.cursor() as cursor:
            cursor.execute(context.cleanup["sql"])
        context.mysql_conn.commit()

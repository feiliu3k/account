#!/usr/bin/env python3

import pymysql


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


def after_scenario(context, scenario):
    if not context.cleanup:
        return

    if "sql" in context.cleanup:
        with context.mysql_conn.cursor() as cursor:
            cursor.execute(context.cleanup["sql"])
        context.mysql_conn.commit()

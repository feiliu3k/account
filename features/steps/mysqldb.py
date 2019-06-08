#!/usr/bin/env python3

from behave import *


@given('创建用户 "{username:str}", 电话 "{telephone:str}", 邮箱 "{email:str}", 密码 "{password:str}"')
def step_impl(context, username, telephone, email, password):
    context.cleanup = {
        "sql": "DELETE FROM accounts WHERE username='{}' OR telephone='{}' OR email='{}'".format(
            username, telephone, email
        )
    }
    context.username = username
    context.telephone = telephone
    context.email = email
    context.password = password
    with context.mysql_conn.cursor() as cursor:
        cursor.execute(context.cleanup["sql"])
        cursor.execute(
            "INSERT INTO accounts (username, telephone, email, password) VALUES (%s, %s, %s, %s)",
            (username, telephone, email, password)
        )
    context.mysql_conn.commit()


@given('删除用户 "{username:str}"')
def step_impl(context, username):
    context.cleanup = {
        "sql": "DELETE FROM accounts WHERE username='{}'".format(
            username
        )
    }
    with context.mysql_conn.cursor() as cursor:
        cursor.execute(context.cleanup["sql"])
    context.mysql_conn.commit()

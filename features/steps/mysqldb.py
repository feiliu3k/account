#!/usr/bin/env python3

from behave import *
from hamcrest import *


@given('mysqldb.accounts 创建用户 username: "{username:str}", telephone: "{telephone:str}", email: "{email:str}", password: "{password:str}"')
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


@given('mysqldb.accounts 删除用户 username: "{username:str}"')
def step_impl(context, username):
    context.cleanup = {
        "sql": "DELETE FROM accounts WHERE username='{}'".format(
            username
        )
    }
    with context.mysql_conn.cursor() as cursor:
        cursor.execute(context.cleanup["sql"])
    context.mysql_conn.commit()


@then('检查 mysqldb.accounts，存在记录 username: "{username:str}", telephone: "{telephone:str}", email: "{email:str}", password: "{password:str}"')
def step_impl(context, username, telephone, email, password):
    with context.mysql_conn.cursor() as cursor:
        cursor.execute(
            "SELECT * FROM accounts WHERE username='{}'".format(username)
        )
        res = cursor.fetchall()
        assert_that(len(res), equal_to(1))
        account = res[0]
        assert_that(username, equal_to(account["username"]))
        assert_that(telephone, equal_to(account["telephone"]))
        assert_that(email, equal_to(account["email"]))
        assert_that(password, equal_to(account["password"]))

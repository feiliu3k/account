#!/usr/bin/env python3

from behave import *
from hamcrest import *
import requests
import json


@when('用户 "{username:str}" 使用密码 "{password:str}" 登陆')
def step_impl(context, username, password):
    res = requests.post("{}/login".format(context.config["url"]), json={
        "username": username,
        "password": password,
    })
    context.status = res.status_code
    if context.status == 200:
        context.res = json.loads(res.text)
    else:
        context.res = res.text
    print({
        "status": context.status,
        "res": context.res,
    })


@then('检查状态码应该为 "{status:int}"')
def step_impl(context, status):
    assert_that(context.status, equal_to(status))


@then('检查有效性应该为 "{valid:bool}"')
def step_impl(context, valid):
    assert_that(context.res["valid"], equal_to(valid))


@then('检查token长度应该为 "{tokenlen:int}"')
def step_impl(context, tokenlen):
    assert_that(len(context.res["token"]), equal_to(tokenlen))


@then('检查redis中token')
def step_impl(context):
    res = context.redis_client.get(context.res["token"])
    account = json.loads(res)
    assert_that(context.username, equal_to(account["username"]))
    assert_that(context.telephone, equal_to(account["telephone"]))
    assert_that(context.email, equal_to(account["email"]))
    assert_that(context.password, equal_to(account["password"]))

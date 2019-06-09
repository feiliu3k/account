#!/usr/bin/env python3

from behave import *
from hamcrest import *
import requests
import json


@when('请求 /login, username: "{username:str}", password: "{password:str}"')
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


@then('检查登陆返回包体 res.body, valid: {valid:bool}, tokenlen: {tokenlen:int}')
def step_impl(context, valid, tokenlen):
    assert_that(context.res["valid"], equal_to(valid))
    assert_that(len(context.res["token"]), equal_to(tokenlen))

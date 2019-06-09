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


@then('检查有效性应该为 "{valid:bool}"')
def step_impl(context, valid):
    assert_that(context.res["valid"], equal_to(valid))


@then('检查token长度应该为 "{tokenlen:int}"')
def step_impl(context, tokenlen):
    assert_that(len(context.res["token"]), equal_to(tokenlen))

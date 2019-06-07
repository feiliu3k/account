#!/usr/bin/env python3

from behave import *
from hamcrest import *
import requests
import json


register_type(int=int)
register_type(str=lambda x: x if x != "N/A" else "")
register_type(bool=lambda x: True if x == "true" else False)


@given('创建用户 "{username:str}", 电话 "{telephone:str}", 邮箱 "{email:str}", 密码 "{password:str}"')
def step_impl(context, username, telephone, email, password):
    print(username, telephone, email, password)


@when('用户 "{username:str}" 使用密码 "{password:str}" 登陆')
def step_impl(context, username, password):
    res = requests.post("http://127.0.0.1:6060/login", json={
        "username": username,
        "password": password,
    })
    context.status = res.status_code
    if context.status == 200:
        context.res = json.loads(res.text)
    else:
        context.res = res.text
    print({
        "status": res.status_code,
        "res": context.res,
    })


@then('检查返回 "{status:int}" 登陆状态 "{valid:bool}"')
def step_impl(context, status, valid):
    assert_that(context.status, equal_to(status))
    if status == 200:
        assert_that(context.res["valid"], equal_to(valid))
        if valid:
            assert_that(context.res["token"], is_not(""))
            assert_that(len(context.res["token"]), equal_to(32))
        else:
            assert_that(context.res["token"], equal_to(""))

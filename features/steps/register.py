#!/usr/bin/env python3

from behave import *
from hamcrest import *
import requests
import json


@when('请求 /register, username: "{username:str}", telephone: "{telephone:str}", email: "{email:str}", password: "{password:str}"')
def step_impl(context, username, telephone, email, password):
    context.username = username
    context.telephone = telephone
    context.email = email
    context.password = password
    res = requests.post("{}/register".format(context.config["url"]), json={
        "username": username,
        "telephone": telephone,
        "email": email,
        "password": password,
    })
    context.status = res.status_code
    context.body = res.text
    if context.status == 200:
        context.res = json.loads(res.text)
    print({
        "status": context.status,
        "body": context.body,
    })


@then('检查注册返回包体 res.body, success: {success:bool}')
def step_impl(context, success):
    assert_that(context.res["success"], equal_to(success))

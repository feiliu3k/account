#!/usr/bin/env python3

from behave import *
from hamcrest import *
import requests
import json


@when('注册用户 "{username:str}", 电话 "{telephone:str}", 邮箱 "{email:str}", 密码 "{password:str}"')
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
    if context.status == 200:
        context.res = json.loads(res.text)
    else:
        context.res = res.text
    print({
        "status": context.status,
        "res": context.res,
    })

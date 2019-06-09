#!/usr/bin/env python3

from behave import *
from hamcrest import *
import json


@then('检查 rediscache.token')
def step_impl(context):
    res = context.redis_client.get(context.res["token"])
    account = json.loads(res)
    assert_that(context.username, equal_to(account["username"]))
    assert_that(context.telephone, equal_to(account["telephone"]))
    assert_that(context.email, equal_to(account["email"]))
    assert_that(context.password, equal_to(account["password"]))

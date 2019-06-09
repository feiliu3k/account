#!/usr/bin/env python3

from behave import *
from hamcrest import *


@then('检查状态码 res.status_code: {status:int}')
def step_impl(context, status):
    assert_that(context.status, equal_to(status))

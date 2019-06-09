#!/usr/bin/env python3

from behave import *
from hamcrest import *


@then('检查状态码应该为 "{status:int}"')
def step_impl(context, status):
    assert_that(context.status, equal_to(status))

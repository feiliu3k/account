Feature: register 注册测试

    Scenario Outline: 注册成功
        Given 删除用户 "<username>"
        When 注册用户 "<username>", 电话 "<telephone>", 邮箱 "<email>", 密码 "<password>"
        Then 检查状态码应该为 "<status>"
        Examples:
            | username   | telephone      | email                  | password                         | status |
            | hatlonely1 | +8612345678901 | hatlonely1@foxmail.com | e010597fcf126d58fdfa36e636f8fc9e | 200    |
            | hatlonely2 | N/A            | hatlonely2@foxmail.com | 0fe808594e47df1a336bafd8ab32f326 | 200    |
            | hatlonely3 | +8612345678903 | N/A                    | de9baf2c5dde96f0a8b371117e936d4b | 200    |

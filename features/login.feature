Feature: login 登陆测试

    Scenario Outline: Username/Telephone/Email 登陆
        Given 创建用户 "hatlonely", 电话 "+8612345678901", 邮箱 "hatlonely@foxmail.com", 密码 "e010597fcf126d58fdfa36e636f8fc9e"
        When 用户 "<username>" 使用密码 "<password>" 登陆
        Then 检查返回 "<status>" 登陆状态 "<valid>"
        Examples:
            | username              | password                         | status | valid |
            | hatlonely             | e010597fcf126d58fdfa36e636f8fc9e | 200    | true  |
            | +8612345678901        | e010597fcf126d58fdfa36e636f8fc9e | 200    | true  |
            | hatlonely@foxmail.com | e010597fcf126d58fdfa36e636f8fc9e | 200    | true  |
            | notexistsuser         | e010597fcf126d58fdfa36e636f8fc9e | 200    | false |
            | hatlonely             | wrong_password                   | 200    | false |

    Scenario Outline: 异常登陆
        Given 创建用户 "hatlonely", 电话 "+8612345678901", 邮箱 "hatlonely@foxmail.com", 密码 "e010597fcf126d58fdfa36e636f8fc9e"
        When 用户 "<username>" 使用密码 "<password>" 登陆
        Then 检查返回 "<status>" 登陆状态 "<valid>"
        Examples:
            | username | password                         | status | valid |
            | N/A      | e010597fcf126d58fdfa36e636f8fc9e | 400    | false |
            | N/A      | N/A                              | 400    | false |

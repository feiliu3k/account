Feature: login 登陆测试

    Scenario Outline: 登陆成功
        Given 创建用户 "hatlonely", 电话 "+8612345678901", 邮箱 "hatlonely@foxmail.com", 密码 "e010597fcf126d58fdfa36e636f8fc9e"
        When 用户 "<username>" 使用密码 "<password>" 登陆
        Then 检查状态码 res.status_code: <status>
        Then 检查登陆返回包体 res.body, valid: <valid>, tokenlen: <tokenlen>
        Then 检查 rediscache.token
        Examples:
            | username              | password                         | status | valid | tokenlen |
            | hatlonely             | e010597fcf126d58fdfa36e636f8fc9e | 200    | true  | 32       |
            | +8612345678901        | e010597fcf126d58fdfa36e636f8fc9e | 200    | true  | 32       |
            | hatlonely@foxmail.com | e010597fcf126d58fdfa36e636f8fc9e | 200    | true  | 32       |

    Scenario Outline: 登陆失败
        Given 创建用户 "hatlonely", 电话 "+8612345678901", 邮箱 "hatlonely@foxmail.com", 密码 "e010597fcf126d58fdfa36e636f8fc9e"
        When 用户 "<username>" 使用密码 "<password>" 登陆
        Then 检查状态码 res.status_code: <status>
        Then 检查登陆返回包体 res.body, valid: <valid>, tokenlen: <tokenlen>
        Examples:
            | username      | password                         | status | valid | tokenlen |
            | notexistsuser | e010597fcf126d58fdfa36e636f8fc9e | 200    | false | 0        |
            | hatlonely     | wrong_password                   | 200    | false | 0        |

    Scenario Outline: 异常登陆
        Given 创建用户 "hatlonely", 电话 "+8612345678901", 邮箱 "hatlonely@foxmail.com", 密码 "e010597fcf126d58fdfa36e636f8fc9e"
        When 用户 "<username>" 使用密码 "<password>" 登陆
        Then 检查状态码 res.status_code: <status>
        Examples:
            | username | password                         | status |
            | N/A      | e010597fcf126d58fdfa36e636f8fc9e | 400    |
            | N/A      | N/A                              | 400    |

Feature: login 登陆测试

    Scenario Outline: 登陆成功
        Given mysqldb.accounts 创建用户 username: "hatlonely", telephone: "12345678901", email: "hatlonely@foxmail.com", password: "e010597fcf126d58fdfa36e636f8fc9e"
        When 请求 /login, username: "<username>", password: "<password>"
        Then 检查状态码 res.status_code: <status>
        Then 检查登陆返回包体 res.body, valid: <valid>, tokenlen: <tokenlen>
        Then 检查 rediscache.token
        Examples:
            | username              | password                         | status | valid | tokenlen |
            | hatlonely             | e010597fcf126d58fdfa36e636f8fc9e | 200    | true  | 32       |
            | 12345678901           | e010597fcf126d58fdfa36e636f8fc9e | 200    | true  | 32       |
            | hatlonely@foxmail.com | e010597fcf126d58fdfa36e636f8fc9e | 200    | true  | 32       |

    Scenario Outline: 登陆失败
        Given mysqldb.accounts 创建用户 username: "hatlonely", telephone: "12345678901", email: "hatlonely@foxmail.com", password: "e010597fcf126d58fdfa36e636f8fc9e"
        When 请求 /login, username: "<username>", password: "<password>"
        Then 检查状态码 res.status_code: <status>
        Then 检查登陆返回包体 res.body, valid: <valid>, tokenlen: <tokenlen>
        Examples:
            | username      | password                         | status | valid | tokenlen |
            | notexistsuser | e010597fcf126d58fdfa36e636f8fc9e | 200    | false | 0        |
            | hatlonely     | wrong_password                   | 200    | false | 0        |

    Scenario Outline: 异常登陆
        Given mysqldb.accounts 创建用户 username: "hatlonely", telephone: "12345678901", email: "hatlonely@foxmail.com", password: "e010597fcf126d58fdfa36e636f8fc9e"
        When 请求 /login, username: "<username>", password: "<password>"
        Then 检查状态码 res.status_code: <status>
        Then 检查返回包体 res.body，包含字符串 <body>
        Examples:
            | username | password                         | status | body                          |
            | N/A      | e010597fcf126d58fdfa36e636f8fc9e | 400    | username or password is empty |
            | N/A      | N/A                              | 400    | username or password is empty |

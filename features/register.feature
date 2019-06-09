Feature: register 注册测试

    Scenario Outline: 注册成功
        Given mysqldb.accounts 删除用户 username: "<username>"
        When 请求 /register, username: "<username>", telephone: "<telephone>", email: "<email>", password: "<password>"
        Then 检查状态码 res.status_code: <status>
        Then 检查注册返回包体 res.body, success: <success>
        Then 检查 mysqldb.accounts，存在记录 username: "<username>", telephone: "<telephone>", email: "<email>", password: "<password>"
        Examples:
            | username   | telephone   | email                  | password                         | status | success |
            | hatlonely1 | 12345678901 | hatlonely1@foxmail.com | e010597fcf126d58fdfa36e636f8fc9e | 200    | true    |
            | hatlonely2 | N/A         | hatlonely2@foxmail.com | 0fe808594e47df1a336bafd8ab32f326 | 200    | true    |
            | hatlonely3 | 12345678903 | N/A                    | de9baf2c5dde96f0a8b371117e936d4b | 200    | true    |

    Scenario: 重复注册
        Given mysqldb.accounts 删除用户 username: "<username>"
        Given mysqldb.accounts 创建用户 username: "hatlonely", telephone: "12345678901", email: "hatlonely@foxmail.com", password: "e010597fcf126d58fdfa36e636f8fc9e"
        When 请求 /register, username: "hatlonely", telephone: "12345678901", email: "hatlonely@foxmail.com", password: "e010597fcf126d58fdfa36e636f8fc9e"
        Then 检查状态码 res.status_code: 500
        Then 检查返回包体 res.body，包含字符串 "username [hatlonely] is already exists"
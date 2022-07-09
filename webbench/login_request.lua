wrk.method = "POST"
-- json string
wrk.body = '{"user": "xiaoming", "password": "123"}'
-- 设置content-type
wrk.headers["Content-Type"] = "application/json"



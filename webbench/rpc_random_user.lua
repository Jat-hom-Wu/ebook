#!/usr/bin/lua5.1


-- 随机化请求查询参数
request = function()
    local rand_idx = math.random(10000000, 10020000)
    -- local tracking_number = tns[rand_idx]
    local user_name = "stress_test_"..rand_idx
    local method = "POST"
    local path = "/ebook/rpclogin"
    local body = string.format('{"user": "%s", "password": "1234567890"}', user_name)
    local headers = {}
    headers["Content-Type"] = "application/json"
    return wrk.format(method, path, headers, body)
end

-- response = function(status, headers, body)
--     if status ~= 200 then
--         print("Faild, http status: " .. string(status))
--     else
--         -- local jsonparse = lunajson.decode(body)
--         -- local result_code = jsonparse["data"]["trackings"][1]["status"]["subcode"]
--         -- local tracking_number = jsonparse["data"]["trackings"][1]["tracking_number"]
--         -- print(result_code, tracking_number)
--     end
-- end

-- done = function(summary, latency, requests)
--     io.write("------------------------------\n")
--     for _, p in pairs({ 50, 90, 99, 99.999 }) do
--        n = latency:percentile(p)
--        io.write(string.format("%g%%,%d\n", p, n))
--     end
-- end

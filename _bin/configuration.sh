# GIN
export GIN_MODE="debug"
# SERVER APP
export SERVER_CONFIG="{
    \"pg-conn\": {
        \"dsn\": \"\",
        \"max-idle-conns\": 10,
        \"max-open-conns\": 10
    },
    \"redis\": {
        \"addr\": \"r-aaa.redis.rds.aliyuncs.com:6379\",
        \"password\": \"\",
        \"db\": 1
    },
    \"time-out\": 10,
    \"server-name\": \"server\"
}"

export PROJECT_CONFIG="{
    \"store-service\": {
        \"base-url\": \"https://aaa.cn/v1/int\",
        \"token\": \"\",
        \"timeout\": 5
    },
   \"name\": \"order-service\"
}"

echo "$SERVER_CONFIG" >s_settings.txt
export SERVER_CONFIG=$(echo "$SERVER_CONFIG" | base64)
echo "$SERVER_CONFIG" >>s_settings.txt

echo "$PROJECT_CONFIG" >p_settings.txt
export PROJECT_CONFIG=$(echo "$PROJECT_CONFIG" | base64)
echo "$PROJECT_CONFIG" >>p_settings.txt
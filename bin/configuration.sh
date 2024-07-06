# GIN
export GIN_MODE="debug"
# Server Conf
# pg and redis connection should work
export SERVER_CONFIG="{
    \"pg-conn\": {
        \"dsn\": \"host=pgm-aaa.pg.rds.aliyuncs.com port=5432 dbname=a user=a password=aaa\",
        \"max-idle-conns\": 10,
        \"max-open-conns\": 10
    },
    \"redis\": {
        \"addr\": \"124.222.50.11:6379\",
        \"username\": \"\",
        \"password\": \"\",
        \"db\": 1
    },
    \"time-out\": 10,
    \"server-name\": \"server\"
}"
# Project env
export PROJECT_CONFIG="{
    \"pilot-service\": {
        \"base-url\": \"https://pilot.samarkand-global.cn/v1/int\",
        \"token\": \"\",
        \"timeout\": 5
    },    
    \"encryption\": {
        \"key\": \"\",
        \"iv\": \"\"
    },
   \"name\": \"order-service\"
}"

echo "$SERVER_CONFIG" >s_settings.txt
export SERVER_CONFIG=$(echo "$SERVER_CONFIG" | base64)
echo "$SERVER_CONFIG" >>s_settings.txt

echo "$PROJECT_CONFIG" >p_settings.txt
export PROJECT_CONFIG=$(echo "$PROJECT_CONFIG" | base64)
echo "$PROJECT_CONFIG" >>p_settings.txt
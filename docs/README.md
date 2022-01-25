<div id="top"></div>

<!-- PROJECT SHIELDS -->
<p align="center">
<a href="https://github.com/hominsu/Ali-DDNS/graphs/contributors"><img src="https://img.shields.io/github/contributors/hominsu/XFileCrypt.svg?style=for-the-badge" alt="Contributors"></a>
<a href="https://github.com/hominsu/Ali-DDNS/network/members"><img src="https://img.shields.io/github/forks/hominsu/Ali-DDNS.svg?style=for-the-badge" alt="Forks"></a>
<a href="https://github.com/hominsu/Ali-DDNS/stargazers"><img src="https://img.shields.io/github/stars/hominsu/Ali-DDNS.svg?style=for-the-badge" alt="Stargazers"></a>
<a href="https://github.com/hominsu/Ali-DDNS/issues"><img src="https://img.shields.io/github/issues/hominsu/Ali-DDNS.svg?style=for-the-badge" alt="Issues"></a>
<a href="https://github.com/hominsu/Ali-DDNS/blob/master/LICENSE"><img src="https://img.shields.io/github/license/hominsu/Ali-DDNS.svg?style=for-the-badge" alt="License"></a>
<a href="https://github.com/hominsu/Ali-DDNS/actions/workflows/code_ql_analysis.yml"><img src="https://img.shields.io/github/workflow/status/hominsu/Ali-DDNS/CodeQL%20Analysis?style=for-the-badge" alt="CodeQL"></a>
</p>


<!-- PROJECT LOGO -->
<br/>
<div align="center">
<!--   <a href="https://github.com/hominsu/Ali-DDNS">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h3 align="center">Ali-DDNS</h3>

  <p align="center">
    DDNS service by using Ali openapi
    <br/>
    <a href="https://hominsu.github.io/Ali-DDNS/"><strong>Explore the docs » (you are here)</strong></a>
    <br/>
    <br/>
    <a href="https://github.com/hominsu/Ali-DDNS">View Demo</a>
    ·
    <a href="https://github.com/hominsu/Ali-DDNS/issues">Report Bug</a>
    ·
    <a href="https://github.com/hominsu/Ali-DDNS/issues">Request Feature</a>
  </p>
</div>

## Description

DDNS service by using Ali openapi

## Details

```
# Architecture

- Server: Obtain domain records from Ali and receive domain name change requests from clients
- Client: Gets the domain record from the server and requests to change the domain record

               +---+    Get Domain Record    +---+
               |   +------------------------>|   |
+------------+ | G |<------------------------+ G | +------------+    +---------+
| Client API | | R |   Return Domain Record  | R | | Server API +--->| Ali API |<------+
+------------+ | P |                         | P | +------+-----+    +---------+       |
               | C +------------------------>| C |        |          +-------+    +----+----+
               |   |   Update Domain Record  |   |        +--------->| Redis +--->| CronJob |
               +---+                         +---+                   +-------+    +---------+
                                                                           ^
               +---+                         +---+                +---+    |
+------------+ | H |  register/login/logout  | H |  grpc-gateway  | G | +--+---------+
|     Web    | | T |<----------------------->| T |<-------------->| R | | Interface  |
+------------+ | T |   add/del Domain Name   | T |                | P | |    API     |
               | P |                         | P |                | C | +------------+
               +---+                         +---+                +---+
```

## How to use

- Clone the project both server and client:

  ```bash
  git clone --depth=1 https://github.com/hominsu/Ali-DDNS.git
  ```

- In server-service, fill in your  Ali Access-Key in the `docker-compose.yml`, set your redis db password. The interface use jwt for authentication, may sure you set the jwt token (e.g. `www.hauhau.cn`). Usually the operator updates the public IP in the early morning, so use the Cron expression (e.g. `CRON_TZ=Asia/Shanghai 1/10 2-4 * * *`) to specify that it updates at a higher frequency during the early morning hours and at a slower rate (like once per hour) at other times. Of course you can define your own time.

  ```yaml
  redis-ddns:
    image: redis:alpine
    container_name: ali-ddns-redis-ddns
    # 设置 Redis 的密码，下面记得主服务中填写对应的密码
    command: redis-server --port 6380 --requirepass redis-ddns-password
    
  ali-ddns-server-service:
    # 阿里云仓库:  registry.cn-shenzhen.aliyuncs.com/hominsu/ali-ddns-server-service:latest
    # GHCR:      ghcr.io/hominsu/ali-ddns-server-service:latest
    # DockerHub: hominsu/ali-ddns-server-service:latest
    image: hominsu/ali-ddns-server-service:latest
    container_name: ali-ddns-server-service
    # build:
    #   context: .
    #   dockerfile: ./Dockerfile
    depends_on:
      - redis-session
      - redis-ddns
    restart: always
    environment:
    # 设置时区，不然 logs 的时间不对
    TZ: "Asia/Shanghai" # 时区
  
    # 设置阿里云的 AK，建议使用 RAM 用户，只分配 AliyunDNSFullAccess 权限
    ALIDDNSSERVER_ACCESSKEY_ID: "*"                                   # 阿里云 AK ID
    ALIDDNSSERVER_ACCESSKEY_SECRET: "*"                               # 阿里云 AK SECRET
  
    ALIDDNSSERVER_BASIC_ENDPOINT: "alidns.cn-shenzhen.aliyuncs.com"   # 阿里云服务地址
    ALIDDNSSERVER_BASIC_INTERFACE_PORT: "50001"                       # WEB 服务监听端口
    ALIDDNSSERVER_BASIC_DOMAIN_GRPC_NETWORK: "tcp"                    # RPC 协议
    ALIDDNSSERVER_BASIC_DOMAIN_GRPC_PORT: "50002"                     # RPC 服务端口
    ALIDDNSSERVER_BASIC_INTERFACE_GRPC_NETWORK: "tcp"                 # RPC 协议
    ALIDDNSSERVER_BASIC_INTERFACE_GRPC_PORT: "50003"                  # RPC 服务端口
  
    # 保存域名相关信息的 Redis，要改的只有密码，和上面设置的密码相同
    ALIDDNSSERVER_DOMAIN_RECORD_REDIS_ADDR: "redis-ddns"              # 保存域名信息的 Redis 地址
    ALIDDNSSERVER_DOMAIN_RECORD_REDIS_PORT: "6380"                    # 保存域名信息的 Redis 端口
    ALIDDNSSERVER_DOMAIN_RECORD_REDIS_PASSWORD: "redis-ddns-password" # 保存域名信息的 Redis 密码
    ALIDDNSSERVER_DOMAIN_RECORD_REDIS_DB: "1"                         # 保存域名信息的 Redis 数据库
    
    ALIDDNSCLIENT_OPTION_JWT_TOKEN: "www.hauhau.cn"                   # jwt token
    ALIDDNSCLIENT_OPTION_TTL: "3600"                                  # 每隔多少秒向服务端获取更新信息
    ALIDDNSCLIENT_OPTION_DELAY_CHECK_CRON: "CRON_TZ=Asia/Shanghai 1/10 2-4 * * *" # 每天的 2-4 点的 1m 开始，每 10m 执行一次
  
  ```
  
- Up the server-service

  ```bash
  cd Ali-DDNS/deploy/docker-compose/ddns-server
  docker-compose up -d
  ```

- Up the client-service

  ```bash
  cd Ali-DDNS/deploy/docker-compose/ddns-client
  docker-compose up -d
  ```

- Configuring the server service

  1. Register a user

     Send an http request via `cURL`:

     ```bash
     curl --location --request POST 'http://127.0.0.1:50001/v1/register' \ 
     --header 'Content-Type: application/json' \
     --data-raw '{
         "username": "admin",
         "password": "passwd"
     }'
     ```

     Or use `wget`:

     ```bash
     wget --no-check-certificate --quiet \
       --method POST \
       --timeout=0 \
       --header 'Content-Type: application/json' \
       --body-data '{
         "username": "admin",
         "password": "passwd"
     }' \
        'http://127.0.0.1:50001/v1/register'
     ```

     You will get the following output if the request is successful

     ```bash
     {"status":true}
     ```

  2. Login to get the token

     `cURL`:

     ```bash
     curl --location --request POST 'http://127.0.0.1:50001/v1/login' \
     --header 'Content-Type: application/json' \
     --data-raw '{
         "username": "admin",
         "password": "passwd"
     }'
     ```

     `wget`:

     ```bash
     wget --no-check-certificate --quiet \
       --method POST \
       --timeout=0 \
       --header 'Content-Type: application/json' \
       --body-data '{
         "username": "admin",
         "password": "passwd"
     }' \
        'http://127.0.0.1:50001/v1/login'
     ```

     You will get the following output with `token` and `username` if the request is successful 

     ```bash
     {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFkbWluIiwiaXNzIjoiMTI3LjAuMC4xIiwic3ViIjoidXNlciB0b2tlbiIsImV4cCI6MTY0MzEwNTY3MSwiaWF0IjoxNjQzMTAyMDcxfQ.EmYB_PApYocKSbdyT0ykUMPMJErMykv3AASBcYngJTQ", "username":"admin"}
     ```

  3. Add the domain name you need to monitor, note that the `{username}` in the url (`/v1/{username}/domain_name`) needs to be your own username (e.g. `admin`), and the token is obtained in the previous step

     `cURL`:

     ```bash
     curl --location --request POST 'http://127.0.0.1:50001/v1/admin/domain_name' \
     --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFkbWluIiwiaXNzIjoiMTI3LjAuMC4xIiwic3ViIjoidXNlciB0b2tlbiIsImV4cCI6MTY0MzEwNTY3MSwiaWF0IjoxNjQzMTAyMDcxfQ.EmYB_PApYocKSbdyT0ykUMPMJErMykv3AASBcYngJTQ' \
     --header 'Content-Type: application/json' \
     --data-raw '{
         "domain_name": "haomingsu.cn"
     }'
     ```

     `wget`:

     ```bash
     wget --no-check-certificate --quiet \
       --method POST \
       --timeout=0 \
       --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFkbWluIiwiaXNzIjoiMTI3LjAuMC4xIiwic3ViIjoidXNlciB0b2tlbiIsImV4cCI6MTY0MzEwNTY3MSwiaWF0IjoxNjQzMTAyMDcxfQ.EmYB_PApYocKSbdyT0ykUMPMJErMykv3AASBcYngJTQ' \
       --header 'Content-Type: application/json' \
       --body-data '{
         "domain_name": "haomingsu.cn"
     }' \
        'http://127.0.0.1:50001/v1/admin/domain_name'
     ```

     You will get the following output if the request is successful

     ```bash
     {"status":true, "domainName":"haomingsu.cn"}
     ```

  4. Check the domain name is added, note that the `{username}` in the url (`/v1/{username}/domain_name`) needs to be your own username (e.g. `admin`), and the token is obtained in the first two steps

     `cURL`:

     ```bash
     curl --location --request GET 'http://127.0.0.1:50001/v1/admin/domain_name' \
     --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFkbWluIiwiaXNzIjoiMTI3LjAuMC4xIiwic3ViIjoidXNlciB0b2tlbiIsImV4cCI6MTY0MzEwNTY3MSwiaWF0IjoxNjQzMTAyMDcxfQ.EmYB_PApYocKSbdyT0ykUMPMJErMykv3AASBcYngJTQ'
     ```

     `wget`:

     ```bash
     wget --no-check-certificate --quiet \
       --method GET \
       --timeout=0 \
       --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFkbWluIiwiaXNzIjoiMTI3LjAuMC4xIiwic3ViIjoidXNlciB0b2tlbiIsImV4cCI6MTY0MzEwNTY3MSwiaWF0IjoxNjQzMTAyMDcxfQ.EmYB_PApYocKSbdyT0ykUMPMJErMykv3AASBcYngJTQ' \
        'http://127.0.0.1:50001/v1/admin/domain_name'
     ```

     You will get the following output if the request is successful

     ```bash
     {"domainNames":["haomingsu.cn"]}
     ```

## Container Repository

- Docker Hub: 

  ```bash
  docker pull hominsu/ali-ddns-client-service:latest
  docker pull hominsu/ali-ddns-server-service:latest
  ```

- GitHub Container Repository: 

  ```bash
  docker pull ghcr.io/hominsu/ali-ddns-client-service:latest
  docker pull ghcr.io/hominsu/ali-ddns-server-service:latest
  ```

- Ali Container Repository: `registry.cn-shenzhen.aliyuncs.com/hominsu/ali-ddns-client:latest`

  ```bash
  docker pull registry.cn-shenzhen.aliyuncs.com/hominsu/ali-ddns-client-service:latest
  docker pull registry.cn-shenzhen.aliyuncs.com/hominsu/ali-ddns-server-service:latest
  ```

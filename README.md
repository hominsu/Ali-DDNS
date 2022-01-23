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
<br />
<div align="center">
<!--   <a href="https://github.com/hominsu/Ali-DDNS">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h3 align="center">Ali-DDNS</h3>

  <p align="center">
    DDNS service by using Ali openapi
    <br />
    <a href="https://github.com/hominsu/Ali-DDNS"><strong>Explore the docs »</strong></a>
    <br />
    <br />
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

               +---+    Get Domain Record     +---+
               |   +------------------------->|   |
+------------+ | G |<-------------------------+ G | +------------+     +---------+
| Client API | | R |   Return Domain Record   | R | | Server API +---->| Ali API |<------+
+------------+ | P |                          | P | +------+-----+     +---------+       |
               | C +------------------------->| C |        |                             |
               |   |   Update Domain Record   |   |        |           +-------+    +----+----+
               +---+                          +---+        +---------->| Redis +--->| CronJob |
               +---+                          +---+                    +-------+    +---------+
               | H |                          | H | +------------+         ^
+------------+ | T |  register/login/logout   | T | | Interface  |         |
|     Web    | | T |------------------------->| T | |    API     +---------+
+------------+ | P |   add/del Domain Name    | P | +------------+ 
               | S |                          | S |
               +---+                          +---+
```

## How to use

- Clone the project both server and client:

  ```bash
  ❯ git clone --depth=1 https://github.com/hominsu/Ali-DDNS.git
  ```

- In server-service, fill in your  Ali Access-Key in the `docker-compose.yml`, set your redis db pass, also you can set 

  ```yaml
  redis-ddns:
    image: redis:alpine
    container_name: ali-ddns-redis-ddns
    # 设置 Redis 的密码，下面记得主服务中填写对应的密码
    command: redis-server --port 6380 --requirepass redis-ddns-password
  
  redis-session:
    image: redis:alpine
    container_name: ali-ddns-redis-session
    # 设置 Redis 的密码，下面记得主服务中填写对应的密码
    command: redis-server --port 6381 --requirepass redis-session-password
    
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
  
    GIN_MODE: "release"
  
    # 设置阿里云的 AK，建议使用 RAM 用户，只分配 AliyunDNSFullAccess 权限
    ALIDDNSSERVER_ACCESSKEY_ID: "*"            # 阿里云 AK ID
    ALIDDNSSERVER_ACCESSKEY_SECRET: "*"  # 阿里云 AK SECRET
  
    ALIDDNSSERVER_BASIC_ENDPOINT: "alidns.cn-shenzhen.aliyuncs.com"   # 阿里云服务地址
    ALIDDNSSERVER_BASIC_WEB_PORT: "50001"                             # WEB 服务监听端口
    ALIDDNSSERVER_BASIC_RPC_NETWORK: "tcp"                            # RPC 协议
    ALIDDNSSERVER_BASIC_RPC_PORT: "50002"                             # RPC 服务端口
  
    # 保存域名相关信息的 Redis，要改的只有密码，和上面设置的密码相同
    ALIDDNSSERVER_DOMAIN_RECORD_REDIS_ADDR: "redis-ddns"              # 保存域名信息的 Redis 地址
    ALIDDNSSERVER_DOMAIN_RECORD_REDIS_PORT: "6380"                    # 保存域名信息的 Redis 端口
    ALIDDNSSERVER_DOMAIN_RECORD_REDIS_PASSWORD: "redis-ddns-password" # 保存域名信息的 Redis 密码
    ALIDDNSSERVER_DOMAIN_RECORD_REDIS_DB: "1"                         # 保存域名信息的 Redis 数据库
  
    # 保存 session 的 Redis，要改的有 Redis 的密码，以及 session 的密码
    ALIDDNSSERVER_SESSION_SIZE: "10"                                  # session
    ALIDDNSSERVER_SESSION_REDIS_NETWORK: "tcp"                        # session 协议
    ALIDDNSSERVER_SESSION_REDIS_ADDRESS: "redis-session"              # 保存 session 信息的 Redis 地址
    ALIDDNSSERVER_SESSION_REDIS_PORT: "6381"                          # 保存 session 信息的 Redis 端口
    ALIDDNSSERVER_SESSION_REDIS_PASSWORD: "redis-session-password"    # 保存 session 信息的 Redis 密码
    ALIDDNSSERVER_SESSION_SECRET: "secret"                            # session 密码
    ALIDDNSCLIENT_OPTION_TTL: "3600"                                  # 每隔多少秒向服务端获取更新信息
    ALIDDNSCLIENT_OPTION_DELAY_CHECK_CRON: "CRON_TZ=Asia/Shanghai 1/10 2-4 * * *" # 每天的 2-4 点的 1m 开始，每 10m 执行一次
  
  ```

- Up the server-service

  ```bash
  ❯ cd Ali-DDNS/deploy/docker-compose/ddns-server
  ❯ docker-compose up -d
  ```

- Up the client-service

  ```bash
  ❯ cd Ali-DDNS/deploy/docker-compose/ddns-client
  ❯ docker-compose up -d
  ```


## Container Repository

- Docker Hub: 

  ```bash
  ❯ docker pull hominsu/ali-ddns-client-service:latest
  ❯ docker pull hominsu/ali-ddns-server-service:latest
  ```

- GitHub Container Repository: 

  ```bash
  ❯ docker pull ghcr.io/hominsu/ali-ddns-client-service:latest
  ❯ docker pull ghcr.io/hominsu/ali-ddns-server-service:latest
  ```

- Ali Container Repository: `registry.cn-shenzhen.aliyuncs.com/hominsu/ali-ddns-client:latest`

  ```bash
  ❯ docker pull registry.cn-shenzhen.aliyuncs.com/hominsu/ali-ddns-client-service:latest
  ❯ docker pull registry.cn-shenzhen.aliyuncs.com/hominsu/ali-ddns-server-service:latest
  ```

  

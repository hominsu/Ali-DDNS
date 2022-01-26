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

- Using openssl to generate SANs certs

    1. Create CA's certs:

       ```bash
       ❯ openssl genrsa -out ca.key 4096
       Generating RSA private key, 4096 bit long modulus (2 primes)
       ..................................................................++++
       .........................++++
       e is 65537 (0x010001)
       
       ❯ openssl req -new -x509 -days 3650 -key ca.key -out ca.crt
       You are about to be asked to enter information that will be incorporated
       into your certificate request.
       What you are about to enter is what is called a Distinguished Name or a DN.
       There are quite a few fields but you can leave some blank
       For some fields there will be a default value,
       If you enter '.', the field will be left blank.
       -----
       Country Name (2 letter code) [AU]:CN
       State or Province Name (full name) [Some-State]:Guangdong
       Locality Name (eg, city) []:Foshan
       Organization Name (eg, company) [Internet Widgits Pty Ltd]:hominsu
       Organizational Unit Name (eg, section) []:hominsu
       Common Name (e.g. server FQDN or YOUR name) []:localhost
       Email Address []:1774069959@qq.com
       ```

    2. Prepare the openssl configuration file

       copy the default openssl to current dir

        - `linux`:

          ```bash
          cp /etc/pki/tls/openssl.cnf .

        - `macos`:

          ```bash
          cp /System/Library/OpenSSL/openssl.cnf .
          ```

       Modify the following options in `openssl.cnf`

        - Find `[ CA_default ]` and uncomment `copy_extensions = copy`:

          ```bash
          [ CA_default ]
          # Extension copying option: use with caution.
          copy_extensions = copy
          ```

        - Find `[ req ]` and uncomment `req_extensions = v3_req`:

          ```bash
          [ req ]
          req_extensions = v3_req # The extensions to add to a certificate request
          ```

        - Find `[ v3_req ]` and add `subjectAltName = @alt_names`, then add the new tag `[ alt_names ]` and the field:

          ```bash
          [ v3_req ]       
                           
          # Extensions to add to a certificate request
                           
          basicConstraints = CA:FALSE
          keyUsage = nonRepudiation, digitalSignature, keyEncipherment
          subjectAltName = @alt_names
                           
          [ alt_names ]    
          DNS.1 = localhost
          DNS.2 = *.example.com
          ```

    3. Generate server-side certs

       ```bash
       ❯ openssl genpkey -algorithm RSA -out server.key
       .....................................................................................................+++++
       ..........+++++
       
       ❯ openssl req -new -nodes -key server.key -out server.csr -days 3650 -config ./openssl.cnf -extensions v3_req
       Ignoring -days; not generating a certificate
       You are about to be asked to enter information that will be incorporated
       into your certificate request.
       What you are about to enter is what is called a Distinguished Name or a DN.
       There are quite a few fields but you can leave some blank
       For some fields there will be a default value,
       If you enter '.', the field will be left blank.
       -----
       Country Name (2 letter code) [AU]:CN
       State or Province Name (full name) [Some-State]:Guangdong
       Locality Name (eg, city) []:Foshan
       Organization Name (eg, company) [Internet Widgits Pty Ltd]:hominsu
       Organizational Unit Name (eg, section) []:hominsu
       Common Name (e.g. server FQDN or YOUR name) []:localhost
       Email Address []:1774069959@qq.com
       
       Please enter the following 'extra' attributes
       to be sent with your certificate request
       A challenge password []:your_password
       An optional company name []:hominsu
       
       ❯ openssl x509 -req -days 3650 -in server.csr -out server.pem -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req
       Signature ok
       subject=C = CN, ST = Guangdong, L = Foshan, O = hominsu, OU = hominsu, CN = localhost, emailAddress = 1774069959@qq.com
       Getting CA Private Key
       ```

    4. Generate client-side certs

       ```bash
       ❯ openssl genpkey -algorithm RSA -out client.key
       ..................................................................................+++++
       ..............................................+++++
       
       ❯ openssl req -new -nodes -key client.key -out client.csr -days 3650 -config ./openssl.cnf -extensions v3_req
       Ignoring -days; not generating a certificate
       You are about to be asked to enter information that will be incorporated
       into your certificate request.
       What you are about to enter is what is called a Distinguished Name or a DN.
       There are quite a few fields but you can leave some blank
       For some fields there will be a default value,
       If you enter '.', the field will be left blank.
       -----
       Country Name (2 letter code) [AU]:CN
       State or Province Name (full name) [Some-State]:Guangdong
       Locality Name (eg, city) []:Foshan
       Organization Name (eg, company) [Internet Widgits Pty Ltd]:hominsu
       Organizational Unit Name (eg, section) []:hominsu
       Common Name (e.g. server FQDN or YOUR name) []:localhost
       Email Address []:1774069959@qq.com
       
       Please enter the following 'extra' attributes
       to be sent with your certificate request
       A challenge password []:your_password
       An optional company name []:hominsu
       
       ❯ openssl x509 -req -days 3650 -in client.csr -out client.pem -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req
       Signature ok
       subject=C = CN, ST = Guangdong, L = Foshan, O = hominsu, OU = hominsu, CN = localhost, emailAddress = 1774069959@qq.com
       Getting CA Private Key
       ```

  The full file is shown below:

  ```bash
  ❯ tree
  .
  ├── ca.crt
  ├── ca.key
  ├── ca.srl
  ├── client.csr
  ├── client.key
  ├── client.pem
  ├── openssl.cnf
  ├── server.csr
  ├── server.key
  └── server.pem
  
  0 directories, 10 files
  ```

  In the server-side, the `ca.crt`, `server.pem`, `server.key` is using to set credentials for server connections, and the `grpc-gateway` will use the `ca.crt`, `client.pem`, `client.key` to set credentials for connections between `grpc-gateway` and `grpc server`

  In the client-side, for safety reasons, you should generate the certs for each client with the CA's certs (`ca.crt` and `ca.key`) if others use your grpc-services, or you can just use the same certs as gateway (not recommended)

  Then create a cert directory and copy the certs into it

  On the server side, the file structure is shown below:

  ```bash
  ❯ tree cert
  cert
  ├── ca.crt
  ├── client.key
  ├── client.pem
  ├── server.key
  └── server.pem
  
  0 directories, 5 files
  ```

  On the server side, the file structure is shown below:

  ```bash
  ❯ tree cert
  cert
  ├── ca.crt
  ├── client.key
  └── client.pem
  
  0 directories, 3 files
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
    
    ALIDDNSCLIENT_OPTION_JWT_TOKEN: "jwt_token"                       # jwt token
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

## Other

On the server-side, the logs are stored in the `ads/logs/ads.log`

```bash
[root@iZwz9diii276grug5qq3byZ ddns-server-service]# cat ads/logs/ads.log 
2022-01-26T13:59:54.423+0800    info    runtime/proc.go:255     service.id: e4fd4c9652af, service.name: ali-ddns-server-service, service.version: 1.2.5, git sha1: a3936d5d8b6044bbbed686d6b2222c2c5813fa39, build stamp: 1643173896
2022-01-26T13:59:54.428+0800    info    grpclog/grpclog.go:37   [core]original dial target is: "localhost:50003"        {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.429+0800    info    grpclog/grpclog.go:37   [core]parsed dial target is: {Scheme:localhost Authority: Endpoint:50003 URL:{Scheme:localhost Opaque:50003 User: Host: Path: RawPath: ForceQuery:false RawQuery: Fragment: RawFragment:}}    {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.429+0800    info    grpclog/grpclog.go:37   [core]fallback to scheme "passthrough"  {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.429+0800    info    grpclog/grpclog.go:37   [core]parsed dial target is: {Scheme:passthrough Authority: Endpoint:localhost:50003 URL:{Scheme:passthrough Opaque: User: Host: Path:/localhost:50003 RawPath: ForceQuery:false RawQuery: Fragment: RawFragment:}}   {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.429+0800    info    grpclog/grpclog.go:37   [core]Channel authority set to "localhost"      {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.429+0800    info    grpclog/grpclog.go:37   [core]ccResolverWrapper: sending update to cc: {[{localhost:50003  <nil> <nil> 0 <nil>}] <nil> <nil>} {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.429+0800    info    grpclog/grpclog.go:37   [core]ClientConn switching balancer to "pick_first"     {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.429+0800    info    grpclog/grpclog.go:37   [core]Channel switches to new LB policy "pick_first"    {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.429+0800    info    grpclog/grpclog.go:37   [core]Subchannel Connectivity change to CONNECTING      {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.429+0800    info    grpclog/grpclog.go:37   [core]Subchannel picks a new address "localhost:50003" to connect       {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.429+0800    info    grpclog/grpclog.go:37   [core]pickfirstBalancer: UpdateSubConnState: 0xc000171e40, {CONNECTING <nil>}   {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.429+0800    info    grpclog/grpclog.go:37   [core]Channel Connectivity change to CONNECTING {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.444+0800    info    grpclog/grpclog.go:37   [core]Subchannel Connectivity change to READY   {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.444+0800    info    grpclog/grpclog.go:37   [core]pickfirstBalancer: UpdateSubConnState: 0xc000171e40, {READY <nil>}        {"system": "grpc", "grpc_log": true}
2022-01-26T13:59:54.444+0800    info    grpclog/grpclog.go:37   [core]Channel Connectivity change to READY      {"system": "grpc", "grpc_log": true}
2022-01-26T14:13:57.084+0800    info    zap/server_interceptors.go:39   finished unary call with code OK        {"grpc.start_time": "2022-01-26T14:13:57+08:00", "system": "grpc", "span.kind": "server", "grpc.service": "server.service.v1.DomainService", "grpc.method": "GetDomainRecord", "grpc.code": "OK", "grpc.time_ms": 1.3580000400543213}
2022-01-26T14:23:57.094+0800    info    zap/server_interceptors.go:39   finished unary call with code OK        {"grpc.start_time": "2022-01-26T14:23:57+08:00", "system": "grpc", "span.kind": "server", "grpc.service": "server.service.v1.DomainService", "grpc.method": "GetDomainRecord", "grpc.code": "OK", "grpc.time_ms": 1.2719999551773071}
2022-01-26T14:33:44.731+0800    warn    grpclog/grpclog.go:46   [core]grpc: Server.Serve failed to create ServerTransport: connection error: desc = "ServerHandshake(\"47.103.37.203:36666\") failed: tls: client didn't provide a certificate"       {"system": "grpc", "grpc_log": true}
2022-01-26T14:33:44.797+0800    warn    grpclog/grpclog.go:46   [core]grpc: Server.Serve failed to create ServerTransport: connection error: desc = "ServerHandshake(\"47.103.37.203:36682\") failed: tls: first record does not look like a TLS handshake"   {"system": "grpc", "grpc_log": true}
2022-01-26T14:33:44.897+0800    warn    grpclog/grpclog.go:46   [core]grpc: Server.Serve failed to create ServerTransport: connection error: desc = "ServerHandshake(\"47.103.37.203:36696\") failed: tls: client didn't provide a certificate"       {"system": "grpc", "grpc_log": true}
2022-01-26T14:33:44.954+0800    warn    grpclog/grpclog.go:46   [core]grpc: Server.Serve failed to create ServerTransport: connection error: desc = "ServerHandshake(\"47.103.37.203:36710\") failed: tls: first record does not look like a TLS handshake"   {"system": "grpc", "grpc_log": true}
2022-01-26T14:33:57.008+0800    info    zap/server_interceptors.go:39   finished unary call with code OK        {"grpc.start_time": "2022-01-26T14:33:57+08:00", "system": "grpc", "span.kind": "server", "grpc.service": "server.service.v1.DomainService", "grpc.method": "GetDomainRecord", "grpc.code": "OK", "grpc.time_ms": 1.125}
```

## Netty Echo Backend - Docker Image

- Build - `$ docker build -t netty-echo-backend .`

1. For `h1c` requests
    - Run - `$ docker run -d -p 8688:8688 netty-echo-backend`
    - Test - `$ curl -v http://localhost:8688 -d "Hello Netty!"`
2. For `h1` requests
    - Run - `$ docker run -d -p 8688:8688 -e "SSL=true" netty-echo-backend`
    - Test - `$ curl -kv https://localhost:8688 -d "Hello Netty!"`
3. For `h2c` requests
    - Run - `$ docker run -d -p 8688:8688 -e "HTTP2=true" netty-echo-backend`
    - Test - `$ curl -v --http2 http://localhost:8688 -d "Hello Netty!"`
4. For `h2` requests
    - Run - `$ docker run -d -p 8688:8688 -e "HTTP2=true" -e "SSL=true" netty-echo-backend`
    - Test - `$ curl -kv --http2 https://localhost:8688 -d "Hello Netty!"`

> Docker Hub URL: https://hub.docker.com/r/ldclakmal/netty-echo-backend

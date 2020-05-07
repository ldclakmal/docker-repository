## Netty Echo Backend

![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/netty-echo-backend)

> Docker Hub URL: https://hub.docker.com/r/ldclakmal/netty-echo-backend

---

- Build - `$ docker build -t netty-echo-backend .`
- Run
    - h1c -`$ docker run -d -p 8688:8688 netty-echo-backend`
    - h1 - `$ docker run -d -p 8688:8688 -e "SSL=true" netty-echo-backend`
    - h2c - `$ docker run -d -p 8688:8688 -e "HTTP2=true" netty-echo-backend`
    - h2 - `$ docker run -d -p 8688:8688 -e "HTTP2=true" -e "SSL=true" netty-echo-backend`
- Test
    - h1c - `$ curl -v http://localhost:8688 -d "Hello Netty!"`
    - h1 - `$ curl -kv https://localhost:8688 -d "Hello Netty!"`
    - h2c - `$ curl -v --http2 http://localhost:8688 -d "Hello Netty!"`
    - h2 - `$ curl -kv --http2 https://localhost:8688 -d "Hello Netty!"`

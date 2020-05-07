## GO HTTP Passthrough Service

![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/go-http-passthrough)

Go Version: 1.14.2

Docker Hub URL: https://hub.docker.com/r/ldclakmal/go-http-passthrough

---

- Build - `$ docker build -t go-http-passthrough .`
- Run
    - HTTP/1.1 - `$ docker run -d --network="host" -p 9090:9090 -e "VERSION=1" -e "BACKENDHOST=localhost" -e "BACKENDPORT=9191" -e "BACKENDPATH=/hello/sayHello" go-http-passthrough`
    - HTTP/2 - `$ docker run -d --network="host" -p 9090:9090 -e "VERSION=2" -e "MAXSTREAM=1000" -e "BACKENDHOST=localhost" -e "BACKENDPORT=9191" -e "BACKENDPATH=/hello/sayHello" go-http-passthrough`
- Test
    - HTTP/1.1- `$ curl -kv https://localhost:9090/passthrough -d "Hello GoLang"`
    - HTTP/2- `$ curl --http2 -kv https://localhost:9090/passthrough -d "Hello GoLang"`

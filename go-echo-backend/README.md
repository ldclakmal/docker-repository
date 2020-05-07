## GO Echo Backend

![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/go-echo-backend)

Go Version: 1.14.2

Docker Hub URL: https://hub.docker.com/r/ldclakmal/go-echo-backend

---

- Build - `$ docker build -t go-echo-backend .`
- Run
    - HTTP/1.1 - `$ docker run -d -p 9191:9191 -e "VERSION=1" go-echo-backend`
    - HTTP/2 - `$ docker run -d -p 9191:9191 -e "VERSION=2" -e "MAXSTREAM=1000" go-echo-backend`
- Test
    - HTTP/1.1 - `$ curl -kv https://localhost:9191/hello/sayHello -d "Hello GoLang"`
    - HTTP/2- `$ curl --http2 -kv https://localhost:9191/hello/sayHello -d "Hello GoLang"`

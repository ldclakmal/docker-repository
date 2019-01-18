## GO Echo Backend - Docker Image

- Build - `$ docker build -t go-echo-backend .`
---
- Run HTTP/1.1 - `$ docker run -d -p 9191:9191 -e "VERSION=1" go-echo-backend`
- Test - `$ curl -kv https://localhost:9191/hello/sayHello -d "Hello GoLang"`
---
- Run HTTP/2 - `$ docker run -d -p 9191:9191 -e "VERSION=2" -e "MAXSTREAM=1000" go-echo-backend`
- Test HTTP/2- `$ curl --http2 -kv https://localhost:9191/hello/sayHello -d "Hello GoLang"`

> Docker Hub URL: https://hub.docker.com/r/ldclakmal/go-echo-backend

## Ballerina Echo Backend

![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/ballerina-echo-backend)

> Docker Hub URL: https://hub.docker.com/r/ldclakmal/ballerina-echo-backend

---

- Build - `$ docker build -t ballerina-echo-backend .`
- Run - `$ docker run -d -p 9191:9191 ballerina-echo-backend`
- Test
    - HTTP/1.1 - `$ curl -v http://localhost:9191/hello/sayHello -d "Hello Ballerina"`
    - HTTP/2.0 - `$ curl --http2 -v http://localhost:9191/hello/sayHello -d "Hello Ballerina"`


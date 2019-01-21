## Ballerina Echo Backend - Docker Image

- Build - `$ docker build -t ballerina-echo-backend .`
- Run - `$ docker run -d -p 9191:9191 ballerina-echo-backend`
- Test
    1. `$ curl -v http://localhost:9191/hello/sayHello -d "Hello Ballerina"`
    2. `$ curl --http2 -v http://localhost:9191/hello/sayHello -d "Hello Ballerina"`

> Docker Hub URL: https://hub.docker.com/r/ldclakmal/ballerina-echo-backend

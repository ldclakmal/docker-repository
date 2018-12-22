## Ballerina HTTP/2 Echo Backend - Docker Image

- Build - `$ docker build -t ballerina-http2-echo-backend .`
- Run - `$ docker run -d -p 9191:9191 ballerina-http2-echo-backend`
- Test - `$ curl --http2 -v http://localhost:9191/hello/sayHello -H "Content-Type: application/json" -d '{"message":"Hello Ballerina"}'`

> Docker Hub URL: https://hub.docker.com/r/ldclakmal/ballerina-http2-echo-backend

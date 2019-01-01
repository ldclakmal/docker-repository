## Ballerina HTTP Passthrough Service - Docker Image

- Build - `$ docker build -t ballerina-http-passthrough .`
- Run - `$ docker run -d --network="host" -p 9090:9090 ballerina-http-passthrough`
- Test - `$ curl -v http://localhost:9090/passthrough -d "Hello Ballerina"`

> Docker Hub URL: https://hub.docker.com/r/ldclakmal/ballerina-http-passthrough

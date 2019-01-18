## WSO2 Transport - HTTP Passthrough Service - Docker Image

- Build - `$ docker build -t wso2-transport-http-passthrough .`
- Run - `$ docker run -d --network="host" -p 9090:9090 -e "HOST=localhost" -e "PORT=9191" wso2-transport-http-passthrough`
- Test - `$ curl -v http://localhost:9090 -d "Hello WSO2"`

> Docker Hub URL: https://hub.docker.com/r/ldclakmal/wso2-transport-http-passthrough

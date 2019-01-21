## WSO2 Transport - HTTP Echo Backend - Docker Image

- Build - `$ docker build -t wso2-transport-echo-backend .`
- Run - `$ docker run -d -p 9191:9191 wso2-transport-echo-backend`
- Test - `$ curl -v http://localhost:9191 -d "Hello WSO2"`

> Docker Hub URL: https://hub.docker.com/r/ldclakmal/wso2-transport-echo-backend

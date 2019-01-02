## GO Echo Backend - Docker Image

- Build - `$ docker build -t go-echo-backend .`
- Run - `$ docker run -d -p 9191:9191 go-echo-backend`
- Test - `$ curl -v http://localhost:9191/hello/sayHello -d "Hello GoLang"`

> Docker Hub URL: https://hub.docker.com/r/ldclakmal/go-echo-backend

## Envoy HTTP Simple Proxy - Docker Image

- Build - `$ docker build -t envoy .`
- Run - `$ docker run -d --network="host" --name envoy -p 9901:9901 -p 9090:9090 envoy`
- Test - `$ curl -v http://localhost:9090`

(Since this is a proxy, there should be a backend HTTP service already started on port `9191` with request path `/`. i.e. `$ curl -v http://localhost:9191` should be up and running)

> Docker Hub URL:  https://hub.docker.com/r/ldclakmal/envoy

#### Resources:
1. https://www.envoyproxy.io/docs/envoy/latest/start/start
2. https://github.com/envoyproxy/envoy/tree/master/examples

# Docker Repository - [ldclakmal/**](https://hub.docker.com/u/ldclakmal)

![Docker Logo](https://www.docker.com/wp-content/uploads/2022/03/horizontal-logo-monochromatic-white.png)

This is a collection of Dockerfiles for images that can be used to different integration scenarios for microservices. All the images are published on DockerHub at https://hub.docker.com/u/ldclakmal.

### Netty - https://netty.io

- Netty Echo Backend - [ldclakmal/netty-echo-backend](https://hub.docker.com/r/ldclakmal/netty-echo-backend) ![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/netty-echo-backend)

    A Docker image of Netty HTTP service which echo back the request payload. Supports both HTTP/1.1 and HTTP/2 versions and SSL. Also, supports any type of request payload such as text, json, xml etc.

### WSO2 - https://wso2.com

- WSO2 IS STS - [ldclakmal/wso2is-sts](https://hub.docker.com/r/ldclakmal/wso2is-sts) ![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/wso2is-sts)

    A Docker image of WSO2 IS as a STS. Supports **JWT**, **OAuth2** token issuing and validation.

### Ballerina - https://ballerina.io

- Ballerina Echo Backend - [ldclakmal/ballerina-echo-backend](https://hub.docker.com/r/ldclakmal/ballerina-echo-backend) ![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/ballerina-echo-backend)

    A Docker image of Ballerina service which echo back the request payload. Supports both **HTTP/1.1** and **HTTP/2**. Also, supports any type of request payload such as text, json, xml etc.

- Ballerina HTTP Passthrough Service - [ldclakmal/ballerina-http-passthrough](https://hub.docker.com/r/ldclakmal/ballerina-http-passthrough) ![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/ballerina-http-passthrough)

    A Docker image of Ballerina passthrough service which forward the request to configured backend and forward the response to the client. Supports only **HTTP/1.1** and any type of request payload such as text, json, xml etc.

### Go - https://golang.org

- GO Echo Backend - [ldclakmal/go-echo-backend](https://hub.docker.com/r/ldclakmal/go-echo-backend) ![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/go-echo-backend)

    A Docker image of Go service which echo back the request payload. Supports both **HTTP/1.1** and **HTTP/2** with **SSL**. Also, supports any type of request payload such as text, json, xml etc.

- Go HTTP Passthrough Service - [ldclakmal/go-http-passthrough](https://hub.docker.com/r/ldclakmal/go-http-passthrough) ![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/go-http-passthrough)

    A Docker image of Go passthrough service which forward the request to configured backend and forward the response to the client. Supports both **HTTP/1.1** and **HTTP/2.0** with **SSL**. Also, supports any type of request payload such as text, json, xml etc.

### Envoy - https://www.envoyproxy.io

- Envoy HTTP Proxy - [ldclakmal/envoy](https://hub.docker.com/r/ldclakmal/envoy) ![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/envoy)

    A Docker image of simple Envoy proxy service which forward the request to configured backend and forward the response to the client. Supports only **HTTP/1.1**.

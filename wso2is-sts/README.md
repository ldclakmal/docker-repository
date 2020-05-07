## WSO2 IS STS

![Docker Pulls](https://img.shields.io/docker/pulls/ldclakmal/wso2is-sts)

WSO2 IS Version: 5.7.0

Java Version: 8

Docker Hub URL: https://hub.docker.com/r/ldclakmal/wso2is-sts

---

- Build - `$ docker build -t wso2is-sts .`
- Run - `$ docker run -d -p 9443:9443 wso2is-sts`
- Test
    - JWT - `curl -kv -u 3VTwFk7u1i366wzmvpJ_LZlfAV4a:TNOH0ZklJm8sqS9U3IMMAZVth78a 
                  -H "Content-Type: application/x-www-form-urlencoded;charset=UTF-8" 
                  -d "grant_type=client_credentials&scope=view-order" 
                  https://localhost:9443/oauth2/token`
    - OAuth2 - `curl -kv -u FlfJYKBD2c925h4lkycqNZlC2l4a:PJz0UhTJMrHOo68QQNpvnqAY_3Aa 
                  -H "Content-Type: application/x-www-form-urlencoded;charset=UTF-8" 
                  -d "grant_type=client_credentials&scope=view-order" 
                  https://localhost:9443/oauth2/token`

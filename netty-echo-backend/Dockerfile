FROM java:8
WORKDIR /
ADD netty-http-echo-service.jar netty-http-echo-service.jar
COPY ./cert/keystore.p12 /etc/netty/cert/keystore.p12
EXPOSE 8688
ENV HTTP2=false
ENV SSL=false
CMD java -jar netty-http-echo-service.jar --ssl $SSL --http2 $HTTP2 --key-store-file /etc/netty/cert/keystore.p12 --key-store-password ballerina

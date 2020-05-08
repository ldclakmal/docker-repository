FROM maven:3.6.3-jdk-8
LABEL maintainer="ldclakmal@gmail.com"
WORKDIR /
COPY ./cert/keystore.p12 /etc/netty/cert/keystore.p12
RUN git clone https://github.com/ldclakmal/netty-http-transport-sample.git
WORKDIR /netty-http-transport-sample
RUN mvn clean install
EXPOSE 8688
ENV HTTP2=false
ENV SSL=false
CMD java -jar target/netty-http-echo-service.jar --ssl $SSL --http2 $HTTP2 --key-store-file /etc/netty/cert/keystore.p12 --key-store-password ballerina

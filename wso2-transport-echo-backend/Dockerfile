FROM maven:3.6.3-jdk-8
WORKDIR /
RUN git clone --branch v1.0 https://github.com/ldclakmal/wso2-http-transport-sample.git
WORKDIR /wso2-http-transport-sample
RUN mvn clean install
EXPOSE 9191
CMD java -cp target/wso2-http-transport.jar lk.avix.http.listener.HttpEchoServer

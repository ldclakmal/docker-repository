FROM java:8
WORKDIR /
ADD wso2-transport-http-passthrough.jar wso2-transport-http-passthrough.jar
EXPOSE 9090
CMD java -jar -Dhost=$HOST -Dport=$PORT wso2-transport-http-passthrough.jar

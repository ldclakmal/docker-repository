FROM java:8
WORKDIR /
ADD wso2-transport-http-passthrough.jar wso2-transport-http-passthrough.jar
EXPOSE 9090
CMD java -jar wso2-transport-http-passthrough.jar -host $HOST -port $PORT

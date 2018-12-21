FROM openjdk:jre-alpine
LABEL maintainer="chanakal@wso2.com"

# Ballerina runtime distribution filename.
ARG BALLERINA_DIST=ballerina-0.990.0.zip

# Add Ballerina runtime.
COPY ${BALLERINA_DIST} /root/
COPY http2-echo-backend.bal /root/

# Create folders, unzip distribution, create users, & set permissions.
RUN mkdir -p /ballerina/files \
    && addgroup troupe \
    && adduser -S -s /bin/bash -g 'ballerina' -G troupe -D ballerina \
    && apk add --update --no-cache bash \
    && unzip /root/${BALLERINA_DIST} -d /ballerina/ > /dev/null 2>&1 \
    && mv /root/http2-echo-backend.bal /ballerina/ \
    && mv /ballerina/ballerina* /ballerina/runtime \
    && mkdir -p /ballerina/runtime/logs \
    && chown -R ballerina:troupe /ballerina \
    && rm -rf /root/${BALLERINA_DIST} > /dev/null 2>&1 \
    && rm -rf /var/cache/apk/*

ENV BALLERINA_HOME /ballerina/runtime
ENV PATH $BALLERINA_HOME/bin:$PATH

WORKDIR /home/ballerina

USER ballerina

CMD $BALLERINA_HOME/bin/ballerina run /ballerina/http2-echo-backend.bal

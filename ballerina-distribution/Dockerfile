FROM java:8
LABEL maintainer="ldclakmal@gmail.com"
RUN mkdir /ballerina
COPY setup.sh /ballerina
WORKDIR /ballerina
CMD sh setup.sh $GITHUB_USER $GITHUB_TOKEN && tail -f /dev/null

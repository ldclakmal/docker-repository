FROM ballerina/ballerina:swan-lake-alpha5
LABEL maintainer="ldclakmal@gmail.com"
ADD https://raw.githubusercontent.com/ldclakmal/ballerina-security/master/scenarios/sts/cert/private.key /home/ballerina/cert/private.key
ADD https://raw.githubusercontent.com/ldclakmal/ballerina-security/master/scenarios/sts/cert/public.crt /home/ballerina/cert/public.crt
ADD https://raw.githubusercontent.com/ldclakmal/ballerina-security/master/scenarios/sts/ballerina-sts.bal /home/ballerina/ballerina-sts.bal
USER root
RUN chmod 766 /home/ballerina/ballerina-sts.bal
RUN chmod 766 /home/ballerina/cert/private.key
RUN chmod 766 /home/ballerina/cert/public.crt
USER ballerina
EXPOSE 9445
CMD bal run ballerina-sts.bal

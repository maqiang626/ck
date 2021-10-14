FROM centos:centos7.9.2009

MAINTAINER maqiang

ENV MY_SERVICE_PORT=9001

LABEL multi.label1="value1" multi.label2="value2" other="value3"

ADD bin/amd64/httpserver /httpserver

# httpserver listen port
EXPOSE 9001

ENTRYPOINT /httpserver

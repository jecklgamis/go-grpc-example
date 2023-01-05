FROM ubuntu:22.04
MAINTAINER Jerrico Gamis <jecklgamis@gmail.com>

RUN apt-get update -y


RUN mkdir -p /app/bin

COPY bin/server-linux-amd64 /app/bin/server

RUN  chmod +x /app/bin/*

EXPOSE 4000

WORKDIR /app
COPY docker-entrypoint.sh /
CMD ["/docker-entrypoint.sh"]


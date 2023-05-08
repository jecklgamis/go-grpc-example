FROM ubuntu:22.04
MAINTAINER Jerrico Gamis <jecklgamis@gmail.com>

RUN apt update -y && apt install -y golang python3 python3-pip dumb-init && rm -rf /var/lib/apt/lists/*
COPY requirements.txt /
RUN pip install --upgrade pip &&  pip install -r /requirements.txt

COPY supervisor.ini /etc/supervisor.d/
RUN mkdir -p /var/log/supervisor

ENV APP_HOME /app

RUN groupadd -r app && useradd -r -gapp app
RUN mkdir -m 0755 -p ${APP_HOME}/bin

COPY bin/*linux* ${APP_HOME}/bin/
COPY docker-entrypoint.sh /

RUN chown -R app:app ${APP_HOME}
RUN chmod +x /docker-entrypoint.sh

EXPOSE 4000
EXPOSE 8080

WORKDIR ${APP_HOME}
CMD ["/docker-entrypoint.sh"]

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/docker-entrypoint.sh"]



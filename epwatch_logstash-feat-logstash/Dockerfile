FROM alpine:3.15

LABEL maintainer "https://github.com/redzenova"

COPY --from=golang:1.18-alpine /usr/local/go/ /usr/local/go/
 
ENV PATH="/usr/local/go/bin:${PATH}"
# Install repository jdk11 python3/pip3
# ENV PYTHONUNBUFFERED=1
# RUN apk add --no-cache openjdk11-jre python3 py3-pip go screen
RUN apk add --no-cache openjdk11-jre screen

# ENV Logstash setting
ENV LOGSTASH_VERSION 8.3.0
ENV LOGSTASH_DOWNLOAD_URL https://artifacts.elastic.co/downloads/logstash
ENV LOGSTASH_TARBALL "${LOGSTASH_DOWNLOAD_URL}/logstash-oss-${LOGSTASH_VERSION}-linux-x86_64.tar.gz"
ENV GO_CONFIG_PATH /usr/share/logstash/app/config
# ENV Filebeat setting
# ENV FILEBEAT_VERSION 8.3.0
# ENV FILEBEAT_DOWNLOAD_URL https://artifacts.elastic.co/downloads/beats/filebeat
# ENV FILEBEAT_TARBALL "${FILEBEAT_DOWNLOAD_URL}/filebeat-oss-${FILEBEAT_VERSION}-linux-x86_64.tar.gz"
#ENV FILEBEAT_PATH /usr/share/filebeat
#ENV PATH $FILEBEAT_PATH:$PATH

# Start and enable SSH
# COPY sshd_config /etc/ssh/

# RUN apk add openssh \
#      && echo "root:Docker!" | chpasswd \
#      && cd /etc/ssh/ \
#      && ssh-keygen -A

# Provide a non-root user to run the process.
RUN addgroup --gid 1000 logstash && \
  adduser -u 1000 -G logstash \
  -h /usr/share/logstash -H -D \
  logstash

RUN apk add --no-cache libzmq bash
# RUN apk add --no-cache -t .build-deps wget ca-certificates gnupg1 openssl \
RUN apk add --no-cache -t .build-deps wget ca-certificates gnupg1 \
  && set -ex \
  && cd /tmp \
  && wget --progress=bar:force -O logstash-${LOGSTASH_VERSION}-linux-x86_64.tar.gz "$LOGSTASH_TARBALL"; \
  tar -xzf logstash-${LOGSTASH_VERSION}-linux-x86_64.tar.gz \
  && mv logstash-$LOGSTASH_VERSION /usr/share/logstash \
  && rm -rf /usr/share/logstash/jdk \
  && chown --recursive logstash:logstash /usr/share/logstash/ \
  && chown -R logstash:root /usr/share/logstash \
  && chmod -R g=u /usr/share/logstash \
  && find /usr/share/logstash -type d -exec chmod g+s {} \; \
  && ln -s /usr/share/logstash /opt/logstash \
  #&& rm -rf /tmp/* \
  && apk del --purge .build-deps

RUN apk add --no-cache libc6-compat

# RUN  cd /tmp \
#   && mkdir /usr/share/filebeat \ 
#   && wget -O filebeat-${FILEBEAT_VERSION}-linux-x86_64.tar.gz "$FILEBEAT_TARBALL"; \
#   tar -xzf filebeat-${FILEBEAT_VERSION}-linux-x86_64.tar.gz \
#   && mv filebeat-${FILEBEAT_VERSION}-linux-x86_64  filebeat \
#   && mv filebeat /usr/share \
#   && mkdir /usr/share/filebeat/logs /usr/share/filebeat/data \
#   && chown --recursive logstash:logstash /usr/share/filebeat/ \
#   && chown -R logstash:root /usr/share/filebeat \
#   && chmod -R g=u /usr/share/filebeat \
#   && find /usr/share/filebeat -type d -exec chmod g+s {} \; \
#   && ln -s /usr/share/filebeat /opt/filebeat \
#   && rm /tmp/*

# ENV PATH /usr/share/filebeat/:/sbin:$PATH
ENV PATH /usr/share/logstash/bin:/sbin:$PATH
ENV LS_SETTINGS_DIR /usr/share/logstash/config
ENV LANG='en_US.UTF-8' LC_ALL='en_US.UTF-8'

RUN set -ex; \
  if [ -f "$LS_SETTINGS_DIR/log4j2.properties" ]; then \
  cp "$LS_SETTINGS_DIR/log4j2.properties" "$LS_SETTINGS_DIR/log4j2.properties.dist"; \
  truncate -s 0 "$LS_SETTINGS_DIR/log4j2.properties"; \
  fi


WORKDIR /usr/share/logstash

COPY config/logstash /usr/share/logstash/config/
COPY config/pipeline/ /usr/share/logstash/pipeline/
# COPY data/ /usr/share/logstash/data/
COPY patterns/ /usr/share/logstash/patterns/
COPY logstash-entrypoint.sh /
RUN chmod +x /logstash-entrypoint.sh
RUN chown --recursive logstash:root config/ pipeline/ patterns/ 

RUN mkdir app
RUN cd /usr/share/logstash/app
# COPY requirements.txt /usr/share/logstash/app
# RUN pip install -r /usr/share/logstash/app/requirements.txt
# COPY app.py /usr/share/logstash/app
COPY config/goapp/ $GO_CONFIG_PATH
COPY go.mod /usr/share/logstash/app
COPY go.sum /usr/share/logstash/app
COPY internal /usr/share/logstash/app/internal/
COPY hooks /usr/share/logstash/app/
COPY cmd /usr/share/logstash/app/cmd/
RUN chown --recursive logstash:root app/

# WORKDIR /usr/share/filebeat
# COPY config/filebeat/ /usr/share/filebeat/
# RUN chmod go-w /usr/share/filebeat/filebeat.yml

WORKDIR /usr/share/logstash/app
RUN go version
RUN go mod tidy
RUN go build cmd/app/main.go
RUN chown logstash:root main
RUN rm -r cmd/ && \ 
    rm -r internal/ && \
    rm -r go.mod && \
    rm -r go.sum

WORKDIR /usr/share/logstash
USER 1000

EXPOSE 6514/udp 6514/tcp 9600 

USER root
ENV JAVA_HOME /usr/lib/jvm/java-11-openjdk 

ENTRYPOINT ["/logstash-entrypoint.sh"]
CMD ["-e", ""]
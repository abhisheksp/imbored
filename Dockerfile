FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/bored /opt/bin/bored

CMD ["/opt/bin/bored"]
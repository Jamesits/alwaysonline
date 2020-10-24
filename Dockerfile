# build stage
FROM golang:1.15.3-buster as builder

WORKDIR /root/alwaysonline
COPY . /root/alwaysonline/
RUN bash ./build.sh

# production stage
FROM debian:buster-slim
LABEL maintainer="docker@public.swineson.me"

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/group /etc/

COPY --from=builder /root/alwaysonline/build/alwaysonline /usr/local/bin/

# nope
# See: https://github.com/moby/moby/issues/8460
# USER nobody:nogroup

EXPOSE 53/tcp 53/udp 80/tcp
ENTRYPOINT [ "/usr/local/bin/alwaysonline" ]
CMD []
HEALTHCHECK NONE

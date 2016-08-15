FROM scratch

MAINTAINER Mitch Anderson <mitch@nuvi.com>

ADD gimlet /

ENV PORT=8080
EXPORT 8080
CMD ["/gimlet"]

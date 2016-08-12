FROM scratch

MAINTAINER Mitch Anderson <mitch@nuvi.com>

ADD gimlet /

ENV PORT=8080

CMD ["/gimlet"]

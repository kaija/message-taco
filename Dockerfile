FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/kaija/message-taco

ENV USER kaija
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET b3KnGMmYxp8Ly4Wy

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://kaija@localhost:5432/message-taco?sslmode=disable

WORKDIR /go/src/github.com/kaija/message-taco

RUN godep go build

EXPOSE 8888
CMD ./message-taco
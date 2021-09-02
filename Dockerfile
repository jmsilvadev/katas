FROM golang:1.13-alpine
ENV CGO_ENABLED 0
RUN mkdir /cycloid
COPY pkg/. /cycloid
WORKDIR /cycloid
RUN apk add git && \
    go get -u golang.org/x/lint/golint
ENTRYPOINT ./run-tests.sh
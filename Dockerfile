FROM golang:1.17-alpine AS build

RUN mkdir /build
WORKDIR /build

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN export GO111MODULE=on
RUN go get github.com/fidellr/fishery_api
RUN cd /build && git clone https://github.com/fidellr/fishery_api.git

RUN cd /build/fishery_api && go build

EXPOSE 8080

ENTRYPOINT [ "/build/fishery_api/examples/main" ]
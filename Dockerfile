FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/fidellr/fishery_api.git
RUN cd /build && git clone https://github.com/fidellr/fishery_api.git

RUN cd /build/fishery_api/examples && go build

EXPOSE 8080

ENTRYPOINT [ "/build/fishery_api/examples/main" ]
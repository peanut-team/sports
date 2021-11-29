FROM golang:1.16-stretch AS build

ARG ARCH

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=${ARCH}

WORKDIR /go/cache

WORKDIR /app

ARG GOPROXY
ARG GOSUMDB=sum.golang.org

ADD go.mod .
RUN go mod download

ADD . .

RUN go build -ldflags "-s -w" -o ./dist/api-service ./main.go

FROM alpine:3.9

WORKDIR /app

ARG PROD_NAME

ENV PROD_NAME ${PROD_NAME}

ADD config config

COPY --from=build /app/dist/api-service /app/dist/api-service

ARG VERSION
ENV VERSION ${VERSION}

ENTRYPOINT ["/app/dist/api-service"]

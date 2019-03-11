FROM golang:1.11.2-alpine3.8 AS build-env
WORKDIR /src

COPY . /src

# go module needs alpine-sdk and git
RUN apk update \
    && apk add alpine-sdk \
    && apk add git

# test and build
RUN go test ./... \
    && GOOS=linux GOARCH=amd64 go build -o /simple

FROM alpine:3.8
RUN apk update \ 
    && apk add ca-certificates \
    && apk add curl \
    && rm -rf /var/cache/apk/*

COPY ./index.gohtml /index.gohtml
COPY --from=build-env /simple /

EXPOSE 8080

ENTRYPOINT /simple
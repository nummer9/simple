FROM golang:1.11.0-alpine3.8 AS build-env
WORKDIR /go/src/github.com/BloodyRainer/simple

COPY . /go/src/github.com/BloodyRainer/simple

# go module needs alpine-sdk and git
RUN apk update \
    && apk add alpine-sdk \
    && apk add git

# enable go module
ENV GO111MODULE=on
RUN go mod vendor \
    && go build -o simple \
    && mv ./simple /simple

FROM alpine
RUN apk update \ 
    && apk add ca-certificates \
    && apk add curl \
    && rm -rf /var/cache/apk/*
COPY --from=build-env /simple /

EXPOSE 8080
ENTRYPOINT /simple
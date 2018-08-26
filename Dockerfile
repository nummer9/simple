FROM golang:alpine AS build-env
WORKDIR /go/src/github.com/BloodyRainer/simple
COPY . /go/src/github.com/BloodyRainer/simple
RUN go build -o simple \
     && mv ./simple /simple

FROM alpine
RUN apk update \ 
    && apk add ca-certificates \
    && apk add curl \
    && rm -rf /var/cache/apk/*
COPY --from=build-env /simple /

EXPOSE 8080
ENTRYPOINT /simple
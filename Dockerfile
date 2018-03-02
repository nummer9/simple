FROM ubuntu:17.10

RUN apt-get update && apt-get install -y \
curl

COPY ./simple /usr/local/bin/simple

CMD ["simple"]
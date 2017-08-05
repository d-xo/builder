FROM golang:1.8

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    ruby-full \
&& apt-get clean

RUN gem install aruba cucumber

RUN go get github.com/golang/lint/golint

VOLUME /go/src/github.com/xwvvvvwx/builder
WORKDIR /go/src/github.com/xwvvvvwx/builder

# https://stackoverflow.com/a/5395932
RUN chmod go-w /go/
RUN chmod go-w /go/bin

CMD /bin/bash

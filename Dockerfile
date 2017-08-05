FROM docker:stable

# -------------------------------------------------------------------

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin:/usr/local/go/bin

# -------------------------------------------------------------------

RUN apk update && apk add \
    build-base \
    git \
    go \
    libffi-dev \
    ruby \
    ruby-dev \
    ruby-irb \
    ruby-rdoc

RUN gem install aruba cucumber
RUN go get github.com/golang/lint/golint

# -------------------------------------------------------------------

VOLUME /go/src/github.com/xwvvvvwx/builder
WORKDIR /go/src/github.com/xwvvvvwx/builder

# -------------------------------------------------------------------

CMD /bin/ash

# -------------------------------------------------------------------

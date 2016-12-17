FROM golang:1.6.3

RUN go get -u github.com/kardianos/govendor
RUN go get -u github.com/spf13/cobra/cobra
RUN go get -u github.com/jteeuwen/go-bindata/go-bindata
RUN go get -u github.com/pressly/goose/cmd/goose

ENV PROJECTPATH=/go/src/github.com/replicatedcom/gin-example

EXPOSE 6060 8080

WORKDIR $PROJECTPATH

CMD ["bin/gin-example api"]

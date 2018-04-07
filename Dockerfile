FROM golang:stretch
RUN go get github.com/KaesaHuanyu/minitimespace
WORKDIR /go/src/github.com/KaesaHuanyu/minitimespace
RUN go install
ENTRYPOINT [ "go/bin/minitimespace" ]
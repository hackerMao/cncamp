#FROM golang:latest
#
#ENV GOPROXY https://goproxy.cn,direct
#WORKDIR $GOPATH/src/github.com/cncamp/works/http_server
#COPY . $GOPATH/src/github.com/cncamp/works/http_server
#RUN go build .
#
#EXPOSE 8088
#ENTRYPOINT ["./http_server"]


FROM golang:1.17.1-alpine as builder
WORKDIR $GOPATH/src/github.com/cncamp/works/http_server
ENV GOPROXY=https://goproxy.cn
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/http_server

FROM scratch
# -ldflags "-s -w" 去除构建包的调试信息
COPY --from=builder /bin/http_server /opt/app/http_server
EXPOSE 8088
CMD ["/opt/app/http_server"]
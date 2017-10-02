FROM golang:alpine

WORKDIR /go/src/github.com/gladmo/wechat
COPY . .

RUN go-wrapper install && \
	ln -s /go/src/github.com/gladmo/wechat/conf.yaml /go/bin/conf.yaml && \
	ln -s /go/src/github.com/gladmo/wechat/public /go/bin/public

EXPOSE 2222
CMD ["wechat", "spider", "lengtoo"]
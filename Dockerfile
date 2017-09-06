FROM golang:alpine

WORKDIR /go/src/github.com/gladmo/wechat
COPY . .

RUN apk add --no-cache git && \
	go-wrapper install && \
	apk del git

CMD ["main"] # ["app"]
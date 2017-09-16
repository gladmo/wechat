# wechat
wechat spider is a WeChat public number spider, you can write other rule to crawl you want things.

# Features
- [x] Crawl weChat public number lengtoo, parse Graphic Jokes and Text Jokes.
- [x] Verify Code server, When a verification code appears, start the service waiting for verification.
- [x] Notice to dingTalk.
- [ ] TODO

# Usage

```go
	go get github.com/gladmo/wechat

	cd $GOROOT/src/github.com/gladmo/wechat

	# import db struct from jokes.sql

	# modify conf.local to conf.yaml and set you dsn and phantomjs

	# docker run --network phantomjs --restart always -p 8910:8910 -d wernight/phantomjs phantomjs --webdriver=8910

	glide up

	# run in local
	go run main.go spider lengtoo

	# run in docker
	# 1. build
	docker build -t wechat-spider:latest -f Dockerfile .

	# 2. run if phantomjs run in you local, add --network phantomjs flag
	docker run --name wechat-spider --network phantomjs -v "$PWD"/public:/go/bin/public -d wechat-spider:latest
```

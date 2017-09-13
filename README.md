# wechat
wechat spider is a WeChat public number spider, you can write other rule to crawl you want things.

# Features
- [x] Crawl weChat public number lengtoo, parse Graphic Jokes and Text Jokes.
- [x] Verify Code server, When a verification code appears, start the service waiting for verification.
- [x] Notice to dingTalk.
- [ ] TODO

# Usage

```go
	go get github.com/gladmo/wehat

	# import sql from createdb/jokes.sql

	# modify conf/databases.local to databases.yaml and set you dsn

	# docker run --restart always -p 8910:8910 -d wernight/phantomjs phantomjs --webdriver=8910

	glide up

	go run main.go spider lengtoo
```

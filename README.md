# wechat
wechat spider is a WeChat public number spider, you can write other rule to crawl you want things.

# Features
- [x] Crawl weChat public number lengtoo, parse Graphic Jokes and Text Jokes.
- [x] Verify Code server, When a verification code appears, start the service waiting for verification.
- [x] Notice to dingTalk.
- [ ] TODO

# Usage
docker run --restart always -p 8910:8910 -d wernight/phantomjs phantomjs --webdriver=8910

```go
	go get github.com/gladmo/wehat

	glide up

	go run main.go spider lengtoo
```

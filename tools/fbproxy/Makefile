all: fbproxy.tar.gz

fbproxy.tar.gz: Dockerfile main.go go.mod go.sum
	tar -zcf $@ $?

.PHONY: all

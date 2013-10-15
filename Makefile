export GOPATH=$(shell pwd)
test:
	go test -v goblin

test-gomega:
	go get github.com/onsi/gomega
	go test -v gomega



export GOPATH=$(shell pwd)
test:
	go get github.com/onsi/gomega
	go test -v goblin

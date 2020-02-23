GOPATH:=$(shell go env GOPATH)

.PHONY: build
build_xprocess:
	go build -o bin/xprocess-server
build_xprocess_linux:
	GOOS=linux GOARCH=amd64 go build -o bin/xprocess-server
.PHONY: docker
docker:
	docker build . -t registry.cn-shanghai.aliyuncs.com/chenwentao/xprocess-backend:0.0.9
.PHONY: all clean build docker docker-build docker-run

GORUN = go run
GOBUILD = go build
APPNAME = accounts

all: clean build

proto-build:
	rm -rf proto/*.go
	protoc -I ./proto ./proto/$(APPNAME).proto --go_out=plugins=grpc:proto

clean:
	rm -rf bin

build:
	$(GOBUILD) -o bin/$(APPNAME) cmd/$(APPNAME)/*.go


docker-build:
	test $(DOCKERREPO)
	docker build . -t $(DOCKERREPO)

docker-push:
	test $(DOCKERREPO)
	docker push $(DOCKERREPO)

docker: docker-build docker-push

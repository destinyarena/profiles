.PHONY: all clean build docker docker-build docker-run

GORUN = go run
GOBUILD = go build
APPNAME = profiles

all: clean build

proto-build:
	rm -rf proto/*.go
	protoc --go_out=. --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		   ./proto/profiles.proto

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

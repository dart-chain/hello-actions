APP_NAME 	:= hello-actions
VERSION 	?= 

.Phony: build-docker run test clean docker

default: build

check-version:
	@if [ -z "${VERSION}" ]; then echo "VERSION is not set"; exit 1; fi

build-docker: check-version
	docker build -t $(APP_NAME):${VERSION} . --build-arg VERSION=${VERSION}
	docker tag $(APP_NAME):${VERSION} $(APP_NAME):latest

build: check-version
	go build -o $(APP_NAME) -ldflags "-X main.version=${VERSION}" cmd/main.go
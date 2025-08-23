APP_NAME := telegram-bot-go
VERSION := $(shell git describe --tags --always)
ARCH := amd64
OS := linux
IMAGE := ghcr.io/stanley29/$(APP_NAME):v1.0.0-$(VERSION)-$(OS)-$(ARCH)

build:
	go build -o bot ./main.go

docker-build:
	docker build -t $(IMAGE) .

docker-push:
	docker push $(IMAGE)

all: build docker-build docker-push

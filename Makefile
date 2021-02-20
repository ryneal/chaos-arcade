# Makefile for releasing chaos-arcade
#
# The release version is controlled from pkg/version

TAG?=latest
NAME:=chaos-arcade
DOCKER_REPOSITORY:=ryneal
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
GIT_COMMIT:=$(shell git describe --dirty --always)
VERSION:=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')
EXTRA_RUN_ARGS?=

run:
	go run -ldflags "-s -w -X github.com/ryneal/chaos-arcade/pkg/version.REVISION=$(GIT_COMMIT)" cmd/chaos-arcade/* \
	--level=debug --grpc-port=9999 --backend-url=https://httpbin.org/status/401 --backend-url=https://httpbin.org/status/500 \
	--ui-logo=https://raw.githubusercontent.com/ryneal/chaos-arcade/gh-pages/cuddle_clap.gif $(EXTRA_RUN_ARGS)

test:
	go test -v -race ./...

build:
	GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/ryneal/chaos-arcade/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/chaos-arcade ./cmd/chaos-arcade/*
	GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/ryneal/chaos-arcade/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/podcli ./cmd/podcli/*

fmt:
	gofmt -l -s -w ./
	goimports -l -w ./

build-charts:
	helm lint charts/*
	helm package charts/*

build-container:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .

build-base:
	docker build -f Dockerfile.base -t $(DOCKER_REPOSITORY)/chaos-arcade-base:latest .

push-base: build-base
	docker push $(DOCKER_REPOSITORY)/chaos-arcade-base:latest

test-container:
	@docker rm -f chaos-arcade || true
	@docker run -dp 9898:9898 --name=chaos-arcade $(DOCKER_IMAGE_NAME):$(VERSION)
	@docker ps
	@TOKEN=$$(curl -sd 'test' localhost:9898/token | jq -r .token) && \
	curl -sH "Authorization: Bearer $${TOKEN}" localhost:9898/token/validate | grep test

push-container:
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_IMAGE_NAME):latest
	docker push $(DOCKER_IMAGE_NAME):$(VERSION)
	docker push $(DOCKER_IMAGE_NAME):latest
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) quay.io/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) quay.io/$(DOCKER_IMAGE_NAME):latest
	docker push quay.io/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker push quay.io/$(DOCKER_IMAGE_NAME):latest

version-set:
	@next="$(TAG)" && \
	current="$(VERSION)" && \
	sed -i '' "s/$$current/$$next/g" pkg/version/version.go && \
	sed -i '' "s/tag: $$current/tag: $$next/g" charts/chaos-arcade/values.yaml && \
	sed -i '' "s/tag: $$current/tag: $$next/g" charts/chaos-arcade/values-prod.yaml && \
	sed -i '' "s/appVersion: $$current/appVersion: $$next/g" charts/chaos-arcade/Chart.yaml && \
	sed -i '' "s/version: $$current/version: $$next/g" charts/chaos-arcade/Chart.yaml && \
	sed -i '' "s/chaos-arcade:$$current/chaos-arcade:$$next/g" kustomize/deployment.yaml && \
	sed -i '' "s/chaos-arcade:$$current/chaos-arcade:$$next/g" deploy/webapp/frontend/deployment.yaml && \
	sed -i '' "s/chaos-arcade:$$current/chaos-arcade:$$next/g" deploy/webapp/backend/deployment.yaml && \
	sed -i '' "s/chaos-arcade:$$current/chaos-arcade:$$next/g" deploy/bases/frontend/deployment.yaml && \
	sed -i '' "s/chaos-arcade:$$current/chaos-arcade:$$next/g" deploy/bases/backend/deployment.yaml && \
	echo "Version $$next set in code, deployment, chart and kustomize"

release:
	git tag $(VERSION)
	git push origin $(VERSION)

swagger:
	go get github.com/swaggo/swag/cmd/swag
	cd pkg/api && $$(go env GOPATH)/bin/swag init -g server.go
include Makefile.var

LOCAL_ARCH := $(shell uname -m)
ifeq ($(LOCAL_ARCH),x86_64)
    ARCH ?= amd64
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 5),armv8)
    ARCH ?= arm64
else ifeq ($(LOCAL_ARCH),aarch64)
    ARCH ?= arm64
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 4),armv)
    ARCH ?= arm
else
    $(error This system's architecture $(LOCAL_ARCH) isn't supported)
endif

DOCKER_BUILD_FLAGS ?=

goimports := golang.org/x/tools/cmd/goimports@v0.1.5
golangci_lint := github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.0

.PHONY: test
test:
	go test $(shell go list ./...) ${opt}

.PHONY: run
run:
	go run main.go

.PHONY: swag.update
swag.update:
	swag init

.PHONY: lint
lint:
	@go run $(golangci_lint) run


build:
	docker build $(DOCKER_BUILD_FLAGS) \
		--build-arg HUB=$(HUB) \
		--build-arg VERSION=$(VERSION) \
		--build-arg PROD_NAME=$(PROD_NAME) \
		--build-arg GOPROXY=$(GOPROXY) \
		--build-arg GOSUMDB=$(GOSUMDB) \
		--build-arg ARCH=$(ARCH) \
		-t $(HUB)/$(PROD_NAME)-api-service:$(VERSION) \
		-f Dockerfile . ; \

.PHONY: build

release: build
	docker push $(HUB)/$(PROD_NAME)-api-service:$(VERSION)

.PHONY: release

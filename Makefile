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

.PHONY: mysql.up
mysql.up:
	docker-compose up -d --build

.PHONY: mysql.down
mysql.down:
	docker-compose down -v

.PHONY: lint
lint:
	@go run $(golangci_lint) run

.PHONY: format
format:
	@find . -type f -name '*.go' | xargs gofmt -s -w
	@for f in `find . -name '*.go'`; do \
	    awk '/^import \($$/,/^\)$$/{if($$0=="")next}{print}' $$f > /tmp/fmt; \
	    mv /tmp/fmt $$f; \
	done
	@go run $(goimports) -w -local github.com/tetratelabs/proxy-wasm-go-sdk `find . -name '*.go'`

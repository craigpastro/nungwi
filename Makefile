.PHONY: buf-mod-update
buf-mod-update:
	@test -s ./proto/buf.lock || buf mod update proto

.PHONY: buf-format
buf-format: buf-mod-update
	buf format -w

.PHONY: buf-generate
buf-generate: buf-format
	buf generate

.PHONY: lint
lint: buf-generate
	golangci-lint run

.PHONY: test
test: buf-generate
	go test -v ./...

.PHONY: build
build: buf-generate
	CGO_ENABLED=0 go build -o ./bin/nungwi ./cmd/nungwi

.PHONY: run
run: build
	./bin/nungwi

WIRECMD=wire
GOCMD=go

.PHONY: wire run fmt tests

wire:
	@command -v wire >/dev/null 2>&1 || $(GOCMD) install github.com/google/wire/cmd/wire@latest
	@cd internal/infrastructure && $(WIRECMD)

mocks-download:
	$(GOCMD) mod download
	$(GOCMD) install -mod=mod go.uber.org/mock/mockgen@latest

mocks-gen: mocks-download 
	@~/go/bin/mockgen -source=internal/infrastructure/http/client.go -destination=internal/usecase/mocks/client.go -typed=true -package=mock

fmt:
	go fmt ./...

test-clean: fmt
	$(GOCMD) clean -testcache

tests: fmt test-clean
	$(GOCMD) test -cover -p=1 ./...

run:
	cd cmd && $(GOCMD) run main.go --url=http://google.com --requests=20 --concurrency=10
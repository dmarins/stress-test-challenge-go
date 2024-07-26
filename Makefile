WIRECMD=wire
GOCMD=go

.PHONY: wire run

wire:
	@command -v wire >/dev/null 2>&1 || $(GOCMD) install github.com/google/wire/cmd/wire@latest
	@cd cmd/api && $(WIRECMD)

run:
	cd cmd && $(GOCMD) run main.go wire_gen.go

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
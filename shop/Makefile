BINDIR ?= ./bin
IMAGE_REPOSITORY ?= "johz/shop"
IMAGE_TAG ?= "dev"

GOOS ?= linux
GOARCH ?= arm

### function for get build os and arch args
### param: $(1) OS
### param: $(2) ARCH
go_build_args = GOOS=$(1) GOARCH=$(2)

go_test_args = GOARCH=$(1)

.PHONY: help
help:  ## Display this help message.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build Binary
	@$(call go_build_args,$(GOOS),$(GOARCH)) CGO_ENABLED=0 go build -o $(BINDIR)/shop .

.PHONY: lint
lint:  ## Apply go lint check
	@$(call go_test_args,$(GOARCH)) golangci-lint run --timeout 10m ./...

.PHONY: test
test:  ## Run unit test cases
	@$(call go_test_args,$(GOARCH)) go test -count=1 ./... -ldflags="-s=false" -gcflags="-l" --cover

.PHONY: build-image
build-image: build
	@docker build -t $(IMAGE_REPOSITORY):$(IMAGE_TAG) .
	@rm ./bin/shop
PACKAGE_NAME := "play-with-locust"
GO_FILES_ALL := $(shell find . -name '*.go' | grep -v /vendor/)

BUILD_IMAGE     ?= alittleprogramming/go-dev:1.14.1
ADDTL_LINTERS   := "goconst,gofmt,goimports,unparam"
SKIP_DIRS_LINTERS := "client,vendor,dist"

.PHONY: fmt
fmt:
	@docker run                                                 \
	    -i                                                      \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd)/servers:/src                                         \
	    -w /src                                                 \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    --env CGO_ENABLED=0                     		        \
	    --env GO111MODULE=on                                    \
	    --env GOFLAGS="-mod=vendor"                             \
	    $(BUILD_IMAGE)                                          \
	    hack/go-dev.sh fmt $(GO_FILES_ALL)

.PHONY: lint
lint:
	@docker run                                                 \
	    -i                                                      \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd)/servers:/src                                         \
	    -w /src                                                 \
	    -v $$(pwd)/.go-build/cache:/.cache                      \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    --env GO111MODULE=on                                    \
	    --env GOFLAGS="-mod=vendor"                             \
	    $(BUILD_IMAGE)                                          \
		hack/go-dev.sh lint $(ADDTL_LINTERS) $(SKIP_DIRS_LINTERS)


.PHONY: build
build: fmt
	@GO111MODULE=on \
	GOFLAGS="-mod=vendor" \
	CGO_ENABLED=0 \
	go build -o ./servers/$(PACKAGE_NAME) .

DISPATCHER_PAIR_LIMIT 		?= 5
DISPATCHER_ALLOCATE_DRIVER	?= false

.PHONY: run
run: build
	@./servers/$(PACKAGE_NAME)										\
			--pair-limit=$(DISPATCHER_PAIR_LIMIT)			\
			--allocate-driver=$(DISPATCHER_ALLOCATE_DRIVER)

.PHONY: clean
clean: ## Remove previous build
	@rm -f ./servers/$(PACKAGE_NAME)
	@rm -rf .go-build

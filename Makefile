SHELL = /bin/bash
GO-VER = go1.12

default: test

# #### GO Binary Management ####
deps-go-binary:
	echo "Expect: $(GO-VER)" && \
		echo "Actual: $$(go version)" && \
	 	go version | grep $(GO-VER) > /dev/null


HAS_GO_IMPORTS := $(shell command -v goimports;)

deps-goimports: deps-go-binary
ifndef HAS_GO_IMPORTS
	go get -u golang.org/x/tools/cmd/goimports
endif

# #### CLEAN ####
clean: deps-go-binary
	rm -rf build/*
	go clean --modcache


# #### DEPS ####

deps: deps-goimports deps-go-binary
	go mod download

test: deps lint
	ginkgo -r --randomizeAllSpecs --randomizeSuites --race --trace .

lint: deps-goimports
	go vet
	git ls-files | grep '.go$$' | xargs goimports -l -w

#!/bin/bash

lint() {
    LINTERS_LIST=$1
    SKIP_DIRS_LINTERS=$2
    echo "=========================="
	echo "|== Running linter   : ==|"
	echo "=========================="
    golangci-lint run \
        --enable ${LINTERS_LIST} \
        --timeout=10m \
        --skip-files=.*\\.my\\.go$,lib/bad.go,generated.*\.go$ \
        --skip-dirs-use-default \
        --skip-dirs=${SKIP_DIRS_LINTERS}
}

fmt() {
    echo "=========================="
	echo "|== Running goimports: ==|"
	echo "=========================="
	goimports -w $@
	echo

	echo "=========================="
	echo "|== Running gofmt:     ==|"
	echo "=========================="
	gofmt -s -w $@
	echo
}

"$@"

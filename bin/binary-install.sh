#!/usr/bin/env bash

export GOBIN=$(pwd)
export GOPATH=$(pwd)/..
/usr/local/go/bin/go install $(pwd)/../src/domhauton.com/membranecli/membrane.go
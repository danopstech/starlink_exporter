// +build tools

package main

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

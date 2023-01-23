//go:build tools
// +build tools

package mendix

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/daixiang0/gci"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "golang.org/x/tools/cmd/godoc"
)

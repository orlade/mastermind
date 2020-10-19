SYSLGO_SYSL=specs/mastermind.sysl
SYSLGO_PACKAGES=Mastermind
SYSLGO_APP.Mastermind = Mastermind

-include local.mk
include codegen.mk

all: external_types

.PHONY: external_types
external_types:
	cp internal/github/external_types.go gen/pkg/servers/Mastermind/github/external_types.go

tidy:
	go mod tidy
	gofmt -s -w .
	goimports -w .

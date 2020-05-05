SHELL += -eu

BLUE  := \033[0;34m
GREEN := \033[0;32m
RED   := \033[0;31m
NC    := \033[0m

GO111MODULE := on
GOPATH ?= ${HOME}/go
GO_BIN := ${GOPATH}/bin
GOPRIVATE := github.com/AlpacaLabs

include makefiles/*.mk
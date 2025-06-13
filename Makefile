SHELL=/bin/bash -o pipefail

# Define a new Path variable into current executed subshell
export PWD := $(shell pwd)


go.mod:
	@go mod download


.PHONY: all
all:  go.mod ## Render all
	@go run main.go

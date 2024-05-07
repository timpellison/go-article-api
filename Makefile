# sort of reminiscient of
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")

.PHONY: hello
hello:
	echo "Hello!"

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)
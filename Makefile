SHELL := /bin/bash

.PHONY: all \
	lint lint-go lint-js \
	test test-go test-js

all:
	@make test -j8

lint-go:
	@lint-go

lint-js:
	@lint-js

lint: lint-go lint-js

test-go: lint-go
	@test-go

test-js: lint-js
	@test-js

test: test-js test-go
	@echo "fanfare! all tests passed!"

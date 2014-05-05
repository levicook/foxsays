SHELL := /bin/bash

SRC_JS      := $(shell find src -type f -name '*.js')
SRC_PAGE_JS := $(shell find src/*/pages -type f -name 'main.js')
PKG_PAGE_JS := $(SRC_PAGE_JS:src/%.js=pkg/%.js)

.PHONY: all clean dist pkg pkg-less

all:
	@make clean
	@make dist -j8

clean:
	@rm -fr dist/**
	@rm -fr pkg/**

dist: pkg
	@find dist -type f -name '*.js' | xargs rm -f
	@gulp dist

pkg: pkg-less pkg/vendor-main.js $(PKG_PAGE_JS)

pkg/vendor-main.js:
	@gulp pkg:vendor-main.js

pkg-less:
	@gulp pkg:less

pkg/%.js: $(SRC_PAGE_JS)
	@mkdir -p $(@D)
	@browserify -t rfileify $(subst pkg,src,$@) -o $@

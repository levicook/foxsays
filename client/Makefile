SHELL := /bin/bash

SRC_JS        := $(shell find src -type f -name '*.js')
SRC_PAGE_JS   := $(shell find src/*/pages -type f -name 'main.js')
PKG_PAGE_JS   := $(SRC_PAGE_JS:src/%.js=pkg/%.js)

SRC_PAGE_LESS := $(shell find src/*/pages -type f -name 'main.less')
DIST_PAGE_CSS := $(SRC_PAGE_LESS:src/%.less=dist/assets/%.css)

.PHONY: all clean dist dist-css pkg

all:
	@make clean
	@make dist -j8

clean:
	@rm -fr dist/**
	@rm -fr pkg/**

dist: pkg dist-css
	@find dist -type f -name '*.js' | xargs rm -f
	@gulp dist

###pkg: pkg-css pkg/vendor-main.js $(PKG_PAGE_JS)
pkg: pkg/vendor-main.js $(PKG_PAGE_JS)

pkg/vendor-main.js:
	@gulp pkg:vendor-main.js

pkg/%.js: $(SRC_PAGE_JS)
	@mkdir -p $(@D)
	@browserify -t rfileify $(subst pkg,src,$@) -o $@

dist-css: $(DIST_PAGE_CSS)

dist/%.css: $(SRC_PAGE_LESS)
	@lessc $(subst dist/assets,src,$(subst css,less,$@)) $@

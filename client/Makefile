SHELL := /bin/bash

SRC_JS        := $(shell find src -type f -name '*.js')
SRC_LESS      := $(shell find src -type f -name '*.less')

SRC_PAGE_LESS := $(shell find src/*/pages -type f -name 'main.less')
DST_PAGE_CSS  := $(SRC_PAGE_LESS:src/%.less=dist/assets/%.min.css)

SRC_PAGE_JS   := $(shell find src/*/pages -type f -name 'main.js')
DST_PAGE_JS   := $(SRC_PAGE_JS:src/%.js=dist/assets/%.min.js)

#-------------

.PHONY: all clean dist dist-css dist-js

all:
	@make clean
	@make dist -j8

clean:
	@rm -fr dist/**

#-------------

dist: dist-css dist-js
	@gulp dist
	@find dist/assets -type f -name '*.min.js' -o -name '*.min.css' | xargs rm -f
	@rewrite-asset-paths

dist-css: $(DST_PAGE_CSS)

dist-js: $(DST_PAGE_JS)

dist/%.css: $(SRC_LESS)
	@lessc --compress $(subst .min.css,.less,$(subst dist/assets,src,$@)) $@

dist/%.js: $(SRC_PAGE_JS)
	@mkdir -p $(@D)
	@browserify --ig -t rfileify $(subst .min,,$(subst dist/assets,src,$@)) | uglifyjs --compress > $@

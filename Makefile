.SUFFIXES:

# Determine whether we're being verbose or not
export V = false
export CMD_PREFIX   = @
export NULL_REDIR = 2>/dev/null >/dev/null
ifeq ($(V),true)
	CMD_PREFIX =
	NULL_REDIR =
endif

# Default build type is debug.
TYPE ?= debug

# We ignore the ".map" file in release mode, to keep the size of our binary
# to a minimum.  In debug mode, we also always load the assets from disk.
ifeq ($(TYPE),release)
export ASSET_FLAGS   :=
export BINDATA_FLAGS := "-ignore=.*\\.map"
export WEBPACK_FLAGS := --optimize-minimize --optimize-dedupe
export NODE_ENV      := NODE_ENV=production
else
export ASSET_FLAGS   := --debug
export BINDATA_FLAGS := -debug
export WEBPACK_FLAGS :=
export NODE_ENV      :=
endif

RED     := \e[0;31m
GREEN   := \e[0;32m
YELLOW  := \e[0;33m
NOCOLOR := \e[0m

WEBPACK_BIN  := ./node_modules/webpack/bin/webpack.js

JS_FILES     := $(shell find frontend/js/ -name '*.js' -or -name '*.jsx')
STATIC_FILES := $(shell find frontend/css -name '*.css') \
				$(shell find frontend/fonts -type f)
BUILD_FILES  := $(patsubst frontend/%,build/%,$(STATIC_FILES))

# The final list of resources to embed in our binary.
RESOURCES    := build/index.html \
				build/js/bundle.js \
				build/js/bundle.js.map \
				$(BUILD_FILES)

######################################################################

.PHONY: all
all: dependencies build/dashboard


build/dashboard: descriptors.go resources.go *.go
	@printf "  $(GREEN)GO$(NOCOLOR)       $@\n"
	$(CMD_PREFIX)godep go build -o $@ .

descriptors.go: $(RESOURCES)
	@printf "  $(GREEN)ASSETS$(NOCOLOR)   $@\n"
	$(CMD_PREFIX)python ./gen_descriptors.py \
		$(ASSET_FLAGS) \
		--prefix "build/" \
		$(RESOURCES) > $@

resources.go: $(RESOURCES)
	@printf "  $(GREEN)BINDATA$(NOCOLOR)  $@\n"
	$(CMD_PREFIX)go-bindata \
		-ignore='^.*(\.gitignore|dashboard)$$' \
		$(BINDATA_FLAGS) \
		-prefix "./build" \
		-o $@ \
		$(sort $(dir $^))

######################################################################

build/js/bundle.js: $(JS_FILES)
	@printf "  $(GREEN)WEBPACK$(NOCOLOR)  $@\n"
	$(CMD_PREFIX)$(NODE_ENV) node $(WEBPACK_BIN) \
		--progress --colors $(WEBPACK_FLAGS) $(NULL_REDIR)

build/index.html: frontend/index.html
	@printf "  $(GREEN)CP$(NOCOLOR)       $< ==> $@\n"
	$(CMD_PREFIX)cp $< $@

build/css/%: frontend/css/%
	@printf "  $(GREEN)CP$(NOCOLOR)       $< ==> $@\n"
	@mkdir -p $(dir $@)
	$(CMD_PREFIX)cp $< $@

build/fonts/%: frontend/fonts/%
	@printf "  $(GREEN)CP$(NOCOLOR)       $< ==> $@\n"
	@mkdir -p $(dir $@)
	$(CMD_PREFIX)cp $< $@

######################################################################

# This is a phony target that checks to ensure our various dependencies are installed
.PHONY: dependencies
dependencies:
	@command -v go-bindata >/dev/null 2>&1 || { printf >&2 "go-bindata is not installed, exiting...\n"; exit 1; }
	@command -v godep      >/dev/null 2>&1 || { printf >&2 "godep is not installed, exiting...\n"; exit 1; }
	@command -v node       >/dev/null 2>&1 || { printf >&2 "node.js is not installed, exiting...\n"; exit 1; }
	@test -d node_modules/webpack    || { printf >&2 "npm dependencies not satisfied, exiting...\n"; exit 1; }
	@# Since webpack doesn't seem to exit with an error if this isn't present...
	@test -d node_modules/jsx-loader || { printf >&2 "npm dependencies not satisfied, exiting...\n"; exit 1; }

######################################################################

.PHONY: clean
CLEAN_FILES := build/dashboard resources.go descriptors.go $(RESOURCES)
clean:
	@printf "  $(YELLOW)RM$(NOCOLOR)       $(CLEAN_FILES)\n"
	$(CMD_PREFIX)$(RM) -r $(CLEAN_FILES)

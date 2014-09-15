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
else
export ASSET_FLAGS   := --debug
export BINDATA_FLAGS := -debug
endif

RED     := \e[0;31m
GREEN   := \e[0;32m
YELLOW  := \e[0;33m
NOCOLOR := \e[0m

# Input resources to embed in our binary.
RESOURCES :=

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
		-ignore=\\.gitignore \
		$(BINDATA_FLAGS) \
		-prefix "./build" \
		-o $@ \
		$(sort $(dir $^))

# This is a phony target that checks to ensure our various dependencies are installed
.PHONY: dependencies
dependencies:
	@command -v go-bindata >/dev/null 2>&1 || { printf >&2 "go-bindata is not installed, exiting...\n"; exit 1; }
	@command -v godep      >/dev/null 2>&1 || { printf >&2 "godep is not installed, exiting...\n"; exit 1; }

######################################################################

.PHONY: clean
CLEAN_FILES := build/dashboard resources.go descriptors.go
clean:
	@printf "  $(YELLOW)RM$(NOCOLOR)       $(CLEAN_FILES)\n"
	$(CMD_PREFIX)$(RM) -r $(CLEAN_FILES)

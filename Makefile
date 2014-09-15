.SUFFIXES:

.PHONY: all
all: build/dashboard


build/dashboard: *.go
	godep go build -o $@ $^

#!/usr/bin/make -f

SHELL = /bin/bash
#.SHELLFLAGS = -ecx
.SHELLFLAGS = -ec

GO ?= go

default: build
PACKAGE = github.com/stephram/audioinfo

APP_NAME = audioinfo
PRODUCT = audioinfo

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD\#\#*/})

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

HOME = $(shell echo $${HOME})

BUILD_FOLDER = $(shell echo `pwd`/build)

# build variables
BRANCH_NAME     ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE      ?= $(shell date '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT      ?= $(shell git rev-list -1 HEAD)
VERSION         ?= 0.0.1
AUTHOR          ?= $(shell git log -1 --pretty=format:'%an')
AUTHOR_EMAIL    ?= $(shell git log -1 --pretty=format:'%ae')

BUILD_OVERRIDES = \
	-X "$(PACKAGE)/pkg/app.Name=$(APP_NAME)" \
	-X "$(PACKAGE)/pkg/app.Product=$(PRODUCT)" \
	-X "$(PACKAGE)/pkg/app.Branch=$(BRANCH_NAME)" \
	-X "$(PACKAGE)/pkg/app.BuildDate=$(BUILD_DATE)" \
	-X "$(PACKAGE)/pkg/app.Commit=$(GIT_COMMIT)" \
	-X "$(PACKAGE)/pkg/app.Version=$(VERSION)" \
	-X "$(PACKAGE)/pkg/app.Author=$(AUTHOR)" \
	-X "$(PACKAGE)/pkg/app.AuthorEmail=$(AUTHOR_EMAIL)" \

info:
	@echo "HOME            : $(HOME)"
	@echo "APP_NAME        : $(APP_NAME)"
	@echo "PRODUCT         : $(PRODUCT)"
	@echo "BRANCH_NAME     : $(BRANCH_NAME)"
	@echo "BUILD_DATE      : $(BUILD_DATE)"
	@echo "GIT_COMMIT      : $(GIT_COMMIT)"
	@echo "VERSION         : $(VERSION)"
	@echo "BUILD_FOLDER    : $(BUILD_FOLDER)"
	@echo "AUTHOR          : $(AUTHOR)"
	@echo "AUTHOR_EMAIL    : $(AUTHOR_EMAIL)"
	@echo "TARGET          : $(TARGET)"
	@echo "SRC             : $(SRC)"
	@echo "BUILD_OVERRIDES : $(BUILD_OVERRIDES)"

install:
	go get -u github.com/go-audio/wav
	go get -u github.com/sirupsen/logrus
	go get -u github.com/oklog/ulid
	# Install gometalinter
	curl -L https://git.io/vp6lP | sh

clean:
	rm -rf $(BUILD_FOLDER)

lint:
	gometalinter --vendor --deadline=2m ./cmd/... ./internal/...

$(TARGET) : $(SRC)
	CGO_ENABLED=0 GOARCH=amd64 \
		go build -a \
		-installsuffix cgo \
		-ldflags='-w -s $(BUILD_OVERRIDES)' \
		-o $(BUILD_FOLDER)/audioinfo cmd/audioinfo/main.go

build: $(TARGET)
	@true

run:
	@find $(HOME) -name [aA-zZ]*.wav -exec $(BUILD_FOLDER)/audioinfo {} \;

runJSON:
	@find $(HOME) -name [aA-zZ]*.wav -exec $(BUILD_FOLDER)/audioinfo {} \;

runTEXT:
	@find $(HOME) -name [aA-zZ]*.wav -exec $(BUILD_FOLDER)/audioinfo -h -ofmt=text {} \;

watch:
	@yolo -i . -e vendor -e build -c $(run)



#!/usr/bin/make -f

SHELL = /bin/bash
.SHELLFLAGS = -ecx

GO ?= go

default: build
PACKAGE = github.com/stephram/grypton

# build variables
BRANCH_NAME ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE  ?= $(shell date '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT  ?= $(shell git rev-list -1 HEAD)
VERSION     ?= 0.0.0

BUILD_OVERRIDES = \
	-X "$(PACKAGE)/pkg/app.Branch=$(BRANCH_NAME)" \
	-X "$(PACKAGE)/pkg/app.BuildDate=$(BUILD_DATE)" \
	-X "$(PACKAGE)/pkg/app.Commit=$(GIT_COMMIT)" \
	-X "$(PACKAGE)/pkg/app.Version=$(VERSION)" \

BUILD_FOLDER = $(shell echo `pwd`/build)

install:
	go get -u github.com/go-audio/wav
	go get -u github.com/sirupsen/logrus

clean:
	rm -rf $(BUILD_FOLDER)

build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
		go build -a \
		-installsuffix cgo \
		-ldflags='-w -s $(BUILD_OVERRIDES)' \
		-o $(BUILD_FOLDER)/audioinfo cmd/audioinfo/main.go

#	go build -o $(BUILD_FOLDER)/audioinfo cmd/audioinfo/main.go

run:
	find $(HOME) -name [aA-zZ]*.wav -exec $(BUILD_FOLDER)/audioinfo -h {} \;

runJSON:
	#find $(HOME) -name *.wav -exec $(BUILD_FOLDER)/audioinfo -ofmt=json {} \;
	find $(HOME) -name [aA-zZ]*.wav -exec $(BUILD_FOLDER)/audioinfo {} \;

runTEXT:
	find $(HOME) -name [aA-zZ]*.wav -exec $(BUILD_FOLDER)/audioinfo -h -ofmt=text {} \;





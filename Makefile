AUTOMAKE_OPTIONS = foreign

SHELL:=/bin/bash
.PHONY : help install

help:   ## Show this help message.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

go-get:  ## Go get all dependencies
	go get ./...

build:  ## Run go build (don't use this)
	GOPATH=${PWD} go build bcrypt-password-hash

install: ## Run go install
	GOPATH=${PWD} go install bcrypt-password-hash

info: ## Show bcrypt-pasword-hash How-To
	./bin/bcrypt-password-hash


go-get:
	go get ./...

build:
	GOPATH=${PWD} go build bcrypt-password-hash

install:
	GOPATH=${PWD} go install bcrypt-password-hash

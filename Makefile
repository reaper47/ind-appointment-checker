.ONESHELL:

build:
	go build -ldflags="-s -w" -o bin/ind main.go
	cp -n .env ./bin/.env

run:
	go run main.go serve

release:	
ifdef tag
		sh build.sh github.com/reaper47/ind-appointment-checker $(tag)
		gh release create $(tag) ./release/$(tag)/*
else
		@echo 'Add the tag argument, i.e. `make release tag=v1.0.0`'
endif

test:
	go test ./...
	
%:
	@: 

.PHONY: release build run test
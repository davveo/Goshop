# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=Eshop
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	@$(GOBUILD) -o $(BINARY_NAME) -v
clean:
	@$(GOCLEAN)
	@rm -f $(BINARY_NAME)
	@rm -f $(BINARY_UNIX)
migrate:clean build
	@./${BINARY_NAME}
run: clean build
	@./${BINARY_NAME}

push: clean
	git status && git commit -am"$(msg)" && git push

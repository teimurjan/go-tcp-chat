GOCMD=go
GOBUILD=$(GOCMD) build

SRC_ROOT=$(shell pwd)

SERVER_SRC=src/server/main/main.go
CLIENT_SRC=src/client/main/main.go

SERVER_OUT_NAME=srv
CLIENT_OUT_NAME=cli

build:
	$(GOBUILD) -o $(SERVER_OUT_NAME) ${GOPATH}/$(SERVER_SRC)
	$(GOBUILD) -o $(CLIENT_OUT_NAME) ${GOPATH}/$(CLIENT_SRC)

clean:
	@[ -f ./$(SERVER_NAME) ] && rm ./$(SERVER_NAME) || true
	@[ -f ./$(CLIENT_NAME) ] && rm ./$(CLIENT_NAME) || true

run:
	$(GOBUILD) -o $(SERVER_OUT_NAME) ${GOPATH}/$(SERVER_SRC)
	$(GOBUILD) -o $(CLIENT_OUT_NAME) ${GOPATH}/$(CLIENT_SRC)
	./srv ${ARGS}
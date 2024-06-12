.PHONY: all webserver install docs

all: install webserver
build: webserver
NUM_CPU_CORES := $(shell nproc --all)
DOUBLE_CPU_CORES = $(shell expr $(NUM_CPU_CORES) \* 2)

install:
	go get -a
	go mod download

webserver:
	go build -gcflags "all=-N -l" -tags dynamic -o webserver cmd/main.go

docs:
	swag init --generalInfo cmd/main.go

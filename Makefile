BUILD?= $(CURDIR)/bin
$(shell mkdir -p $(BUILD))
VERSION?= v$(shell git rev-list HEAD --count)
ARCH?= $(shell uname -m)
SYSTEM?= $(shell uname)
export CGO_ENABLED=0

.PHONY: mod
mod:
	go mod verify
	go mod vendor
	go mod tidy

.PHONY:	build
build:
	go build -o $(BUILD)/app $(CURDIR)/cmd

.PHONY: test
test:
	go test ./... -count=1

.PHONY: dockerbuild
dockerbuild:
	docker build -t app_${SYSTEM}_${ARCH}:${VERSION} $(CURDIR)/.

.PHONY: docker-compose-up
docker-compose-up:
	docker-compose -f $(CURDIR)/docker-compose.yml up -d

.PHONY: docker-compose-down
docker-compose-down:
	docker-compose -f $(CURDIR)/docker-compose.yml down -v

.PHONY: docker-scale
docker-scale:
	docker-compose scale app=${NUMBER}

.PHONY: clean
clean:
	go clean $(CURDIR)/cmd
	rm -rf $(BUILD)

.DEFAULT_GOAL := build

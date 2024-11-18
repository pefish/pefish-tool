
DEFAULT: build-cur

ifeq ($(GOPATH),)
  GOPATH = $(HOME)/go
endif

build-cur:
	GOPATH=$(GOPATH) go install github.com/pefish/go-build-tool/cmd/...@latest
	go mod tidy
	$(GOPATH)/bin/go-build-tool

install: build-cur
	sudo install -C ./build/bin/linux/pefish-tool /usr/local/bin/pefish-tool

install-service: install
	sudo mkdir -p /etc/systemd/system
	sudo install -C -m 0644 ./script/pefish-tool.service /etc/systemd/system/pefish-tool.service
	sudo systemctl daemon-reload
	@echo
	@echo "pefish-tool service installed."


# Makefile
#
# Required tooling:
#
# - docker
#

help:
	@echo
	@echo "  \033[34mtest            \033[0m       - run the project tests"
	@echo "  \033[34mdev             \033[0m       - starts the project"
	@echo
.PHONY: help

test:
	go test -v -cover ./...
.PHONY: test

dev:
	docker compose build
	docker compose up
.PHONY: dev
.DEFAULT_GOAL := help
.PHONY: up rebuild down ssh test test.coverage build logs docker-cleanup help lint vet fmt deps vendor doc

DOCKER_C := docker-compose
DOCKER_RUN := docker-compose run 
APP_NAME := app

OUTPUT_COVERAGE := pkg/tests/coverage/

up: ## Start docker container
	$(DOCKER_C) pull
	$(DOCKER_C) up -d

rebuild: ## Rebuild docker container
	$(DOCKER_C) pull
	$(DOCKER_C) up --build -d

down: ## Stop docker container
	$(DOCKER_C) down --remove-orphans

ssh: ## Interactive access to container
	$(DOCKER_RUN) --entrypoint="/bin/sh" $(APP_NAME)
	$(DOCKER_C) down --remove-orphans

deps: ## Install dependencies
	$(DOCKER_RUN) --entrypoint="go mod download" $(APP_NAME)
	$(DOCKER_C) down --remove-orphans

vendor: ## Install vendor
	$(DOCKER_RUN) --entrypoint="go mod vendor" $(APP_NAME)
	$(DOCKER_C) down --remove-orphans

lint: ## Checks Code Style
	$(DOCKER_RUN) --entrypoint="./run-lint.sh" $(APP_NAME)
	$(DOCKER_C) down --remove-orphans

vet: ## Finds issues in code
	$(DOCKER_RUN) --entrypoint="go vet ./..." $(APP_NAME)
	$(DOCKER_C) down --remove-orphans

fmt: ## Applies standard formatting
	$(DOCKER_RUN) --entrypoint="go fmt ./..." $(APP_NAME)
	$(DOCKER_C) down --remove-orphans

doc: ## Show package documentation
	$(DOCKER_RUN) --entrypoint="go doc github.com/jmsilvadev/cycloid/katas02" $(APP_NAME)
	$(DOCKER_RUN) --entrypoint="go doc github.com/jmsilvadev/cycloid/katas02_with_encapsulation" $(APP_NAME)
	$(DOCKER_RUN) --entrypoint="go doc github.com/jmsilvadev/cycloid/katas19" $(APP_NAME)
	$(DOCKER_C) down --remove-orphans

benchmark: ## Run benchmark tests
	$(DOCKER_RUN) --entrypoint="./run-bench.sh" $(APP_NAME)
	$(DOCKER_C) down --remove-orphans

test: ## Run all available tests
	$(DOCKER_RUN) --entrypoint="./run-tests.sh" $(APP_NAME)
	$(DOCKER_C) down --remove-orphans

test.coverage: ## Check project test coverage
	$(DOCKER_RUN) --entrypoint="./run-tests.sh" $(APP_NAME)
	open $(OUTPUT_COVERAGE)coverage.html >&- 2>&- || \
	xdg-open $(OUTPUT_COVERAGE)coverage.html >&- 2>&- || \
	gnome-open $(OUTPUT_COVERAGE)coverage.html >&- 2>&-
	$(DOCKER_C) down --remove-orphans

build: ## Build docker image in daemon mode
	$(DOCKER_C) build

logs: ## Watch docker log files
	$(DOCKER_C) logs --tail 100 -f

help:
	@grep -E '^[a-zA-Z._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

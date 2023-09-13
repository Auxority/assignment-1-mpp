.DEFAULT_GOAL := help

.PHONY: help
help: ## Show the available commands
	@printf "\033[33mUsage:\033[0m\n  make [target] [arg=\"val\"...]\n\n\033[33mTargets:\033[0m\n"
	@grep -E '^[-a-zA-Z0-9_\.\/]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[32m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the docker image
	@printf "\033[32mBuilding docker image...\033[0m\n"
	@docker compose build

.PHONY: run
run: ## Run the docker image
	@printf "\033[32mRunning docker image...\033[0m\n"
	@docker compose up -d

.PHONY: stop
stop: ## Stop the docker image
	@printf "\033[32mStopping docker image...\033[0m\n"
	@docker compose stop

.PHONY: clean
clean: ## Clean the docker image
	@printf "\033[32mCleaning docker image...\033[0m\n"
	@docker compose down --rmi all --volumes --remove-orphans

.PHONY: sh
sh: ## Run a shell in the docker image
	@printf "\033[32mRunning shell in docker image...\033[0m\n"
	@docker compose exec app sh

.PHONY: logs
logs: ## Show the logs of the docker image
	@printf "\033[32mShowing logs of docker image...\033[0m\n"
	@docker compose logs -f app

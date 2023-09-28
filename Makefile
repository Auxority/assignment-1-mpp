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

.PHONY: restart
restart: ## Restart the docker image
	@printf "\033[32mRestarting docker image...\033[0m\n"
	@docker compose restart

.PHONY: clean
clean: ## Clean the docker image
	@printf "\033[32mCleaning docker image...\033[0m\n"
	@docker compose down --rmi all --volumes --remove-orphans

.PHONY: sh
sh: ## Run a shell in the docker image
	@printf "\033[32mRunning shell in docker image...\033[0m\n"
	@docker compose exec mpp sh

.PHONY: logs
logs: ## Show the logs of the docker image
	@printf "\033[32mShowing logs of docker image...\033[0m\n"
	@docker compose logs -f mpp --tail 5

.PHONY: test
test: ## Test all the commands
	@printf "\033[32mRunning list command...\033[0m\n"
	@docker compose exec mpp ./bin/main list

	@printf "\033[32mRunning add command...\033[0m\n"
	@docker compose exec mpp ./bin/main add -imdbid tt0000001 -title Carmencita -year 1894 -rating 5.7

	@printf "\033[32mRunning details command...\033[0m\n"
	@docker compose exec mpp ./bin/main details -imdbid tt0000001

	@printf "\033[32mRunning delete command...\033[0m\n"
	@docker compose exec mpp ./bin/main delete -imdbid tt0000001

	@printf "\033[32mRunning API list endpoint...\033[0m\n"
	@docker compose exec mpp wget -qO- localhost:8090/movies

	@printf "\033[32mRunning API details endpoint...\033[0m\n"
	@docker compose exec mpp wget -qO- localhost:8090/movies/tt0111161

	@printf "\033[32mRunning API add endpoint...\033[0m\n"
	@docker compose exec mpp curl -X POST -H "Content-Type: mpplication/json" -d '{"imdb_id": "tt0368226", "title": "The Room", "rating": 3.7, "year": 2003}' localhost:8090/movies

	@printf "\033[32mRunning API delete endpoint...\033[0m\n"
	@docker compose exec mpp curl -X DELETE localhost:8090/movies/tt0368226

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show help
	@egrep -h '^[a-zA-Z0-9_\/-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort -d | awk 'BEGIN {FS = ":.*?## "; printf "Usage: make \033[0;34mTARGET\033[0m \033[0;35m[ARGUMENTS]\033[0m\n\n"; printf "Targets:\n"}; {printf "  \033[33m%-25s\033[0m \033[0;32m%s\033[0m\n", $$1, $$2}'

psql: ## connect to db 
	docker exec -it graph-model psql -U admin -d graph

reset: ## reset db
	cat ./scripts/01_create_tables.sql | docker exec -i graph-model psql -U admin -d graph

run-watch: services-up
	@gow run .

services-up:
	@docker-compose -f infra/compose.yaml up -d

services-stop:
	@docker-compose -f infra/compose.yaml stop

services-down:
	@docker-compose -f infra/compose.yaml down

test-watch:
	@gow test ./...
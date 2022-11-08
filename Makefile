setup:
	docker-compose build
	docker-compose run --rm graphql sql-migrate up

start:
	docker-compose up

end:
	docker-compose down

entry-server-container:
	docker-compose exec graphql ash

entry-db-container:
	docker-compose exec db bash

test-cover:
	docker-compose exec graphql go test -cover ./... -coverprofile=cover.out
	docker-compose exec graphql go tool cover -html=cover.out -o cover.html
DOCKER_COMPOSE=docker compose -f docker/docker-compose.yml
container_name=docker-server-1
DOCKER_EXEC=docker exec -it ${container_name}

start:
	${DOCKER_COMPOSE} up --build -d 
	make migration-init
	make migration-up

stop:
	${DOCKER_COMPOSE} rm -s -v -f 

restart:
	${DOCKER_COMPOSE} restart

livereload:
	git ls-files | entr -c -r -s 'make restart; docker logs -f docker-server-1'

migration-init:
	${DOCKER_EXEC} go run ./cmd/migrate/main.go db init

migration-up:
	${DOCKER_EXEC} go run ./cmd/migrate/main.go db migrate

reset:
	make stop
	make start 

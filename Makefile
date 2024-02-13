
up:
	docker-compose up -d
down:
	docker-compose down
kill:
	docker-compose kill

psql:
	docker exec -it gochatdb psql -U root gochat

createdb:
	docker exec -it gochatdb createdb --username=root --owner=root gochat

dropdb:
	docker exec -it gochatdb dropdb gochat

.PHONY: dpsql createdb dropdb up down kill
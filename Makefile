up:
	docker-compose up -d
stop:
	docker-compose down
kill:
	docker-compose kill

psql:
	docker exec -it gochatdb psql -U root gochat

createdb:
	docker exec -it gochatdb createdb --username=root --owner=root gochat

dropdb:
	docker exec -it gochatdb dropdb gochat

migrate-up:
	cd server/ && migrate -path db/migrations -database "postgresql://root:password@localhost:5432/gochat?sslmode=disable" -verbose up && cd ..
migrate-down:
	cd server/ && migrate -path db/migrations -database "postgresql://root:password@localhost:5432/gochat?sslmode=disable" -verbose down && cd ..

rfe:
	cd client/ && npm run dev && cd ..
rbe:
	cd server/ && go run cmd/main.go && cd ..
	
.PHONY: dpsql createdb dropdb up down kill

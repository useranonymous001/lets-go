postgres:
	docker run --name simple_bank_cont -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb: 
	docker exec -it simple_bank_cont createdb --username=root --owner=root simple_bank_db

dropdb:
	docker exec -it simple_bank_cont dropdb simple_bank_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/root?sslmode=disable" -verbose up


migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/root?sslmode=disable" -verbose force 1 down



.PHONY: createdb dropdb migrateup migratedown

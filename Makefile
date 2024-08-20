# GO 
run:
	go run main.go

test:
	go test ./... -cover

# Database 
db.init:
	docker run -d --name argon2password -p 4444:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 postgres

db.create:
	docker exec -it argon2password createdb --username=root argon2password

db.drop:
	docker exec -it argon2password dropdb argon2password

# Migrations 
migrate.init:	
	migrate create -ext sql -dir db/migrations -seq create_users

migrate.up:
	migrate -path db/migrations -database 'postgres://root:123456@localhost:4444/argon2password?sslmode=disable' -verbose up

migrate.down:
	migrate -path db/migrations -database 'postgres://root:123456@localhost:4444/argon2password?sslmode=disable' -verbose down
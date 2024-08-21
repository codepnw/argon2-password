# GO 
run:
	go run main.go

test:
	go test ./... -cover


PQROOT=root
PQPWD=123456
DBNAME=argon2password
MIGRATEPATH=db/migrations
DSN=postgres://${PQROOT}:${PQPWD}@localhost:4444/${DBNAME}?sslmode=disable

# Database 
db.init:
	docker run -d --name ${DBNAME} -p 4444:5432 -e POSTGRES_USER=${PQROOT} -e POSTGRES_PASSWORD=${PQPWD} postgres

db.create:
	docker exec -it ${DBNAME} createdb --username=root ${DBNAME}

db.drop:
	docker exec -it ${DBNAME} dropdb ${DBNAME}

# Migrations 
migrate.init:	
	migrate create -ext sql -dir ${MIGRATEPATH} -seq create_users

migrate.up:
	migrate -path ${MIGRATEPATH} -database '${DSN}' -verbose up

migrate.down:
	migrate -path ${MIGRATEPATH} -database '${DSN}' -verbose down
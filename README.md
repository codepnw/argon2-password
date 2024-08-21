# Argon 2 Password Hash

## Package
[document](https://pkg.go.dev/github.com/alexedwards/argon2id@v1.0.0)
```bash
go get github.com/alexedwards/argon2id
```

## Run and Test

```bash
make run
```
```bash
make test
```

## Setup
```bash
# docker init postgres database
make db.init
# docker create database
make db.create
# docker drop database
make db.drop


# migrations init
make migrate.init
# migrate up
make migrate.up
# migrate down
make migrate.down
```
change docker and migrations at Makefile
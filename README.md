# pismo

## Application structure

```
├── account
│   ├── model.go
│   ├── repository.go
│   └── serializer.go
├── adapter
│   ├── api
│   │   ├── account.go
│   │   ├── operationtype.go
│   │   └── transaction.go
│   ├── repository
│   │   ├── mysqlconnection
│   │   │   ├── account.go
│   │   │   ├── general.go
│   │   │   ├── operationtype.go
│   │   │   └── transaction.go
│   │   └── sqlite
│   │       ├── account.go
│   │       ├── operationtype.go
│   │       └── transaction.go
│   └── serializer
│       ├── account.go
│       └── transaction.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── operationtype
│   ├── model.go
│   ├── repository.go
│   └── serializer.go
├── README.md
├── test.db
└── transaction
    ├── model.go
    ├── repository.go
    └── serializer.go

9 directories, 28 files
```

## Local development environment

The following commands are needed to build and run the application:

```
go build main.go

./main
```

## Working using docker

To build docker use:

```
docker-compose build --no-cache
```

To run the application use:

```
docker-compose up -d
```

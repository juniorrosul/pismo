# pismo

# Local development environment

The following commands are needed to build and run the application:

```
go build main.go

./main
```

# Working using docker

To build docker use:

```
docker build --no-cache -t github.com/juniorrosul/pismo-api .
```

To run the application use:

```
docker run -p 8080:8080 -d github.com/juniorrosul/pismo-api
```

# Fibonacci-Service-Go
Rest Service - Go - Fibonacci serie

----------------------------------------

Execute this as Go Application:
```
go run Fibonacci.go
```

----------------------------------------

Execute this as Docker Container: (With Docker Compose)
- Run Container:
```
docker-compose up
```

- Run Container as Daemon:
```
docker-compose up -d
```
----------------------------------------

And send a test curl:
```
curl http://localhost:3004/\?number\=10
curl http://localhost:3004/\?number\=11
curl http://localhost:3004/\?number\=12
curl http://localhost:3004/\?number\=13
curl http://localhost:3004/\?number\=[Nº]
```


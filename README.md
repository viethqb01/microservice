# Run service api-gateway

```shell
cd api-gateway 
docker composer up -d
```

# Run service registry

```shell
cd registry 
docker composer up -d
```

# Run service order-service

```shell
cd order-service  
go run main.go
```

# Run service user-service

```shell
cd user-service  
go run main.go
```
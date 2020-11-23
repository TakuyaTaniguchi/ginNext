# ginNext


## start

### Backend
```
docker-compose up -d
go run main.go
```

### Frontend
```
cd frontend
yarn install && yarn dev
```

## stop

```
docker ps
docker stop ${containerID}
```

## mysql
```
mysql -u root -p
```
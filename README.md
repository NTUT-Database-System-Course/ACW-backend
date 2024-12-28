# ACW-Backend

## how to run env
```bash
sudo docker-compose -f .devcontainer/docker-compose.yml up --build
```

You can open up http://localhost:8080/swagger/index.html to see api documentation

## clean up (including reset db server)
```bash
sudo docker-compose -f .devcontainer/docker-compose.yml down -v
```

## access to database directly when service is running
```
sudo docker exec -it devcontainer-db-1 psql -U root -d db
```

## generate docs
```
swag init --parseDependency github.com/guregu/null/v5
```

# Simple Gofiber GraphQL

## How to running using docker
```bash
# running all service (go service and postgres service)
docker-compose up -d --force-recreate

# running project go
docker-compose exec go go run .

```

## Running endpoint 
```
# endpoint add employee 
curl -X POST \
    -H 'Accept: application/json' \
    -H "Content-Type: application/json" \
    -d '{"query": "mutation { addEmployee(name: \"John Doe\", position: \"Software Engineer\", address: \"Jakarta\") { id name position address } }"}' \
    http://localhost:8222/graphql


# endpoint find employee by id
curl -X POST \
    -H 'Accept: application/json' \
    -H "Content-Type: application/json" \
    -d '{"query": "query { employee(id: \"1ea9945e-6595-464b-a091-23f36dabcb47\") { id name position address } }"}' \
    http://localhost:8222/graphql
```


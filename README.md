
// Start Docker Container
docker run --name=mkp-db -e POSTGRES_PASSWORD='1234' -p 5436:5432 -d --rm postgres

// Init schema dir
migrate create -ext sql -dir ./schema -seq init

// Init db
migrate -path ./schema -database 'postgres://postgres:1234@localhost:5436/postgres?sslmode=disable' up

// Remove db
migrate -path ./schema -database 'postgres://postgres:1234@localhost:5436/postgres?sslmode=disable' down

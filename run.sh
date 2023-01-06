#!/bin/bash

export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=Postgres2023!  # definido no docker-compose.yml para desenvolvimento

go run servidor.go

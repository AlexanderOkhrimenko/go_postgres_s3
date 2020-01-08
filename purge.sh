#!/usr/bin/env bash

docker-compose down --remove-orphans
docker image rm go_postgres_s3_worker
docker image rm go_postgres_s3_api
docker volume rm go_postgres_s3_database_postgres
docker volume rm go_postgres_s3_minio
#!/bin/sh

set +e

#  --mount type=bind,source="$(pwd)"/scripts/init.sql,target=/docker-entrypoint-initdb.d/init.sql\
docker run -d --rm --name postgres_test -p '5432:5432' \
	-e POSTGRES_USER='postgres' -e POSTGRES_PASSWORD='querty' \
	postgres:12.7-alpine

echo "Postgres container is starting"

RETRIES=5

until psql postgresql://postgres:querty@localhost:5432/postgres -c "select 1" >/dev/null 2>&1 || [ $RETRIES -eq 0 ]; do
	echo "Waiting for postgres server, $((RETRIES -= 1)) remaining attempts..."
	sleep 1
done

echo "Postgres is ready"

env $(cat .env | xargs) go test -cover ./...

docker stop postgres_test

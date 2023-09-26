postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=hectico2511 postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root go-tool

dropdb:
	docker exec -it postgres15 dropdb go-tool
	
#!make
include .env

PHONY: 

run:
	go run cmd/server/main.go

build-neo4j:
	docker run \
		--name neo4jVs \
		--publish=7474:7474 --publish=${DB_NEO4J_PORT}:${DB_NEO4J_PORT} \
		--volume=$$HOME/neo4j/data:/data \
		--volume=$$HOME/neo4j/conf:/var/lib/neo4j/conf \
		--env=NEO4J_dbms_default__database=${DB_NEO4J_DATABASE} \
		--env=NEO4J_AUTH=none \
		neo4j:latest

run-neo4j:
	docker run neo4jVs

	
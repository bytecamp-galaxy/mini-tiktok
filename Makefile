.PHONY: docker-compose docker-build api-server

docker-compose:
	docker-compose up

docker-build:
	docker build -f docker/api-server/Dockerfile -t vgalaxy/api-server .

api-server:
	docker run --network=host vgalaxy/api-server
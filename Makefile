.PHONY: docker-compose docker-build api-server

docker-compose:
	docker-compose up

docker-build:
	docker build -f docker/api-server/Dockerfile -t vgalaxy/api-server .
	docker build -f docker/user-server/Dockerfile -t vgalaxy/user-server .

api-server:
	docker run --network=host -p 8080:8080 vgalaxy/api-server
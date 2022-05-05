build:
	docker-compose -f docker/docker-compose.yml up -d --build

remove:
	docker-compose -f docker/docker-compose.yml down
	docker rmi my-go-app
	docker volume prune --force
.PHONY: up down start stop restart

up:
	docker-compose -f deployments/docker-compose.yml --project-name btchash160 up

down:
	docker-compose -f deployments/docker-compose.yml --project-name btchash160 down

start:
	docker-compose -f deployments/docker-compose.yml --project-name btchash160 start

stop:
	docker-compose -f deployments/docker-compose.yml --project-name btchash160 stop

restart:
	docker-compose -f deployments/docker-compose.yml --project-name btchash160 restart

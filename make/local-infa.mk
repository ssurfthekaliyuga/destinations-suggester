LOCAL_COMPOSE=./environments/local/docker-compose.yaml
LOCAL_PROJECT=local

up-local:
	docker compose -f $(LOCAL_COMPOSE) -p $(LOCAL_PROJECT) up -d --build

down-local:
	docker compose -f $(LOCAL_COMPOSE) -p $(LOCAL_PROJECT) down

restart-local: down-local up-local

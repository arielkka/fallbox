up-containers:
	docker-compose -f docker-compose.yml up

stop-containers:
	docker stop $(docker ps -a -q)

delete-containers:
	docker container prune $(docker ps -a -q)

migrate-up-handler-db:
	migrate -path ./handler/schema -database 'mysql://root:root@tcp(localhost:3307)/handler_db' up

migrate-up-png-db:
	migrate -path ./png/schema -database 'mysql://root:root@tcp(localhost:3308)/png_db' up

migrate-up-jpg-db:
	migrate -path ./jpg/schema -database 'mysql://root:root@tcp(localhost:3309)/jpg_db' up

migrate-down-handler-db:
	migrate -path ./handler/schema -database 'mysql://root:root@tcp(localhost:3307)/handler_db' down

migrate-down-png-db:
	migrate -path ./png/schema -database 'mysql://root:root@tcp(localhost:3307)/png_db' down

migrate-down-jpg-db:
	migrate -path ./jpg/schema -database 'mysql://root:root@tcp(localhost:3307)/jpg_db' down

migrate-dirty-handler-db:
	migrate -path ./handler/schema -database 'mysql://root:root@tcp(localhost:3307)/handler_db' force 1

migrate-dirty-png-db:
	migrate -path ./png/schema -database 'mysql://root:root@tcp(localhost:3307)/png_db' force 1

migrate-dirty-jpg-db:
	migrate -path ./jpg/schema -database 'mysql://root:root@tcp(localhost:3307)/jpg_db' force 1
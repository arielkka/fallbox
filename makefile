up-containers:
	docker-compose -f docker-compose.yml up

stop-containers:
	docker stop $(docker ps -a -q)

delete-containers:
	docker container prune $(docker ps -a -q)

migrate-up:
	migrate -path ./schema -database 'mysql://root:root@tcp(localhost:3307)/fallbox' up

migrate-down:
	migrate -path ./schema -database 'mysql://root:root@tcp(localhost:3307)/fallbox' down

migrate-dirty:
	migrate -path ./schema -database 'mysql://root:root@tcp(localhost:3307)/fallbox' force 1
up-containers:
	docker-compose -f docker-compose.yml up --build


stop-containers:
	docker stop $(docker ps -qa)

delete-containers:
	docker container prune $(docker ps -q)


migrate-up-all:
	make migrate-up-handler-db && make migrate-up-excel-db && make migrate-up-txt-db

migrate-down-all:
	make migrate-down-handler-db && make migrate-down-excel-db && make migrate-down-txt-db


migrate-up-handler-db:
	migrate -path ./handler/schema -database 'mysql://root:root@tcp(localhost:3307)/handler_db' up

migrate-up-excel-db:
	migrate -path ./excel/schema -database 'mysql://root:root@tcp(localhost:3308)/excel_db' up

migrate-up-txt-db:
	migrate -path ./txt/schema -database 'mysql://root:root@tcp(localhost:3309)/txt_db' up

migrate-down-handler-db:
	migrate -path ./handler/schema -database 'mysql://root:root@tcp(localhost:3307)/handler_db' down

migrate-down-excel-db:
	migrate -path ./excel/schema -database 'mysql://root:root@tcp(localhost:3308)/excel_db' down

migrate-down-txt-db:
	migrate -path ./txt/schema -database 'mysql://root:root@tcp(localhost:3309)/txt_db' down

migrate-dirty-handler-db:
	migrate -path ./handler/schema -database 'mysql://root:root@tcp(localhost:3307)/handler_db' force 1

migrate-dirty-excel-db:
	migrate -path ./excel/schema -database 'mysql://root:root@tcp(localhost:3308)/excel_db' force 1

migrate-dirty-txt-db:
	migrate -path ./txt/schema -database 'mysql://root:root@tcp(localhost:3309)/txt_db' force 1
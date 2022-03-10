DATABASE_URL="mysql://root:root@(127.0.0.1:3306)/todo_db?charset=utf8&parseTime=True&loc=Local&timeout=10ms"

db-migrate-up:
	migrate -database ${DATABASE_URL} -path database/migrations/ up

db-migrate-down:
	migrate -database ${DATABASE_URL} -path database/migrations/ down
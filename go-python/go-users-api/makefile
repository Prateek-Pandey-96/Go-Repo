postgres-init:
	docker run --name postgreDB -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres 

# before running init make sure to create the users db inside postgres
migrate-init:
	migrate create -ext sql -dir migrations/ -seq init_mg

migrate-up:
	migrate -path migrations/ -database "postgresql://postgres:secret@localhost:5432/users?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations/ -database "postgresql://postgres:secret@localhost:5432/users?sslmode=disable" -verbose down
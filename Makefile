postgres:
	docker run --name postgres  -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres:15-alpine  

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 
	
.PHONY: postgres migrateup migratedown
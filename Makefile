migrateup:
	migrate -path migrations -database "postgresql://thinhnn:secret@0.0.0.0:5432/testdb?sslmode=disable" -verbose up

migrateup1:
	migrate -path migrations -database "postgresql://thinhnn:secret@0.0.0.0:5432/testdb?sslmode=disable" -verbose up 1

migratedown:
	migrate -path migrations -database "postgresql://thinhnn:secret@0.0.0.0:5432/testdb?sslmode=disable" -verbose down

migratedown1:
	migrate -path migrations -database "postgresql://thinhnn:secret@0.0.0.0:5432/testdb?sslmode=disable" -verbose down 1

.PHONY: migrateup migratedown migrateup1 migratedown1
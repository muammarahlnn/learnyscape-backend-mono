include ${CURDIR}/.env

migratecreate:
	@migrate create -ext sql -dir ${CURDIR}/db/migrations/ -seq ${name}

migrateforce:
	@migrate -path ${CURDIR}/db/migrations/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose force 1

migratedown:
	@migrate -path ${CURDIR}/db/migrations/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose down

migrateup:
	@migrate -path ${CURDIR}/db/migrations/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose up
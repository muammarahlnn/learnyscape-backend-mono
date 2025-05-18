include ${CURDIR}/.env

migratecreate:
	@migrate create -ext sql -dir ${CURDIR}/db/migration/ -seq ${name}

migrateforce:
	@migrate -path ${CURDIR}/db/migration/ -database "${MIGRATE_DB_URL}" -verbose force 1

migratedown:
	@migrate -path ${CURDIR}/db/migration/ -database "${MIGRATE_DB_URL}" -verbose down

migrateup:
	@migrate -path ${CURDIR}/db/migration/ -database "${MIGRATE_DB_URL}" -verbose up

logs:
	@docker container logs learnyscape-mono-backend
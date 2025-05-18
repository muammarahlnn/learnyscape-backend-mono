include ${CURDIR}/.env

migrate_create:
	@migrate create -ext sql -dir ${CURDIR}/db/migration/ -seq ${name}

migrate_force:
	@migrate -path ${CURDIR}/db/migration/ -database "${MIGRATE_DB_URL}" -verbose force 1

migrate_down:
	@migrate -path ${CURDIR}/db/migration/ -database "${MIGRATE_DB_URL}" -verbose down

migrate_down_one:
	@migrate -path ${CURDIR}/db/migration/ -database "${MIGRATE_DB_URL}" -verbose down 1

migrate_up:
	@migrate -path ${CURDIR}/db/migration/ -database "${MIGRATE_DB_URL}" -verbose up

logs:
	@docker container logs learnyscape-mono-backend

up:
	@docker compose up -d --build
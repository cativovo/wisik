include .env
export

goose_up:
	goose -dir=./pkg/storage/postgres/schema up

sqlc_generate: goose_up
	sqlc generate -f ./pkg/storage/postgres/sqlc.yaml

dev:
	fd -e go -e html -E sqlc_* | entr -rc go run main.go

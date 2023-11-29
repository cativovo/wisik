dev:
	fd -e go -e html -E sqlc_* | entr -rc go run main.go

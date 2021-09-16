run:
	docker-compose up --remove-orphans

build-win64:
	GOOS=windows GOARCH=amd64 go build -o bin/workplace_client64.exe cmd/client/main.go

build-linux64:
	GOOS=linux GOARCH=amd64 go build -o bin/workplace_client_linux64 cmd/client/main.go

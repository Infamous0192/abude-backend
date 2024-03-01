run:
	go run $(FLAGS) main.go

migrate:
	go run $(FLAGS) main.go --migrate

dev:
	air

build:
	go build -o ./bin/application.exe main.go

install:
	go install $(FLAGS)

buildrun:
	go build main.go && .\main.exe

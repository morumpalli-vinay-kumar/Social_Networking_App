start:
	go run .

up:
	./goose up

down:
	./goose down
	
build:
	GOOS=linux GOARCH=amd64 go build /home/ubuntu/app/cmd/goose.go

clean:
	go mod tidy

add:
	git add .

gst:
	git status


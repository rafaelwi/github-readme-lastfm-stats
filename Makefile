deploy:
	mkdir -p api
	go get ./...
	go build -o api/card src/api/card.go

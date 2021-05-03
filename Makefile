deploy:
	mkdir -p functions
	go get ./...
	go build -o functions/card src/card.go

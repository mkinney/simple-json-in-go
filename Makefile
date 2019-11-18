run:
	go run main.go

test:
	go test -v

cover:
	go test -cover

both:
	go test -v -cover
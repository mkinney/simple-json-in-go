run:
	go run main.go

test:
	go test -v

short:
	go test -v -short

cover:
	go test -cover

both:
	go test -v -cover

parallel:
	go test -v -short -parallel 3

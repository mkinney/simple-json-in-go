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

benchtoo:
	go test -v -short -bench .

# run all tests that match 'x' (which there are none)
# so, it only runs the benchmark
bench:
	go test -run x -bench .

run:
	go run main.go

test:
	go test -v

short:
	go test -v -short

# just show percentage of code coverage
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

# show which lines have code coverage and which ones do not
coverage:
	go test -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html
	# open the html page in default browswer (from mac)
	open cover.html

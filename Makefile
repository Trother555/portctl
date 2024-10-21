.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run .

.PHONY: docker-build
docker-build:
	docker build --tag portctl .

.PHONY: docker-run
docker-run:
	docker run -it --rm -p 8080:8080 portctl -inPorts 4 -outPorts 4

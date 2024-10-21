### Usage

Run portctl with 4 input ports and 3 output ports:

```
go run ./cmd/main.go -inPorts 4 -outPorts 3
```

Alternatively, build docker container:

```
make docker-build
```

And run in docker:

```
make docker-run
docker run -it --rm -p 8080:8080 portctl -inPorts 4 -outPorts 4
```

Read port 2 value:

```
curl localhost:8080/read?portNum=2
```

Write to port 2 transaction 4 value 5:

```
curl -X POST "localhost:8080/write?portNum=2&transactionId=4&val=5"
```

### Tests

```
make test
```

# gonats

### Go nuts with NATS!

This is to demonstrate how to pause and resume message publishing.  
To Run the code:

In terminal 1

```bash
go get github.com/nats-io/nats-server
nats-server
```

In terminal 2

```bash
cd cmd/subscribe
go run main.go lorem
```

In terminal 3

```bash
cd cmd/publish
go run main.go lorem test-data.txt
```

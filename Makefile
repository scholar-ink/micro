build:
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -i -o micro ./main.go ./plugins.go
    docker build -t ttouch/micro .
run:
docker run --link=etcd -p 8080:8080 -e MICRO_REGISTRY=etcd -e MICRO_REGISTRY_ADDRESS=http://etcd:2379 ttouch/micro api --handle=proxy